package timeblockapi

import (
	"github.com/danielgtaylor/huma/v2"

	pgmodels "github.com/Sourceware-Lab/realquick/database/postgres/models"
)

type TimeblockGetInput struct {
	ID uint64 `doc:"Id for the timeblock you want to get" example:"999" path:"id"`
}

type TimeblockGetOutput struct {
	Body struct {
		TimeblockPostBodyInput
		ID uint64 `json:"id"`
	}
}

type TimeblockPostInput struct {
	Body TimeblockPostBodyInput `json:"body"`
}

type TimeblockPostOutput struct {
	Body struct {
		ID uint64 `doc:"Id for new user" example:"999" json:"id"`
	}
}
type TimeblockPostBodyInput struct {
	pgmodels.TimeBlock
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

	if i.StartTime.After(i.EndTime.Time.Time) {
		return []error{&huma.ErrorDetail{
			Location: "TimeblockPostBodyInput.startTime",
			Message:  "Start time must be before end time",
			Value:    i.StartTime,
		}}
	}

	return nil
}
