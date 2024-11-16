package biz

import (
	"golang.org/x/net/context"
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
	Quality        int           `json:"quality"`                        //视频质量
	ApprovalStatus int           `json:"approval_status"`                //审核状态
}

type ContentRepo interface {
	Create(context.Context, *Content) error
	Save(context.Context, *Content) (*Content, error)
}
