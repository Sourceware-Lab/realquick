package pgmodels_test

import (
	"errors"
	"testing"
	"time"

	pgmodels "github.com/Sourceware-Lab/realquick/database/postgres/models"
	"github.com/Sourceware-Lab/realquick/internal/utils"
)

func TestVerify(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		timeblock pgmodels.TimeBlock
		wantErr   bool
		errMsg    error
	}{
		{
			name: "missing name",
			timeblock: pgmodels.TimeBlock{
				Name: "",
			},
			wantErr: true,
			errMsg:  pgmodels.ErrMissingName,
		},
		{
			name: "missing days with recur true",
			timeblock: pgmodels.TimeBlock{
				Name:  "test",
				Recur: true,
			},
			wantErr: true,
			errMsg:  pgmodels.ErrMissingDays,
		},
		{
			name: "missing recur with days",
			timeblock: pgmodels.TimeBlock{
				Name:  "test",
				Recur: false,
				Days:  utils.MakePointer("0000000"),
			},
			wantErr: true,
			errMsg:  pgmodels.ErrMissingRecur,
		},
		{
			name: "start after end",
			timeblock: pgmodels.TimeBlock{
				Name:      "test",
				StartDate: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
				EndDate:   utils.MakePointer(time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)),
			},
			wantErr: true,
			errMsg:  pgmodels.ErrStartAfterEnd,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := tt.timeblock.Verify()
			if err != nil != tt.wantErr {
				t.Errorf("Verify() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err != nil && !errors.Is(err, tt.errMsg) {
				t.Errorf("Verify() error = %v, wantErr %v", err, tt.errMsg)
			}
		})
	}
}
