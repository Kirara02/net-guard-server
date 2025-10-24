package repository

import (
	"NetGuardServer/config"
	"NetGuardServer/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ServerRepository defines the interface for server data operations
type ServerRepository interface {
	Create(server *models.Server) error
	FindByID(id uuid.UUID) (*models.Server, error)
	GetAllServers() ([]models.Server, error)
	FindByUserID(userID uuid.UUID) ([]models.Server, error)
	Update(server *models.Server) error
	Delete(id uuid.UUID) error
}

// serverRepository implements ServerRepository
type serverRepository struct {
	db *gorm.DB
}

// NewServerRepository creates a new server repository instance
func NewServerRepository() ServerRepository {
	return &serverRepository{
		db: config.AppConfig.DB,
	}
}

// Create creates a new server
func (r *serverRepository) Create(server *models.Server) error {
	return r.db.Create(server).Error
}

// GetAllServers gets all servers from all users
func (r *serverRepository) GetAllServers() ([]models.Server, error) {
	var servers []models.Server
	err := r.db.Find(&servers).Error
	return servers, err
}

// FindByID finds a server by ID
func (r *serverRepository) FindByID(id uuid.UUID) (*models.Server, error) {
	var server models.Server
	err := r.db.Where("id = ?", id).First(&server).Error
	if err != nil {
		return nil, err
	}
	return &server, nil
}

// FindByUserID finds all servers by user ID
func (r *serverRepository) FindByUserID(userID uuid.UUID) ([]models.Server, error) {
	var servers []models.Server
	err := r.db.Where("created_by = ?", userID).Find(&servers).Error
	return servers, err
}

// Update updates a server
func (r *serverRepository) Update(server *models.Server) error {
	return r.db.Save(server).Error
}

// Delete deletes a server by ID
func (r *serverRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.Server{}, id).Error
}