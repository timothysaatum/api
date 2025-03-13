package routes

import (
	"api/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	referral := app.Group("/api/referrals")
	referral.Get("/", handlers.GetReferrals)
	referral.Get("/:id", handlers.GetReferral)
}
