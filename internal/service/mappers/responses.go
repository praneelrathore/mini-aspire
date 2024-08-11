package mappers

// Status denotes the response status
type Status struct {
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}

// StatusSuccess is a helper function to create a successful Status object
func StatusSuccess(message string) Status {
	return Status{
		Status:  "success",
		Message: message,
	}
}

// StatusFailed is a helper function to create a unsuccessful Status object
func StatusFailed(message string) Status {
	return Status{
		Status:  "failed",
		Message: message,
	}
}
