package utils_test

import (
	"testing"
	"time"

	"github.com/Sourceware-Lab/realquick/internal/utils"
)

func TestParseDate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		input     string
		expected  time.Time
		expectErr bool
	}{
		{
			name:      "valid_ISO_format",
			input:     "2023-10-05",
			expected:  time.Date(2023, 10, 5, 0, 0, 0, 0, time.UTC),
			expectErr: false,
		},
		{
			name:      "valid_date_with_time",
			input:     "2023-10-05T14:30:00",
			expectErr: true,
		},
		{
			name:      "invalid_format",
			input:     "05-10-2023",
			expected:  time.Time{},
			expectErr: true,
		},
		{
			name:      "empty_input",
			input:     "",
			expected:  time.Time{},
			expectErr: true,
		},
		{
			name:      "out_of_range_date",
			input:     "2023-13-32",
			expected:  time.Time{},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result, err := utils.ParseDate(tt.input)
			if (err != nil) != tt.expectErr {
				t.Fatalf("expected error: %v, got: %v", tt.expectErr, err)
			}

			if !tt.expectErr && !result.Equal(tt.expected) {
				t.Errorf("expected: %v, got: %v", tt.expected, result)
			}
		})
	}
}
