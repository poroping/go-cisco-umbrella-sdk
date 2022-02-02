package models

type ErrorResp struct {
	StatusCode *int64               `json:"statusCode,omitempty"`
	Error      *string              `json:"error,omitempty"`
	Message    *string              `json:"message,omitempty"`
	Validation *ErrorRespValidation `json:"validation,omitempty"`
}

type ErrorRespValidation struct {
	Source *string  `json:"source,omitempty"`
	Keys   []string `json:"keys,omitempty"`
}
