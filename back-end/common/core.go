package common

type APIResponse struct {
	Status  string      `json:"status"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message"`
}

// Status enum
type StatusEnum struct {
	Ok           string
	Error        string
	Invalid      string
	NotFound     string
	Unauthorized string
	Forbidden    string
	Existed      string
}

var APIStatus = &StatusEnum{
	Ok:           "OK",
	Error:        "ERROR",
	Invalid:      "INVALID",
	NotFound:     "NOT_FOUND",
	Unauthorized: "UNAUTHORIZED",
	Forbidden:    "FORBIDDEN",
	Existed:      "EXISTED",
}
