package pgmodels_test

import (
	"errors"
	"testing"

	pgmodels "github.com/Sourceware-Lab/realquick/database/postgres/models"
)

//nolint:funlen
func TestTag_Verify(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		tag       pgmodels.Tag
		wantError error
	}{
		{
			name:      "valid tag",
			tag:       pgmodels.Tag{Name: "TagName", Color: "#ff5733"},
			wantError: nil,
		},

		// Name
		{
			name:      "missing name",
			tag:       pgmodels.Tag{Name: "", Color: "#ff5733"},
			wantError: pgmodels.ErrMissingName,
		},
		{
			name:      "name with only spaces",
			tag:       pgmodels.Tag{Name: "   ", Color: "#ff5733"},
			wantError: pgmodels.ErrMissingName,
		},

		// Color
		{
			name:      "color invalid hex",
			tag:       pgmodels.Tag{Name: "TagName", Color: "123456"},
			wantError: pgmodels.ErrInvalidColor,
		},
		{
			name:      "color empty",
			tag:       pgmodels.Tag{Name: "TagName", Color: ""},
			wantError: pgmodels.ErrInvalidColor,
		},
		{
			name:      "color uppercase valid hex",
			tag:       pgmodels.Tag{Name: "TagName", Color: "#ABCDEF"},
			wantError: nil,
		},
		{
			name:      "color missing # ",
			tag:       pgmodels.Tag{Name: "TagName", Color: "123abc"},
			wantError: pgmodels.ErrInvalidColor,
		},
		{
			name:      "color too short",
			tag:       pgmodels.Tag{Name: "TagName", Color: "#123ab"},
			wantError: pgmodels.ErrInvalidColor,
		},
		{
			name:      "color too long ",
			tag:       pgmodels.Tag{Name: "TagName", Color: "#123abcd"},
			wantError: pgmodels.ErrInvalidColor,
		},
		{
			name:      "color invalid hex char ",
			tag:       pgmodels.Tag{Name: "TagName", Color: "#123abz"},
			wantError: pgmodels.ErrInvalidColor,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := tt.tag.Verify()
			if !errors.Is(err, tt.wantError) {
				t.Errorf("Verify() error = %v, wantErr %v", err, tt.wantError)
			}
		})
	}
}
