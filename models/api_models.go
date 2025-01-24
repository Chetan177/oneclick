package models

type HealthResponse struct {
	Status  string `json:"status"`
	Version string `json:"version"`
}

type ProjectRequest struct {
	Name        string `json:"name" validate:"required,min=3,max=100"`
	Description string `json:"description" validate:"required,min=10,max=500"`
}

type ProjectResponse struct {
	Projects []Project `json:"projects"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
