package handlers

import (
    "github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
    code := fiber.StatusInternalServerError
    if e, ok := err.(*fiber.Error); ok {
        code = e.Code
    }

    return c.Status(code).JSON(fiber.Map{
        "error": err.Error(),
    })
}
