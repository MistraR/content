package service

import (
	"content-manage/api/operate"
	"content-manage/internal/biz"
	"context"
	"time"
)

func (a *AppService) UpdateContent(ctx context.Context, req *operate.UpdateContentRequest) (*operate.UpdateContentReply, error) {
	uc := a.uc
	content := req.GetContent()
	err := uc.UpdateContent(ctx, &biz.Content{
		Id:             content.GetId(),
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
	return &operate.UpdateContentReply{}, nil
}
