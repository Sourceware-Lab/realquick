package utils

import (
	"fmt"
	"time"
)

var DateFormats = []string{ //nolint:gochecknoglobals
	time.DateOnly,
}

// ParseDate parses a date string using predefined layouts and returns the parsed time.Time or an error
// if parsing fails.
// It iterates through supported date formats until a valid match is found or returns an error if none match.
func ParseDate(s string) (time.Time, error) {
	parsedTime := time.Time{}

	var err error

	for _, layout := range DateFormats {
		if parsedTime, err = time.Parse(layout, s); err == nil {
			return parsedTime, nil
		}
	}

	return parsedTime, fmt.Errorf("date does not follow YYYY-mm-dd: %w", err)
}
