package ipmaker

import (
	"context"
	"hotgo/api/api/ipmaker/v1/upload"
)

type (
	IUpload interface {
		// UploadFile 上传文件
		UploadFile(ctx context.Context, req *upload.UploadReq) (res *upload.UploadRes, err error)
	}
)

var (
	localUpload IUpload
)

func Upload() IUpload {
	if localUpload == nil {
		panic("implement not found for interface IUpload, forgot register?")
	}
	return localUpload
}

func RegisterUpload(i IUpload) {
	localUpload = i
}
