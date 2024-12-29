package pgmodels

import (
	"errors"
	"time"

	"github.com/peterHoburg/go-date-and-time-extension/dtegorm"
	"gorm.io/gorm"
)

var (
	ErrMissingTimeblockName = errors.New("name is required")
	ErrMissingDays          = errors.New("days are required when recur is true")
	ErrMissingRecur         = errors.New("recur is required when days are provided")
	ErrStartAfterEnd        = errors.New("start date cannot be after end date")
)

type TimeBlock struct {
	ID        uint `gorm:"primarykey"` // identifier for timeblock
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	TagID     uint          `required:"false" doc:"The ID of the Tag"`
	Name      string        `required:"false" doc:"Name for the timeblock" example:"SuperCoolName"`
	Days      *string       `required:"false" doc:"Days of the week where the timeblock reoccurs in binary form. STARTING ON MONDAY! Example is every week on Tuesday" example:"0100000"`
	Recur     bool          `required:"false" default:"false" doc:"Does the timeblock reoccur. If this is set, days also needs to be set."`
	StartDate dtegorm.Date  `required:"true" doc:"Date to start" format:"date"`
	EndDate   *dtegorm.Date `required:"false" doc:"Date to end" format:"date"`
	StartTime dtegorm.Time  `required:"true" doc:"Time to start" format:"time"`
	EndTime   dtegorm.Time  `required:"true" doc:"Time to end" format:"time"`
}

func (t TimeBlock) Verify() error {
	if t.Name == "" {
		return ErrMissingTimeblockName
	}

	if t.Recur && t.Days == nil {
		return ErrMissingDays
	}

	if t.Days != nil && !t.Recur {
		return ErrMissingRecur
	}

	a := t.EndDate
	if t.EndDate != nil && t.StartDate.After(a.Time) {
		return ErrStartAfterEnd
	}

	return nil
}
