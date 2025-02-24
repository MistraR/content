// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"content-manage/internal/biz"
	"content-manage/internal/conf"
	"content-manage/internal/data"
	"content-manage/internal/server"
	"content-manage/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	contentRepo := data.NewContentRepo(dataData, logger)
	contentUsecase := biz.NewContentcase(contentRepo, logger)
	appService := service.NewAppService(contentUsecase)
	grpcServer := server.NewGRPCServer(confServer, appService, logger)
	app := newApp(logger, grpcServer)
	return app, func() {
		cleanup()
	}, nil
}
