package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ServerDownHistory struct {
	ID            uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	ServerID      uuid.UUID `gorm:"type:uuid" json:"server_id"`
	ServerName    string    `gorm:"not null" json:"server_name"`
	URL           string    `gorm:"not null" json:"url"`
	Status        string    `gorm:"not null" json:"status"` // DOWN, RESOLVED
	Timestamp     time.Time `json:"timestamp"`
	CreatedBy     uuid.UUID `gorm:"type:uuid" json:"created_by"`
	Description   string    `json:"description,omitempty"`
	ResolvedBy    *uuid.UUID `gorm:"type:uuid" json:"resolved_by,omitempty"`
	ResolvedAt    *time.Time `json:"resolved_at,omitempty"`
	ResolveNote   string     `json:"resolve_note,omitempty"`
}

// HistoryResponse represents history with user names instead of IDs
type HistoryResponse struct {
	ID            uuid.UUID `json:"id"`
	ServerID      uuid.UUID `json:"server_id"`
	ServerName    string    `json:"server_name"`
	URL           string    `json:"url"`
	Status        string    `json:"status"`
	Timestamp     time.Time `json:"timestamp"`
	CreatedBy     string    `json:"created_by"`     // User name instead of UUID
	ResolvedBy    *string   `json:"resolved_by"`    // User name instead of UUID
	ResolvedAt    *time.Time `json:"resolved_at,omitempty"`
	ResolveNote   string    `json:"resolve_note,omitempty"`
	Description   string    `json:"description,omitempty"`
}

// Hook: auto set UUID & timestamp
func (h *ServerDownHistory) BeforeCreate(tx *gorm.DB) (err error) {
	h.ID = uuid.New()
	h.Timestamp = time.Now()
	return
}
