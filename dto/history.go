package dto

// CreateHistoryRequest represents create history request
type CreateHistoryRequest struct {
	ServerID   string `json:"server_id" validate:"required,uuid4"`
	ServerName string `json:"server_name" validate:"required,min=1,max=255"`
	URL        string `json:"url" validate:"required,url"`
	Status     string `json:"status" validate:"required,oneof=DOWN"`
}

// ResolveHistoryRequest represents resolve history request
type ResolveHistoryRequest struct {
	ResolveNote string `json:"resolve_note" validate:"required,min=1,max=1000"`
}

// HistoryDTO represents history data transfer object
type HistoryDTO struct {
	ID            string  `json:"id"`
	ServerID      string  `json:"server_id"`
	ServerName    string  `json:"server_name"`
	URL           string  `json:"url"`
	Status        string  `json:"status"`
	Timestamp     string  `json:"timestamp"`
	CreatedBy     string  `json:"created_by"`
	ResolvedBy    *string `json:"resolved_by,omitempty"`
	ResolvedAt    *string `json:"resolved_at,omitempty"`
	ResolveNote   string  `json:"resolve_note,omitempty"`
	Description   string  `json:"description,omitempty"`
}

// MonthlyReportDTO represents monthly report data
type MonthlyReportDTO struct {
	Year   int                      `json:"year"`
	Month  int                      `json:"month"`
	Report []map[string]interface{} `json:"report"`
}