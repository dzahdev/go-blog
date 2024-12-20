package json

import (
	"dzrise.ru/internal/model"
	"github.com/gofiber/fiber/v2"
)

func (h *JSONHandlers) CategoryCreate(c *fiber.Ctx) error {
	category := new(model.Category)
	err := c.BodyParser(&category)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	id, err := h.categoryService.Create(c.UserContext(), category)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"id": id,
	})
}
