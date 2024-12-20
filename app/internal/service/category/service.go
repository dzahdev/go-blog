package service

import (
	"dzrise.ru/internal/client/db"
	"dzrise.ru/internal/repository"
)

type serv struct {
	Repo      repository.CategoryRepository
	TxManager db.TxManager
}

func New(repo repository.CategoryRepository, tm db.TxManager) *serv {
	return &serv{
		Repo:      repo,
		TxManager: tm,
	}
}
