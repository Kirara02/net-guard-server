package dto

// CreateServerRequest represents create server request
type CreateServerRequest struct {
	Name string `json:"name" validate:"required,min=1,max=255"`
	URL  string `json:"url" validate:"required,url"`
}

// UpdateServerRequest represents update server request
type UpdateServerRequest struct {
	Name string `json:"name,omitempty" validate:"omitempty,min=1,max=255"`
	URL  string `json:"url,omitempty" validate:"omitempty,url"`
}

// UpdateServerStatusRequest represents update server status request
type UpdateServerStatusRequest struct {
	Status       string `json:"status" validate:"required,oneof=UP DOWN UNKNOWN"`
	ResponseTime int64  `json:"response_time,omitempty"`
}

// ServerDTO represents server data transfer object
type ServerDTO struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	URL        string `json:"url"`
	Status     string `json:"status,omitempty"`
	ResponseTime int64 `json:"response_time,omitempty"`
	LastChecked string `json:"last_checked,omitempty"`
	CreatedBy  string `json:"created_by"`
	CreatedAt  string `json:"created_at"`
}