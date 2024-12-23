package types_test

import (
	"testing"
	"time"

	"github.com/Sourceware-Lab/realquick/internal/types"
)

func TestParseTime(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		input   string
		want    time.Time
		wantErr bool
	}{
		{
			name:    "valid time format",
			input:   "15:04:05",
			want:    time.Date(0, 1, 1, 15, 4, 5, 0, time.UTC),
			wantErr: false,
		},
		{
			name:    "invalid time format",
			input:   "15-04-05",
			want:    time.Time{},
			wantErr: true,
		},
		{
			name:    "empty input",
			input:   "",
			want:    time.Time{},
			wantErr: true,
		},
		{
			name:    "invalid characters",
			input:   "invalid",
			want:    time.Time{},
			wantErr: true,
		},
		{
			name:    "partial time format",
			input:   "15:04",
			want:    time.Time{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := types.Time{}.Parse(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseTime() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr && !got.Equal(tt.want) {
				t.Errorf("ParseTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimeMarshalJSON(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		input   types.Time
		want    string
		wantErr bool
	}{
		{
			name:    "valid time",
			input:   types.Time{Time: time.Date(0, 1, 1, 15, 4, 5, 0, time.UTC)},
			want:    `"15:04:05"`,
			wantErr: false,
		},
		{
			name:    "zero time",
			input:   types.Time{Time: time.Time{}},
			want:    `"00:00:00"`,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := tt.input.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}

			if string(got) != tt.want {
				t.Errorf("MarshalJSON() = %s, want %s", got, tt.want)
			}
		})
	}
}

func TestTimeUnmarshalJSON(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		input   string
		want    types.Time
		wantErr bool
	}{
		{
			name:    "valid time",
			input:   `"15:04:05"`,
			want:    types.Time{Time: time.Date(0, 1, 1, 15, 4, 5, 0, time.UTC)},
			wantErr: false,
		},
		{
			name:    "zero time",
			input:   `"00:00:00"`,
			want:    types.Time{Time: time.Date(0, 1, 1, 0, 0, 0, 0, time.UTC)},
			wantErr: false,
		},
		{
			name:    "invalid JSON format",
			input:   `15:04:05`, // Missing quotes
			want:    types.Time{},
			wantErr: true,
		},
		{
			name:    "invalid time format",
			input:   `"15-04-05"`,
			want:    types.Time{},
			wantErr: true,
		},
		{
			name:    "unexpected input",
			input:   `"invalid"`,
			want:    types.Time{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var got types.Time
			err := got.UnmarshalJSON([]byte(tt.input))
			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr && !got.Equal(tt.want.Time) {
				t.Errorf("UnmarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
