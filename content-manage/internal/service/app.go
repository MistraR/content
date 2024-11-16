package service

import (
	"content-manage/api/operate"
	"content-manage/internal/biz"
)

type AppService struct {
	operate.UnimplementedAppServer
	uc *biz.Contentcase
}

func NewAppService(uc *biz.Contentcase) *AppService {
	return &AppService{uc: uc}
}
