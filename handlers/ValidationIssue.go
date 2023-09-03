package handlers

type ValidationIssue struct {
    Error   string `json:"error"`
    Message string `json:"message,omitempty"`
}