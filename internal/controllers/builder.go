package controllers

import "github.com/gofiber/fiber/v2"

type Builder interface {
	Build(c *fiber.App)
}
