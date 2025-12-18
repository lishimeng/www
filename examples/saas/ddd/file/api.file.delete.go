package file

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"

	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/app-starter/tool"
)

// FileDeleteResp 文件删除响应结构体
type FileDeleteResp struct {
	app.Response
}

// deleteFile 删除文件
func deleteFile(ctx server.Context) {
	var resp FileDeleteResp

	// 获取文件路径参数
	filePath := ctx.C.URLParam("filePath")
	if filePath == "" {
		resp.Code = tool.RespCodeError
		resp.Message = "文件路径不能为空"
		ctx.Json(resp)
		return
	}

	// URL解码文件路径（处理中文字符等特殊字符）
	decodedFilePath, err := url.QueryUnescape(filePath)
	if err != nil {
		// 如果解码失败，使用原始路径
		decodedFilePath = filePath
	}

	// 从URL路径中提取文件名
	fileName := filepath.Base(decodedFilePath)
	if fileName == "" || fileName == "." || fileName == "/" {
		resp.Code = tool.RespCodeError
		resp.Message = "无效的文件路径"
		ctx.Json(resp)
		return
	}

	// URL解码文件名（处理中文字符等特殊字符）
	decodedFileName, decodeErr := url.PathUnescape(fileName)
	if decodeErr != nil {
		// 如果解码失败，使用原始文件名
		decodedFileName = fileName
	}

	// 构建完整的文件路径
	fullPath := filepath.Join("./uploads", decodedFileName)

	// 检查文件是否存在
	if _, statErr := os.Stat(fullPath); os.IsNotExist(statErr) {
		resp.Code = tool.RespCodeNotFound
		resp.Message = "文件不存在"
		ctx.Json(resp)
		return
	}

	// 删除文件
	if err := os.Remove(fullPath); err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = fmt.Sprintf("删除文件失败: %v", err)
		ctx.Json(resp)
		return
	}

	resp.Code = tool.RespCodeSuccess
	resp.Message = "删除成功"
	ctx.Json(resp)
}
