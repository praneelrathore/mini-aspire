package time

import "time"

type ITime interface {
	// GetCurrentDate returns the current date in IST timezone
	GetCurrentDate() (time.Time, error)
}

// Service is a struct for time service object
type Service struct{}

// NewService returns a new instance of the time service
func NewService() *Service {
	return &Service{}
}
