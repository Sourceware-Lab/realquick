package utils

import (
	"fmt"
	"time"
)

var DateFormats = []string{ //nolint:gochecknoglobals
	time.DateOnly,
}

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
