package post

import "context"

func (s *serv) Delete(ctx context.Context, id int64) error {
	return s.Repo.Delete(ctx, id)
}
