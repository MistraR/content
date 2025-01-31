package data

import (
	"content-manage/internal/biz"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type contentRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewContentRepo(data *Data, logger log.Logger) biz.ContentRepo {
	return &contentRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (receiver *contentRepo) Create(ctx context.Context, content *biz.Content) error {
	receiver.log.Infof("contentRepo create content = %+v", content)
	return nil
}

//func (c *contentRepo) Create(ctx context.Context, content *biz.ContentRepo) error {
//	c.log.Infof("contentRepo Create content = %+v", content)
//	return nil
//}
//func (r *contentRepo) Save(ctx context.Context, g *biz.Content) (*biz.Content, error) {
//	return g, nil
//}
