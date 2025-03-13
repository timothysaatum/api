package repositories

import (
	"api/database"
	"api/models"
)

// GetReferrals retrieves referrals based on filtering
func GetReferrals(filters map[string]interface{}) ([]models.Referral, error) {
	var referrals []models.Referral
	query := database.DB.Where(filters).Find(&referrals)
	return referrals, query.Error
}

// GetReferralByID retrieves a single referral by ID
func GetReferralByID(id string) (models.Referral, error) {
	var referral models.Referral
	query := database.DB.First(&referral, "id = ?", id)
	return referral, query.Error
}
