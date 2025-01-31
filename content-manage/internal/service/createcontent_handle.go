package service

import (
	"content-manage/api/operate"
	"content-manage/internal/biz"
	"context"
	"time"
)

func (a *AppService) CreateContent(ctx context.Context, req *operate.CreateContentRequest) (*operate.CreateContentReply, error) {
	uc := a.uc
	content := req.GetContent()
	err := uc.CreateContent(ctx, &biz.Content{
		Title:          content.GetTitle(),
		Description:    content.GetDescription(),
		Author:         content.GetAuthor(),
		VideoURL:       content.GetVideoUrl(),
		Thumbnail:      content.GetThumbnail(),
		Category:       content.GetCategory(),
		Duration:       time.Duration(content.GetDuration()),
		Resolution:     content.GetResolution(),
		FileSize:       content.GetFileSize(),
		Format:         content.GetFormat(),
		Quality:        content.GetQuality(),
		ApprovalStatus: content.GetApprovalStatus(),
	})
	if err != nil {
		return nil, err
	}
	return &operate.CreateContentReply{}, nil
}
