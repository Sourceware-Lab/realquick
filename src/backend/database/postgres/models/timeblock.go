package pgmodels

import (
	"errors"
	"time"

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

	Tag       *Tag       // label and color for timeblock
	Name      string     // name for timeblock
	Days      *string    // days of the week timeblock recurs
	Recur     bool       // whether timeblock recurs
	StartDate time.Time  // start date for timeblock
	EndDate   *time.Time // end date for timeblock
	// TimeStamp utils.TimeStamp // timestamp for timeblock
	Duration time.Duration // duration for a timestamp
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

	if t.EndDate != nil && t.StartDate.After(*t.EndDate) {
		return ErrStartAfterEnd
	}

	return nil
}
