package upload

import (
	"context"
	"hotgo/api/api/meme/v1/upload"
	"hotgo/internal/service/meme"
)

var Upload = cUpload{}

type cUpload struct{}

// UploadFile 上传文件
func (c *cUpload) UploadFile(ctx context.Context, req *upload.UploadReq) (res *upload.UploadRes, err error) {
	res, err = meme.Upload().UploadFile(ctx, req)
	return
}
