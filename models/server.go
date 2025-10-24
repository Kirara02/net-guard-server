package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Server struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name      string    `gorm:"not null" json:"name"`
	URL       string    `gorm:"not null" json:"url"`
	CreatedBy uuid.UUID `gorm:"type:uuid" json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
}

// Auto generate UUID & timestamp
func (s *Server) BeforeCreate(tx *gorm.DB) (err error) {
	s.ID = uuid.New()
	s.CreatedAt = time.Now()
	return
}
