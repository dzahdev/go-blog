package html

import "github.com/gofiber/fiber/v2"

type HtmlHandlers struct {
}

func NewHtmlHandlers() *HtmlHandlers {
	return &HtmlHandlers{}
}

func (h *HtmlHandlers) Index(c *fiber.Ctx) error {
	data := fiber.Map{
		"title":       "Denis Zakharov!",
		"description": "This is a blog about programming by Denis Zakharov.",
	}
	return c.Render("index", data, "layouts/main")
}
