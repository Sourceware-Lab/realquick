package timeblockapi

import (
	"time"

	"github.com/danielgtaylor/huma/v2"
	"gorm.io/gorm"

	pgmodels "github.com/Sourceware-Lab/realquick/database/postgres/models"
)

type TimeblockGetInput struct {
	ID uint
}

type TimeblockGetOutput struct {
	Body TimeblockPostBodyInput
}

type TimeblockPostInput struct {
	Body TimeblockPostBodyInput `json:"body"`
}

type TimeblockPostOutput struct {
	Body struct {
		ID uint `doc:"Id for new user" example:"999" json:"id"`
	}
}
type TimeblockPostBodyInput struct {
	pgmodels.TimeBlock

	// fields to ignore. `json:"-" will instruct huma and json to ignore the field.
	ID        uint           `json:"-"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

func (i *TimeblockPostBodyInput) TableName() string {
	return "time_blocks"
}

func (i *TimeblockPostBodyInput) Resolve(_ huma.Context) []error {
	if i.Recur && i.Days == nil {
		return []error{&huma.ErrorDetail{
			Location: "TimeblockPostBodyInput.days",
			Message:  "If recur is true, days must be set",
			Value:    i.Days,
		}}
	}

	if !i.Recur && i.Days != nil {
		return []error{&huma.ErrorDetail{
			Location: "TimeblockPostBodyInput.days",
			Message:  "If recur is false, days must not be set",
			Value:    i.Days,
		}}
	}

	if i.EndDate != nil && i.StartDate.After(i.EndDate.Time) {
		return []error{&huma.ErrorDetail{
			Location: "TimeblockPostBodyInput.startDate",
			Message:  "Start date must be before end date",
			Value:    i.StartDate,
		}}
	}

	return nil
}
