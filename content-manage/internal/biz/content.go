package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type Content struct {
	Id             int64         `json:"id"`                             //内容标题
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

type FindParams struct {
	ID       int64
	Author   string
	Title    string
	Page     int
	PageSize int
}

type ContentRepo interface {
	Create(context.Context, *Content) error
	Update(context.Context, int64, *Content) error
	Delete(context.Context, int64) error
	IsExist(context.Context, int64) (bool, error)
	Find(context.Context, *FindParams) ([]*Content, int64, error)
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

func (uc *ContentUsecase) UpdateContent(ctx context.Context, g *Content) error {
	uc.log.WithContext(ctx).Infof("UpdateContent: %v", g.Title)
	return uc.repo.Update(ctx, g.Id, g)
}

func (uc *ContentUsecase) DeleteContent(ctx context.Context, id int64) error {
	uc.log.WithContext(ctx).Infof("DeleteContent: %v", id)
	repo := uc.repo
	ok, err := repo.IsExist(ctx, id)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New(10001, "内容不存在", "")
	}
	return uc.repo.Delete(ctx, id)
}

func (uc *ContentUsecase) FindContent(ctx context.Context, params *FindParams) ([]*Content, int64, error) {
	repo := uc.repo
	contents, total, err := repo.Find(ctx, params)
	if err != nil {
		return nil, 0, err
	}
	return contents, total, err
}
