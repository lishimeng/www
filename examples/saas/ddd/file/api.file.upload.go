package file

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/app-starter/tool"
)

// FileUploadResp 文件上传响应结构体
type FileUploadResp struct {
	app.Response
	Data FileInfo `json:"data"`
}

// FileInfo 文件信息结构体
type FileInfo struct {
	Name string `json:"name"` // 文件名
	Url  string `json:"url"`  // 文件访问URL
}

// uploadFile 处理文件上传
func uploadFile(ctx server.Context) {
	var resp FileUploadResp

	// 获取上传的文件，尝试不同的字段名
	var fileHeader *multipart.FileHeader
	var src multipart.File
	var err error

	// 首先尝试 "file" 字段
	src, fileHeader, err = ctx.C.FormFile("file")
	if err != nil {
		// 尝试获取所有文件字段
		form := ctx.C.Request().MultipartForm
		if form == nil {
			err := ctx.C.Request().ParseMultipartForm(32 << 20) // 32MB
			if err != nil {
				resp.Code = tool.RespCodeError
				resp.Message = "请选择要上传的文件"
				ctx.Json(resp)
				return
			}
			form = ctx.C.Request().MultipartForm
		}

		// 查找第一个文件
		if form == nil || len(form.File) == 0 {
			resp.Code = tool.RespCodeError
			resp.Message = "请选择要上传的文件"
			ctx.Json(resp)
			return
		}

		// 获取第一个文件字段的第一个文件
		for _, files := range form.File {
			if len(files) > 0 {
				fileHeader = files[0]
				// 打开文件
				var openErr error
				src, openErr = files[0].Open()
				if openErr != nil {
					resp.Code = tool.RespCodeError
					resp.Message = fmt.Sprintf("打开文件失败: %v", openErr)
					ctx.Json(resp)
					return
				}
				break
			}
		}

		if fileHeader == nil {
			resp.Code = tool.RespCodeError
			resp.Message = "请选择要上传的文件"
			ctx.Json(resp)
			return
		}
	}

	// 创建上传目录（如果不存在）
	uploadDir := "./uploads"
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		err := os.MkdirAll(uploadDir, 0755)
		if err != nil {
			resp.Code = tool.RespCodeError
			resp.Message = fmt.Sprintf("创建上传目录失败: %v", err)
			ctx.Json(resp)
			return
		}
	}

	// 生成唯一的文件名（时间戳 + 原文件名）
	timestamp := time.Now().Format("20060102150405")
	ext := filepath.Ext(fileHeader.Filename)
	name := fileHeader.Filename[:len(fileHeader.Filename)-len(ext)]
	newFileName := fmt.Sprintf("%s_%s%s", name, timestamp, ext)
	filePath := filepath.Join(uploadDir, newFileName)

	// src 已经通过 FormFile 打开，这里直接使用
	defer src.Close()

	// 创建目标文件
	dst, err := os.Create(filePath)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = fmt.Sprintf("创建文件失败: %v", err)
		ctx.Json(resp)
		return
	}
	defer dst.Close()

	// 复制文件内容
	_, err = io.Copy(dst, src)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = fmt.Sprintf("保存文件失败: %v", err)
		ctx.Json(resp)
		return
	}

	// 生成文件访问URL（相对路径）
	fileUrl := fmt.Sprintf("/api/v1/files/%s", newFileName)

	resp.Code = tool.RespCodeSuccess
	resp.Message = "上传成功"
	resp.Data = FileInfo{
		Name: fileHeader.Filename,
		Url:  fileUrl,
	}
	ctx.Json(resp)
}
