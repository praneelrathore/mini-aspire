package time

import (
	"errors"
	"time"
)

// date time formats
const (
	DateFormat  = "2006-01-02"
	ISTLocation = "Asia/Kolkata"
)

// GetCurrentDate returns the current date in IST timezone
func (s *Service) GetCurrentDate() (time.Time, error) {
	location, err := time.LoadLocation(ISTLocation)
	if err != nil {
		return time.Time{}, errors.New("could not load location, err: " + err.Error())
	}

	currentTime := time.Now().In(location)
	return time.Date(
		currentTime.Year(),
		currentTime.Month(),
		currentTime.Day(),
		0,
		0,
		0,
		0,
		currentTime.Location(),
	), nil
}
