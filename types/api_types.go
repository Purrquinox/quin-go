package types

// Response is a generic response type used for API responses.
type Response struct {
	Message string `json:"message" description:"Message of the response"`
}

// ApiError represents an error response in the API.
type ApiError struct {
	Context map[string]string `json:"context,omitempty" description:"Context of the error. Usually used for validation error contexts"`
	Message string            `json:"message" description:"Message of the error"`
}
