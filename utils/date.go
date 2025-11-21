package utils

import (
	"regexp"
	"time"
)

func IsValidDate(date string) bool {

	// check format
	pattern := regexp.MustCompile(`^\d{2}-\d{2}-\d{4}$`)
	if !pattern.MatchString(date) {
		return false
	}

	// validate date
	_, err := time.Parse("01-02-2006", date)
	return err == nil
}
