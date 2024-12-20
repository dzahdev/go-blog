package post

import (
	"context"
	"dzrise.ru/internal/converter"
	"dzrise.ru/internal/model"
	"fmt"
)

func (s *serv) Create(ctx context.Context, post *model.Post) (int64, error) {
	r := converter.ToPostFromService(post)
	fmt.Println(r)
	return s.Repo.Create(ctx, r)
}
