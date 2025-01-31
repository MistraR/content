package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type Content struct {
	Title          string        `json:"title" binding:"required"`       //内容标题
	Description    string        `json:"description" binding:"required"` //描述
	Author         string        `json:"author" binding:"required"`      //作者
	VideoURL       string        `json:"video_url"`                      //视频url
	Thumbnail      string        `json:"thumbnail"`                      //封面图url
	Category       string        `json:"category"`                       //分类
	Duration       time.Duration `json:"duration"`                       //时长
	Resolution     string        `json:"resolution"`                     //分辨率
	FileSize       int64         `json:"file_size"`                      //文件大小
	Format         string        `json:"format"`                         //格式
	Quality        int64         `json:"quality"`                        //视频质量
	ApprovalStatus int64         `json:"approval_status"`                //审核状态
}

type ContentRepo interface {
	Create(context.Context, *Content) error
}

type ContentUsecase struct {
	repo ContentRepo
	log  *log.Helper
}

func NewContentcase(repo ContentRepo, logger log.Logger) *ContentUsecase {
	return &ContentUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateGreeter creates a Greeter, and returns the new Greeter.
func (uc *ContentUsecase) CreateContent(ctx context.Context, g *Content) error {
	uc.log.WithContext(ctx).Infof("CreateContent: %v", g.Title)
	return uc.repo.Create(ctx, g)
}
