package data

import (
	"content-manage/internal/biz"
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"time"
)

type contentRepo struct {
	data *Data
	log  *log.Helper
}

func (receiver *contentRepo) Find(ctx context.Context, param *biz.FindParams) ([]*biz.Content, int64, error) {
	query := receiver.data.db.Model(&ContentDetail{})
	if param.ID != 0 {
		query = query.Where("id=?", param.ID)
	}
	if param.Author != "" {
		query = query.Where("author=?", param.Author)
	}
	if param.Title != "" {
		query = query.Where("title=?", param.Title)
	}
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var page, pageSize = 1, 10
	if param.Page > 0 {
		page = param.Page
	}
	if param.PageSize > 0 {
		pageSize = param.PageSize
	}
	offset := (page - 1) * pageSize
	var data []*ContentDetail
	var result []*biz.Content
	if err := query.Offset(offset).Limit(pageSize).Find(&data).Error; err != nil {
		return nil, 0, err
	}
	for _, r := range data {
		result = append(result, &biz.Content{
			Id:             r.ID,
			Title:          r.Title,
			VideoURL:       r.VideoURL,
			Author:         r.Author,
			Description:    r.Description,
			Thumbnail:      r.Thumbnail,
			Category:       r.Category,
			Duration:       r.Duration,
			Resolution:     r.Resolution,
			FileSize:       r.FileSize,
			Format:         r.Format,
			Quality:        r.Quality,
			ApprovalStatus: r.ApprovalStatus,
		})
	}
	return result, total, nil
}

func (receiver *contentRepo) IsExist(ctx context.Context, id int64) (bool, error) {
	db := receiver.data.db
	var detail ContentDetail
	err := db.Where("id =?", id).First(&detail).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	if err != nil {
		receiver.log.WithContext(ctx).Errorf("ContentDao isExist = %v", err)
		return false, err
	}
	return true, nil
}

func (receiver *contentRepo) Delete(ctx context.Context, id int64) error {
	db := receiver.data.db
	err := db.Where("id =?", id).Delete(&ContentDetail{}).Error
	if err != nil {
		receiver.log.Error("delete error =%v", err)
		return err
	}
	return nil
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
