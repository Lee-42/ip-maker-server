package upload

import (
	"context"
	"hotgo/api/api/ipmaker/v1/upload"
	"hotgo/internal/service/ipmaker"
)

var Upload = cUpload{}

type cUpload struct{}

// UploadFile 上传文件
func (c *cUpload) UploadFile(ctx context.Context, req *upload.UploadReq) (res *upload.UploadRes, err error) {
	res, err = ipmaker.Upload().UploadFile(ctx, req)
	return
}
