package upload

import (
	"context"
	"hotgo/api/api/meme/v1/upload"
	"hotgo/internal/consts"
	"hotgo/internal/library/storager"
	"hotgo/internal/service/meme"

	"github.com/gogf/gf/v2/errors/gerror"
)

type sUpload struct{}

func NewUpload() *sUpload {
	return &sUpload{}
}

func init() {
	meme.RegisterUpload(NewUpload())
}

// UploadFile 上传文件
func (s *sUpload) UploadFile(ctx context.Context, req *upload.UploadReq) (res *upload.UploadRes, err error) {
	if req.File == nil {
		err = gerror.New("请选择要上传的文件")
		return
	}

	// 使用阿里云OSS驱动
	drive := storager.New(consts.UploadDriveOss)
	fullPath, err := drive.Upload(ctx, req.File)
	if err != nil {
		return
	}

	res = &upload.UploadRes{
		Name:     req.File.Filename,
		Url:      fullPath,
		Size:     req.File.Size,
		MimeType: req.File.Header.Get("Content-Type"),
	}
	return
}
