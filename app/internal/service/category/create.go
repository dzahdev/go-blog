package service

import (
	"context"
	"dzrise.ru/internal/converter"
	"dzrise.ru/internal/model"
	"github.com/dzahdev/slugGenerator"
)

func (s *serv) Create(ctx context.Context, c model.Category) (int64, error) {
	m := converter.ToCategoryFromService(c)
	if m.Slug == "" {
		gen := slugGenerator.New(slugGenerator.DefaultOptions())
		m.Slug = gen.Slag(c.Name)
	}
	return s.Repo.Create(ctx, &m)
}
