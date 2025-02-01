package service

import (
	"content-manage/api/operate"
	"content-manage/internal/biz"
	"context"
)

func (a *AppService) FindContent(ctx context.Context, req *operate.FindContentRequest) (*operate.FindContentReply, error) {
	findParams := &biz.FindParams{
		ID:       req.GetId(),
		Author:   req.GetAuthor(),
		Title:    req.GetTitle(),
		Page:     int(req.GetPage()),
		PageSize: int(req.GetPageSize()),
	}
	uc := a.uc
	results, total, err := uc.FindContent(ctx, findParams)
	if err != nil {
		return nil, err
	}
	var contents []*operate.Content
	for _, r := range results {
		contents = append(contents, &operate.Content{
			Id:             r.Id,
			Title:          r.Title,
			VideoUrl:       r.VideoURL,
			Author:         r.Author,
			Description:    r.Description,
			Thumbnail:      r.Thumbnail,
			Category:       r.Category,
			Duration:       r.Duration.Milliseconds(),
			Resolution:     r.Resolution,
			FileSize:       r.FileSize,
			Format:         r.Format,
			Quality:        r.Quality,
			ApprovalStatus: r.ApprovalStatus,
		})
	}
	rsp := &operate.FindContentReply{
		Total:    total,
		Contents: contents,
	}
	return rsp, nil
}
