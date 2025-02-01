package service

import (
	"content-manage/api/operate"
	"context"
)

func (a *AppService) DeleteContent(ctx context.Context, req *operate.DeleteContentRequest) (*operate.DeleteContentReply, error) {
	uc := a.uc
	err := uc.DeleteContent(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &operate.DeleteContentReply{}, nil
}
