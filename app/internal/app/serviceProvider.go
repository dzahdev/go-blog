package app

import (
	"context"
	"dzrise.ru/internal/api"
	"dzrise.ru/internal/api/html"
	"dzrise.ru/internal/api/json"
	"dzrise.ru/internal/client/db"
	"dzrise.ru/internal/client/db/pg"
	"dzrise.ru/internal/config"
	"dzrise.ru/internal/pkg/closer"
	"dzrise.ru/internal/pkg/logger"
	"dzrise.ru/internal/pkg/transaction"
	"dzrise.ru/internal/repository"
	"dzrise.ru/internal/service"
	"github.com/gofiber/fiber/v2"
	"log"
	"log/slog"

	categoryRepository "dzrise.ru/internal/repository/category"
	commentRepository "dzrise.ru/internal/repository/comment"
	postRepository "dzrise.ru/internal/repository/post"

	categoryService "dzrise.ru/internal/service/category"
	postService "dzrise.ru/internal/service/post"
)

type ServiceProvider struct {
	cnf *config.Config
	log *slog.Logger

	dbClient           db.Client
	transactionManager db.TxManager

	server       *fiber.App
	HtmlHandlers *html.HtmlHandlers
	JSONHandlers *json.JSONHandlers

	PostRepo     repository.PostRepository
	CommentRepo  repository.CommentRepository
	CategoryRepo repository.CategoryRepository

	postService     service.PostService
	categoryService service.CategoryService
}

func NewServiceProvider() *ServiceProvider {
	return &ServiceProvider{}
}

func (sp *ServiceProvider) Config() *config.Config {
	if sp.cnf == nil {
		sp.cnf = config.MustLoad()
	}
	return sp.cnf
}

func (sp *ServiceProvider) Logger() *slog.Logger {
	if sp.log == nil {
		sp.log = logger.SetupLogger(sp.Config().Env)
	}
	return sp.log
}

func (sp *ServiceProvider) DBClient(ctx context.Context) db.Client {
	if sp.dbClient == nil {
		cl, err := pg.New(ctx, sp.Config().DSN())
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("failed to ping db: %v", err)
		}
		closer.Append(cl.Close)

		sp.dbClient = cl

	}

	return sp.dbClient
}

func (sp *ServiceProvider) TransactionManager(ctx context.Context) db.TxManager {
	if sp.transactionManager == nil {
		sp.transactionManager = transaction.NewManager(sp.DBClient(ctx).DB(), "tx")
	}
	return sp.transactionManager
}

func (sp *ServiceProvider) PostRepository(ctx context.Context) repository.PostRepository {
	if sp.PostRepo == nil {
		sp.PostRepo = postRepository.New(sp.DBClient(ctx), sp.TransactionManager(ctx))
	}
	return sp.PostRepo
}

func (sp *ServiceProvider) CommentRepository(ctx context.Context) repository.CommentRepository {
	if sp.CommentRepo == nil {
		sp.CommentRepo = commentRepository.New(sp.DBClient(ctx), sp.TransactionManager(ctx))
	}
	return sp.CommentRepo
}

func (sp *ServiceProvider) CategoryRepository(ctx context.Context) repository.CategoryRepository {
	if sp.CategoryRepo == nil {
		sp.CategoryRepo = categoryRepository.New(sp.DBClient(ctx), sp.TransactionManager(ctx))
	}
	return sp.CategoryRepo
}

func (sp *ServiceProvider) PostService(ctx context.Context) service.PostService {
	if sp.postService == nil {
		sp.postService = postService.New(sp.PostRepository(ctx), sp.CategoryRepository(ctx), sp.TransactionManager(ctx))
	}

	return sp.postService

}

func (sp *ServiceProvider) CategoryService(ctx context.Context) service.CategoryService {
	if sp.categoryService == nil {
		sp.categoryService = categoryService.New(sp.CategoryRepository(ctx), sp.TransactionManager(ctx))
	}
}

func (sp *ServiceProvider) GetServer(ctx context.Context) *fiber.App {
	if sp.server == nil {
		sp.server = api.New(sp.GetHtmlHandlers(), sp.GetJsonHandlers(ctx))
	}

	return sp.server
}

func (sp *ServiceProvider) GetHtmlHandlers() *html.HtmlHandlers {
	if sp.HtmlHandlers == nil {
		sp.HtmlHandlers = html.NewHtmlHandlers()
	}
	return sp.HtmlHandlers
}

func (sp *ServiceProvider) GetJsonHandlers(ctx context.Context) *json.JSONHandlers {
	if sp.JSONHandlers == nil {
		sp.JSONHandlers = json.NewJSONHandlers(sp.PostService(ctx))
	}
	return sp.JSONHandlers
}
