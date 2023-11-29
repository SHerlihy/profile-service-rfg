package routes

import (
	"github.com/SHerlihy/profile-service-rfg/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(router fiber.Router) {
	router.Post("/profile/create", controllers.Create)
	router.Post("/profile/update-access", controllers.UpdateAccess)
}
