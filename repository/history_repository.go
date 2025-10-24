package repository

import (
	"NetGuardServer/config"
	"NetGuardServer/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// HistoryRepository defines the interface for history data operations
type HistoryRepository interface {
	Create(history *models.ServerDownHistory) error
	FindByID(id uuid.UUID) (*models.ServerDownHistory, error)
	FindByServerID(serverID uuid.UUID) ([]models.ServerDownHistory, error)
	FindAll(limit int) ([]models.ServerDownHistory, error)
	Update(history *models.ServerDownHistory) error
	GetMonthlyReport(year, month int) ([]map[string]interface{}, error)
}

// historyRepository implements HistoryRepository
type historyRepository struct {
	db *gorm.DB
}

// NewHistoryRepository creates a new history repository instance
func NewHistoryRepository() HistoryRepository {
	return &historyRepository{
		db: config.AppConfig.DB,
	}
}

// Create creates a new history record
func (r *historyRepository) Create(history *models.ServerDownHistory) error {
	return r.db.Create(history).Error
}

// FindByID finds a history record by ID
func (r *historyRepository) FindByID(id uuid.UUID) (*models.ServerDownHistory, error) {
	var history models.ServerDownHistory
	err := r.db.Where("id = ?", id).First(&history).Error
	if err != nil {
		return nil, err
	}
	return &history, nil
}

// Update updates a history record
func (r *historyRepository) Update(history *models.ServerDownHistory) error {
	return r.db.Save(history).Error
}

// GetMonthlyReport gets monthly server down statistics
func (r *historyRepository) GetMonthlyReport(year, month int) ([]map[string]interface{}, error) {
	var results []map[string]interface{}

	query := `
		SELECT
			server_id,
			server_name,
			url,
			COUNT(*) as down_count,
			COUNT(CASE WHEN resolved_at IS NOT NULL THEN 1 END) as resolved_count,
			AVG(EXTRACT(EPOCH FROM (COALESCE(resolved_at, CURRENT_TIMESTAMP) - timestamp))) as avg_resolution_time
		FROM server_down_histories
		WHERE EXTRACT(YEAR FROM timestamp) = $1
		AND EXTRACT(MONTH FROM timestamp) = $2
		AND status = 'DOWN'
		GROUP BY server_id, server_name, url
		ORDER BY down_count DESC
	`

	err := r.db.Raw(query, year, month).Scan(&results).Error
	return results, err
}

// FindByServerID finds history records by server ID
func (r *historyRepository) FindByServerID(serverID uuid.UUID) ([]models.ServerDownHistory, error) {
	var histories []models.ServerDownHistory
	err := r.db.Where("server_id = ?", serverID).Order("timestamp DESC").Find(&histories).Error
	return histories, err
}

// FindAll finds all history records with limit
func (r *historyRepository) FindAll(limit int) ([]models.ServerDownHistory, error) {
	var histories []models.ServerDownHistory
	err := r.db.Order("timestamp DESC").Limit(limit).Find(&histories).Error
	return histories, err
}