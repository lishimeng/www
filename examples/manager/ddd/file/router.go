package file

import (
	"github.com/lishimeng/app-starter/server"
)

// Router 文件管理路由注册
func Router(root server.Router) {
	// 文件上传
	root.Post("/", uploadFile)
	// 文件删除
	root.Delete("/", deleteFile)
	// 文件下载/查看
	// 使用路径参数匹配所有GET请求，支持包含中文字符的文件名
	root.Get("/{path:path}", getFile)
}
