package service

import (
	"content-manage/api/operate"
	"content-manage/internal/biz"
	"context"
)

func (a *AppService) CreateContent(ctx context.Context, req operate.CreateContentRequest) (*operate.CreateContentReply, error) {
	uc := a.uc
	err := uc.CreateContent(ctx, &biz.Content{
		Title:          "",
		Description:    "",
		Author:         "",
		VideoURL:       "",
		Thumbnail:      "",
		Category:       "",
		Duration:       0,
		Resolution:     "",
		FileSize:       0,
		Format:         "",
		Quality:        0,
		ApprovalStatus: 0,
	})
	if err != nil {
		return nil, err
	}
	return &operate.CreateContentReply{}, nil
}
