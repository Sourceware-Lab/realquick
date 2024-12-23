package types

import (
	"errors"
	"fmt"
	"time"
)

type Time struct {
	time.Time
}

func (t Time) Parse(s string) (Time, error) { // TODO should  I pass a * and just mutate the instance?
	parsedTime, err := time.Parse(time.TimeOnly, s)
	if err != nil {
		return Time{}, fmt.Errorf("time does not follow 15:04:05 time only format: %w", err)
	}

	return Time{Time: parsedTime}, nil

}

// MarshalJSON implements the [json.Marshaler] interface.
// The time is a quoted string in the hh:mm:ss format
func (t Time) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(time.TimeOnly)+len(`""`))
	b = append(b, '"')
	formatedTime := t.Format(time.TimeOnly)
	b = append(b, formatedTime...)
	b = append(b, '"')

	return b, nil
}

// UnmarshalJSON implements the [json.Unmarshaler] interface.
// The time must be a quoted string in the RFC 3339 format.
func (t *Time) UnmarshalJSON(data []byte) error {
	tempTime := time.Time{}
	parsedTime := Time{}

	err := tempTime.UnmarshalJSON(data)
	if err != nil {
		if string(data) == "null" {
			return nil
		}
		// TODO(https://go.dev/issue/47353): Properly unescape a JSON string.
		if len(data) < 2 || data[0] != '"' || data[len(data)-1] != '"' {
			return errors.New("Time.UnmarshalJSON: input is not a JSON string")
		}
		data = data[len(`"`) : len(data)-len(`"`)]
		var err error
		parsedTime, err = t.Parse(string(data))
		if err != nil {
			return err
		}
	} else {
		parsedTime, err = Time{}.Parse(tempTime.Format(time.TimeOnly))
		if err != nil {
			return fmt.Errorf("failed to format unmarshaled time: %w", err)
		}
	}
	*t = parsedTime
	return nil
}
