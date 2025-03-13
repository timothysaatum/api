package handlers

import (
	"api/services"
	"github.com/gofiber/fiber/v2"
)

// GetReferrals handles fetching referrals
func GetReferrals(c *fiber.Ctx) error {
	// Example: Fetching query params
	status := c.Query("status")
	fromDate := c.Query("from_date")
	toDate := c.Query("to_date")

	// Define filters map
	filters := make(map[string]interface{})
	if status != "" {
		filters["referral_status"] = status
	}
	if fromDate != "" && toDate != "" {
		filters["date_referred BETWEEN ? AND ?"] = []interface{}{fromDate, toDate}
	}

	// Fetch referrals from service
	referrals, err := services.FetchReferrals(filters)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch referrals"})
	}

	return c.JSON(referrals)
}

// GetReferral handles fetching a single referral by ID
func GetReferral(c *fiber.Ctx) error {
	id := c.Params("id")

	referral, err := services.FetchReferral(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Referral not found"})
	}

	return c.JSON(referral)
}
