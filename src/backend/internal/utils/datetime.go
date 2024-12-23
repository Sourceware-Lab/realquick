package utils

import (
	"fmt"
	"time"
)

var DatetimeFormats = []string{ //nolint:gochecknoglobals
	time.RFC3339,
	time.RFC3339Nano,
}

// ParseDatetime attempts to parse a datetime string using predefined layouts and returns the parsed time or an error.
// It iterates through supported datetime formats (e.g., RFC3339, RFC3339Nano) until a match is found or fails.
func ParseDatetime(s string) (time.Time, error) {
	parsedTime := time.Time{}

	var err error

	for _, layout := range DatetimeFormats {
		if parsedTime, err = time.Parse(layout, s); err == nil {
			return parsedTime, nil
		}
	}

	return parsedTime, fmt.Errorf("datetime does not follow RFC3339: %w", err)
}
