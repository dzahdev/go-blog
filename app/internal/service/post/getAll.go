package post

import (
	"context"
	"dzrise.ru/internal/converter"
	"dzrise.ru/internal/model"
)

func (s *serv) GetAll(ctx context.Context) ([]*model.Post, error) {
	posts, err := s.Repo.GetAll(ctx)

	if err != nil {
		return nil, err
	}

	res := make([]*model.Post, 0, len(posts))

	for _, post := range posts {
		cat, err := s.RepoCat.GetByID(ctx, post.CategoryId)
		if err != nil {
			return nil, err
		}
		r := converter.ToPostFromRepository(post, cat)

		res = append(res, r)

	}
	return res, nil
}
