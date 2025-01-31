package service

import (
	"content-manage/api/operate"
	"content-manage/internal/biz"
)

type AppService struct {
	operate.UnimplementedAppServer
	uc *biz.ContentUsecase
}

func NewAppService(uc *biz.ContentUsecase) *AppService {
	return &AppService{uc: uc}
}
