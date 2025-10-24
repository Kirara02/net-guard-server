package services

import (
	"NetGuardServer/models"
	"NetGuardServer/repository"
	"errors"
	"time"

	"github.com/google/uuid"
)

// HistoryService defines the interface for history business logic
type HistoryService interface {
	CreateHistory(serverID uuid.UUID, serverName, url, status string, createdBy uuid.UUID) (*models.ServerDownHistory, error)
	GetHistoryByID(id uuid.UUID) (*models.ServerDownHistory, error)
	GetHistoryByServerID(serverID uuid.UUID) ([]models.HistoryResponse, error)
	GetAllHistory(limit int) ([]models.HistoryResponse, error)
	ResolveHistory(id uuid.UUID, resolvedBy uuid.UUID, resolveNote string) (*models.ServerDownHistory, error)
	GetMonthlyReport(year, month int) ([]map[string]interface{}, error)
}

// historyService implements HistoryService
type historyService struct {
	historyRepo repository.HistoryRepository
}

// NewHistoryService creates a new history service instance
func NewHistoryService(historyRepo repository.HistoryRepository) HistoryService {
	return &historyService{
		historyRepo: historyRepo,
	}
}

// CreateHistory handles history record creation business logic
func (s *historyService) CreateHistory(serverID uuid.UUID, serverName, url, status string, createdBy uuid.UUID) (*models.ServerDownHistory, error) {
	history := &models.ServerDownHistory{
		ServerID:   serverID,
		ServerName: serverName,
		URL:        url,
		Status:     status,
		CreatedBy:  createdBy,
	}

	if err := s.historyRepo.Create(history); err != nil {
		return nil, errors.New("failed to create history record")
	}

	return history, nil
}

// GetHistoryByID gets a history record by ID
func (s *historyService) GetHistoryByID(id uuid.UUID) (*models.ServerDownHistory, error) {
	history, err := s.historyRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("history record not found")
	}
	return history, nil
}

// ResolveHistory resolves a history record
func (s *historyService) ResolveHistory(id uuid.UUID, resolvedBy uuid.UUID, resolveNote string) (*models.ServerDownHistory, error) {
	history, err := s.historyRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("history record not found")
	}

	// Check if already resolved
	if history.ResolvedAt != nil {
		return nil, errors.New("history record already resolved")
	}

	now := time.Now()
	history.Status = "RESOLVED"
	history.ResolvedBy = &resolvedBy
	history.ResolvedAt = &now
	history.ResolveNote = resolveNote

	if err := s.historyRepo.Update(history); err != nil {
		return nil, errors.New("failed to resolve history record")
	}

	return history, nil
}

// GetMonthlyReport gets monthly server down statistics
func (s *historyService) GetMonthlyReport(year, month int) ([]map[string]interface{}, error) {
	results, err := s.historyRepo.GetMonthlyReport(year, month)
	if err != nil {
		return nil, errors.New("failed to get monthly report")
	}
	return results, nil
}

// GetHistoryByServerID gets history records for a specific server
func (s *historyService) GetHistoryByServerID(serverID uuid.UUID) ([]models.HistoryResponse, error) {
	histories, err := s.historyRepo.FindByServerID(serverID)
	if err != nil {
		return nil, errors.New("failed to get history records")
	}
	return s.convertToHistoryResponse(histories), nil
}

// GetAllHistory gets all history records with limit
func (s *historyService) GetAllHistory(limit int) ([]models.HistoryResponse, error) {
	histories, err := s.historyRepo.FindAll(limit)
	if err != nil {
		return nil, errors.New("failed to get history records")
	}
	return s.convertToHistoryResponse(histories), nil
}

// convertToHistoryResponse converts ServerDownHistory slice to HistoryResponse slice with user names
func (s *historyService) convertToHistoryResponse(histories []models.ServerDownHistory) []models.HistoryResponse {
	responses := make([]models.HistoryResponse, len(histories))

	// Create a map to cache user names to avoid multiple DB queries
	userCache := make(map[uuid.UUID]string)

	for i, history := range histories {
		// Get created_by user name
		createdByName, err := s.getUserName(history.CreatedBy, userCache)
		if err != nil {
			createdByName = "Unknown User"
		}

		// Get resolved_by user name (if exists)
		var resolvedByName *string
		if history.ResolvedBy != nil {
			name, err := s.getUserName(*history.ResolvedBy, userCache)
			if err != nil {
				name = "Unknown User"
			}
			resolvedByName = &name
		}

		responses[i] = models.HistoryResponse{
			ID:          history.ID,
			ServerID:    history.ServerID,
			ServerName:  history.ServerName,
			URL:         history.URL,
			Status:      history.Status,
			Timestamp:   history.Timestamp,
			CreatedBy:   createdByName,
			ResolvedBy:  resolvedByName,
			ResolvedAt:  history.ResolvedAt,
			ResolveNote: history.ResolveNote,
			Description: history.Description,
		}
	}

	return responses
}

// getUserName gets user name by ID with caching
func (s *historyService) getUserName(userID uuid.UUID, cache map[uuid.UUID]string) (string, error) {
	// Check cache first
	if name, exists := cache[userID]; exists {
		return name, nil
	}

	// Query from database
	userRepo := repository.NewUserRepository()
	user, err := userRepo.FindByID(userID)
	if err != nil {
		return "", err
	}

	// Cache the result
	cache[userID] = user.Name
	return user.Name, nil
}