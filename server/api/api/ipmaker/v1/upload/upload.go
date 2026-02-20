package upload

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// UploadReq 上传文件请求
type UploadReq struct {
	g.Meta `path:"/upload/file" method:"post" tags:"上传" summary:"上传文件" mime:"multipart/form-data"`
	File   *ghttp.UploadFile `json:"file" type:"file" dc:"文件"`
}

// UploadRes 上传文件响应
type UploadRes struct {
	Name     string `json:"name" dc:"文件名"`
	Url      string `json:"url" dc:"文件地址"`
	Size     int64  `json:"size" dc:"文件大小"`
	MimeType string `json:"mimeType" dc:"文件类型"`
}
