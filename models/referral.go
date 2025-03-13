package models

import (
	"github.com/google/uuid"
	"time"
	"gorm.io/gorm"
)

// Referral struct maps to the referral table in MySQL
type Referral struct {
	ID                  uuid.UUID `gorm:"type:char(36);primaryKey"`
	ReferralID          string    `gorm:"size:50;uniqueIndex"`
	ReferringFacilityID uint
	FacilityType        string    `gorm:"size:255"`
	PatientName         string    `gorm:"size:200"`
	PatientAge          string    `gorm:"size:10"`
	PatientSex          string    `gorm:"size:20"`
	ClinicalHistory     string    `gorm:"type:text"`
	ToLaboratoryID      uint
	DeliveryID          *uint
	ReferralAttachment  string    `gorm:"size:500"`
	Attachment          string    `gorm:"size:500"`
	RequiresPhlebotomist bool
	SenderFullName      string    `gorm:"size:200"`
	SenderPhone         string    `gorm:"size:20"`
	SenderEmail         string    `gorm:"size:100"`
	ReferralStatus      string    `gorm:"size:100"`
	DateReferred        time.Time `gorm:"index"`
	IsCompleted         bool      `gorm:"index"`
	IsArchived          bool      `gorm:"index"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
}

// BeforeCreate generates a UUID before saving the referral
func (r *Referral) BeforeCreate(tx *gorm.DB) error {
	r.ID = uuid.New()
	return nil
}
