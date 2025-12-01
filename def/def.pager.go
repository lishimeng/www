package def

import "github.com/lishimeng/app-starter/server"

var ( // 可修改
	QueryPageNum         = "page_no"
	QueryPageSize        = "page_size"
	DefaultQueryPageNum  = 1
	DefaultQueryPageSize = 5
)

// GetParamPager 自动读Page参数, 默认pageSize=5
func GetParamPager(ctx server.Context) (pageNum, pageSize int) {
	pageSize = ctx.C.URLParamIntDefault(QueryPageSize, DefaultQueryPageSize)
	pageNum = ctx.C.URLParamIntDefault(QueryPageNum, DefaultQueryPageNum)
	return
}
