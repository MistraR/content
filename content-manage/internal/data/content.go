package data

import (
	"content-manage/internal/biz"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type contentRepo struct {
	data *Data
	log  *log.Helper
}

func (receiver *contentRepo) Update(ctx context.Context, id int64, content *biz.Content) error {
	db := receiver.data.db
	detail := ContentDetail{
		ID:             id,
		Title:          content.Title,
		Description:    content.Description,
		Author:         content.Author,
		VideoURL:       content.VideoURL,
		Thumbnail:      content.Thumbnail,
		Category:       content.Category,
		Duration:       content.Duration,
		Resolution:     content.Resolution,
		FileSize:       content.FileSize,
		Format:         content.Format,
		Quality:        content.Quality,
		ApprovalStatus: content.ApprovalStatus,
	}
	if err := db.Where("id =?", id).Updates(&detail).Error; err != nil {
		receiver.log.WithContext(ctx).Errorf("content update error = %v", err)
		return err
	}
	return nil
}

// NewGreeterRepo .
func NewContentRepo(data *Data, logger log.Logger) biz.ContentRepo {
	return &contentRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

type ContentDetail struct {
	ID             int64         `gorm:"column:id;primary_key"`  //ID
	Title          string        `gorm:"column:title"`           //内容标题
	Description    string        `gorm:"column:description"`     //描述
	Author         string        `gorm:"column:author"`          //作者
	VideoURL       string        `gorm:"column:video_url"`       //视频url
	Thumbnail      string        `gorm:"column:thumbnail"`       //封面图url
	Category       string        `gorm:"column:category"`        //分类
	Duration       time.Duration `gorm:"column:duration"`        //时长
	Resolution     string        `gorm:"column:resolution"`      //分辨率
	FileSize       int64         `gorm:"column:file_size"`       //文件大小
	Format         string        `gorm:"column:format"`          //格式
	Quality        int64         `gorm:"column:quality"`         //视频质量
	ApprovalStatus int64         `gorm:"column:approval_status"` //审核状态
	UpdatedAt      time.Time     `gorm:"column:updated_at"`      //更新时间
	CreatedAt      time.Time     `gorm:"column:created_at"`      //创建时间
}

func (a ContentDetail) TableName() string {
	table := "cms_content.t_content_details"
	return table
}

func (receiver *contentRepo) Create(ctx context.Context, content *biz.Content) error {
	receiver.log.Infof("contentRepo22222 create content = %+v", content)
	detail := ContentDetail{
		Title:          content.Title,
		Description:    content.Description,
		Author:         content.Author,
		VideoURL:       content.VideoURL,
		Thumbnail:      content.Thumbnail,
		Category:       content.Category,
		Duration:       content.Duration,
		Resolution:     content.Resolution,
		FileSize:       content.FileSize,
		Format:         content.Format,
		Quality:        content.Quality,
		ApprovalStatus: content.ApprovalStatus,
	}
	db := receiver.data.db
	if err := db.Create(&detail).Error; err != nil {
		receiver.log.Error("content create error =%v", err)
		return err
	}
	return nil
}
