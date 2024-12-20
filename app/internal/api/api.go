package api

import (
	htmlH "dzrise.ru/internal/api/html"
	"dzrise.ru/internal/api/json"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func New(h *htmlH.HtmlHandlers, j *json.JSONHandlers) *fiber.App {
	engine := html.New("./views", ".html")

	srv := fiber.New(
		fiber.Config{
			Views:        engine,
			ErrorHandler: ErrorHandler,
		})

	//static files
	srv.Static("/", "./public")

	SetHtmlRoutes(srv, h) //html endpoints
	SetJsonRoutes(srv, j) //json endpoints

	return srv
}

func SetHtmlRoutes(a *fiber.App, h *htmlH.HtmlHandlers) {
	a.Get("/", h.Index)
}

func SetJsonRoutes(a *fiber.App, j *json.JSONHandlers) {
	v1 := a.Group("/api/v1")
	post := v1.Group("/post")

	post.Post("/", j.PostCreate)
	post.Get("/:id", j.GetById)
}

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	// Status code defaults to 500
	code := fiber.StatusInternalServerError

	// Retrieve the custom status code if it's a *fiber.Error
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	// Send custom error page
	err = ctx.Status(code).SendFile(fmt.Sprintf("./views/errors/%d.html", code))
	if err != nil {
		// In case the SendFile fails
		return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	// Return from handler
	return nil
}
