package post

import (
	"dzrise.ru/internal/client/db"
	"dzrise.ru/internal/repository"
	"dzrise.ru/internal/service"
)

var _ service.PostService = (*serv)(nil)

type serv struct {
	Repo    repository.PostRepository
	RepoCat repository.CategoryRepository
	Tm      db.TxManager
}

func New(repo repository.PostRepository, repoCat repository.CategoryRepository, tm db.TxManager) *serv {
	return &serv{
		Repo:    repo,
		RepoCat: repoCat,
		Tm:      tm,
	}
}
