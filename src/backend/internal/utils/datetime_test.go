package utils_test

import (
	"testing"
	"time"

	"github.com/Sourceware-Lab/realquick/internal/utils"
)

//nolint:funlen
func TestParseAnyDatetime(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    string
		expected time.Time
		wantErr  bool
	}{
		{
			name:     "ValidISO8601",
			input:    "2023-10-27T15:04:05Z",
			expected: time.Date(2023, 10, 27, 15, 4, 5, 0, time.UTC),
			wantErr:  false,
		},
		{
			name:     "ValidISO8601 +05:00 TZ",
			input:    "2023-10-27T15:04:05+05:00",
			expected: time.Date(2023, 10, 27, 10, 4, 5, 0, time.UTC),
			wantErr:  false,
		},
		{
			name:     "ValidISO8601",
			input:    "2023-10-27T00:00:00Z",
			expected: time.Date(2023, 10, 27, 0, 0, 0, 0, time.UTC),
			wantErr:  false,
		},
		{
			name:     "ValidISO8601 -02:00 TZ",
			input:    "2023-10-27T15:04:05-02:00",
			expected: time.Date(2023, 10, 27, 17, 4, 5, 0, time.UTC),
			wantErr:  false,
		},
		{
			name:    "ValidRFC1123",
			input:   "Mon, 02 Jan 2006 15:04:05 MST",
			wantErr: true,
		},
		{
			name:    "ValidYYYYMMDD",
			input:   "2023-10-27",
			wantErr: true,
		},
		{
			name:    "EmptyString",
			input:   "",
			wantErr: true,
		},
		{
			name:    "InvalidFormat",
			input:   "27-10-2023",
			wantErr: true,
		},
		{
			name:    "GarbageString",
			input:   "random_string",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := utils.ParseDatetime(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseDatetime() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if !tt.wantErr && !got.Equal(tt.expected) {
				t.Errorf("ParseDatetime() got = %v, expected %v", got, tt.expected)
			}
		})
	}
}
