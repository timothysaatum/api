package services

import (
	"api/models"
	"api/repositories"
)

// FetchReferrals retrieves referrals with filtering
func FetchReferrals(filters map[string]interface{}) ([]models.Referral, error) {
	return repositories.GetReferrals(filters)
}

// FetchReferral retrieves a single referral by ID
func FetchReferral(id string) (models.Referral, error) {
	return repositories.GetReferralByID(id)
}
