package json

import (
	"dzrise.ru/internal/model"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func (h *JSONHandlers) PostCreate(c *fiber.Ctx) error {
	post := new(model.Post)
	err := c.BodyParser(&post)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	id, err := h.postService.Create(c.UserContext(), post)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"id": id,
	})
}

func (h *JSONHandlers) GetById(c *fiber.Ctx) error {
	idS := c.Params("id")
	id, err := strconv.ParseInt(idS, 10, 64)

	if err != nil || id <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	post, err := h.postService.GetByID(c.UserContext(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(post)
}
