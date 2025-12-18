package file

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/go-log"
)

// getFile 获取文件（用于下载或查看）
func getFile(ctx server.Context) {
	requestURI := ctx.C.Request().RequestURI
	log.Debug("File request - RequestURI: %s, Method: %s, Path: %s",
		requestURI, ctx.C.Request().Method, ctx.C.Path())

	var fileName string
	var rawFileName string // 保存原始编码的文件名

	// 方法1：优先从RequestURI中提取原始编码的文件名（保留URL编码）
	log.Info("File request - RequestURI: %s", requestURI)

	if strings.Contains(requestURI, "/api/v1/files/") {
		parts := strings.Split(requestURI, "/api/v1/files/")
		if len(parts) > 1 {
			// 获取文件名部分（去除查询参数）
			filePart := strings.Split(parts[1], "?")[0]
			rawFileName = filePart
			log.Info("Extracted from RequestURI: %s", rawFileName)
		}
	}

	// 方法2：从路径参数中获取（Iris框架可能已经自动解码）
	// 尝试多种参数名
	fileName = ctx.C.Params().Get("path")
	if fileName == "" {
		fileName = ctx.C.Params().Get("*")
	}
	log.Info("File request - Param path/*: %s", fileName)

	// 方法3：如果方法2为空，从Path中提取
	if fileName == "" {
		requestURL := ctx.C.Request().URL.Path
		log.Info("File request - URL.Path: %s", requestURL)
		if strings.HasPrefix(requestURL, "/api/v1/files/") {
			fileName = strings.TrimPrefix(requestURL, "/api/v1/files/")
			log.Info("Extracted from URL.Path: %s", fileName)
		}
	}

	// 如果rawFileName为空，使用fileName
	if rawFileName == "" {
		rawFileName = fileName
	}

	if fileName == "" && rawFileName == "" {
		ctx.C.StatusCode(400)
		ctx.C.WriteString("文件路径不能为空")
		return
	}

	// 记录原始文件名和请求信息（用于调试）
	log.Info("File request - fileName: %s, rawFileName: %s, Param: %s, Path: %s, RawPath: %s, URL.Path: %s, RequestURI: %s",
		fileName, rawFileName, ctx.C.Params().Get("*"), ctx.C.Path(), ctx.C.Request().URL.RawPath, ctx.C.Request().URL.Path, ctx.C.Request().RequestURI)

	// URL解码文件名（处理中文字符等特殊字符）
	// Iris框架在获取通配符参数时可能已经自动解码，也可能没有
	// 我们需要尝试多种方式
	var candidates []string
	seen := make(map[string]bool) // 避免重复添加

	// 添加候选文件名的辅助函数
	addCandidate := func(candidate string) {
		if candidate != "" && !seen[candidate] {
			candidates = append(candidates, candidate)
			seen[candidate] = true
		}
	}

	// 候选1：如果fileName不为空，直接使用（可能已被Iris解码）
	if fileName != "" {
		addCandidate(fileName)
	}

	// 候选2：如果rawFileName不为空且与fileName不同，添加原始编码的文件名
	if rawFileName != "" && rawFileName != fileName {
		addCandidate(rawFileName)
	}

	// 候选3：尝试解码rawFileName（如果包含编码字符或看起来像是编码后的）
	if rawFileName != "" {
		// 如果包含 %，明确是URL编码
		if strings.Contains(rawFileName, "%") {
			// 尝试 PathUnescape（用于路径参数，优先）
			if decoded, err := url.PathUnescape(rawFileName); err == nil && decoded != rawFileName {
				addCandidate(decoded)
				log.Info("PathUnescape decoded: %s -> %s", rawFileName, decoded)
			} else if err != nil {
				log.Warn("PathUnescape failed for %s: %v", rawFileName, err)
			}
			// 尝试 QueryUnescape（用于查询参数）
			if decoded, err := url.QueryUnescape(rawFileName); err == nil && decoded != rawFileName {
				addCandidate(decoded)
				log.Info("QueryUnescape decoded: %s -> %s", rawFileName, decoded)
			} else if err != nil {
				log.Warn("QueryUnescape failed for %s: %v", rawFileName, err)
			}
		} else {
			// 即使不包含%，如果rawFileName和fileName不同，也尝试使用rawFileName
			if rawFileName != fileName {
				addCandidate(rawFileName)
			}
		}
	}

	// 候选4：如果fileName包含编码字符，也尝试解码
	if fileName != "" && strings.Contains(fileName, "%") {
		if decoded, err := url.PathUnescape(fileName); err == nil && decoded != fileName {
			addCandidate(decoded)
			log.Debug("fileName PathUnescape decoded: %s -> %s", fileName, decoded)
		}
		if decoded, err := url.QueryUnescape(fileName); err == nil && decoded != fileName {
			addCandidate(decoded)
			log.Debug("fileName QueryUnescape decoded: %s -> %s", fileName, decoded)
		}
	}

	// 候选5：尝试从RawPath获取（如果Iris提供了）
	if rawPath := ctx.C.Request().URL.RawPath; rawPath != "" && strings.HasPrefix(rawPath, "/api/v1/files/") {
		rawPathFileName := strings.TrimPrefix(rawPath, "/api/v1/files/")
		if rawPathFileName != "" && rawPathFileName != fileName && rawPathFileName != rawFileName {
			addCandidate(rawPathFileName)
			log.Debug("Using RawPath: %s", rawPathFileName)
			// 如果包含编码，也解码
			if strings.Contains(rawPathFileName, "%") {
				if decoded, err := url.PathUnescape(rawPathFileName); err == nil && decoded != rawPathFileName {
					addCandidate(decoded)
					log.Debug("RawPath PathUnescape decoded: %s -> %s", rawPathFileName, decoded)
				}
			}
		}
	}

	// 尝试每个候选文件名
	var fullPath string
	var found bool

	// 获取uploads目录的绝对路径
	// 尝试多种方式找到uploads目录
	uploadDir := "./uploads"
	if workDir, err := os.Getwd(); err == nil {
		// 方法1：当前工作目录下的uploads
		candidateDir := filepath.Join(workDir, "uploads")
		if _, err := os.Stat(candidateDir); err == nil {
			uploadDir = candidateDir
		} else {
			// 方法2：在项目根目录查找（假设工作目录在cmd/mes下）
			// 尝试向上查找项目根目录
			projectRoot := workDir
			for i := 0; i < 3; i++ {
				testDir := filepath.Join(projectRoot, "uploads")
				if _, err := os.Stat(testDir); err == nil {
					uploadDir = testDir
					break
				}
				// 向上查找
				parent := filepath.Dir(projectRoot)
				if parent == projectRoot {
					break
				}
				projectRoot = parent
			}
		}
	}

	// 记录使用的上传目录（仅在调试时输出）
	log.Debug("Using upload directory: %s", uploadDir)

	for i, candidate := range candidates {
		// 清理候选文件名（去除可能的前导/尾随空格和斜杠）
		candidate = strings.TrimSpace(candidate)
		candidate = strings.Trim(candidate, "/")
		if candidate == "" {
			log.Info("Skipping empty candidate at index %d", i)
			continue
		}

		// 构建完整文件路径
		fullPath = filepath.Join(uploadDir, candidate)

		// 检查文件是否存在
		if fileInfo, err := os.Stat(fullPath); err == nil && !fileInfo.IsDir() {
			found = true
			log.Info("File found (candidate %d/%d): %s -> %s", i+1, len(candidates), candidate, fullPath)
			break
		}

		log.Info("File not found (candidate %d/%d), tried: %s (candidate: %q)", i+1, len(candidates), fullPath, candidate)
	}

	// 如果所有候选都不存在，返回404
	if !found {
		// 尝试列出目录中的文件（仅用于调试，帮助定位问题）
		checkDir := "./uploads"
		if workDir, err := os.Getwd(); err == nil {
			checkDir = filepath.Join(workDir, "uploads")
		}

		var existingFiles []string
		if entries, err := os.ReadDir(checkDir); err == nil {
			for _, entry := range entries {
				if !entry.IsDir() {
					existingFiles = append(existingFiles, entry.Name())
				}
			}
		}

		// 尝试找到相似的文件名（可能是编码问题）
		var similarFiles []string
		if rawFileName != "" && len(candidates) > 0 {
			// 解码后的文件名应该是最后一个候选
			expectedName := candidates[len(candidates)-1]
			for _, f := range existingFiles {
				if strings.Contains(f, expectedName) || strings.Contains(expectedName, f) {
					similarFiles = append(similarFiles, f)
				}
			}
		}

		errorMsg := fmt.Sprintf("文件不存在\n原始请求: %s\n解码后的候选文件名 (%d个):\n", rawFileName, len(candidates))
		for i, cand := range candidates {
			errorMsg += fmt.Sprintf("  [%d] %q\n", i+1, cand)
		}
		errorMsg += fmt.Sprintf("上传目录: %s\n目录中现有文件数: %d", checkDir, len(existingFiles))

		// 如果目录中文件不多，列出所有文件名（最多20个）
		if len(existingFiles) > 0 && len(existingFiles) <= 20 {
			errorMsg += "\n现有文件列表:\n"
			for _, f := range existingFiles {
				errorMsg += fmt.Sprintf("  - %q\n", f)
			}
		}

		// 如果有相似的文件，特别提示
		if len(similarFiles) > 0 {
			errorMsg += "\n相似的文件名（可能是编码问题）:\n"
			for _, f := range similarFiles {
				errorMsg += fmt.Sprintf("  - %q\n", f)
			}
		}

		log.Error("File not found. %s", errorMsg)
		ctx.C.StatusCode(404)
		ctx.C.WriteString(errorMsg)
		return
	}

	// 发送文件
	ctx.C.SendFile(fullPath, "")
}
