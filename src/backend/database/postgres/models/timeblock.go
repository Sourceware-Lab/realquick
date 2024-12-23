package pgmodels

import (
	"errors"
	"time"

	"gorm.io/gorm"

	"github.com/Sourceware-Lab/realquick/internal/types"
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

	TagID     uint          // ID of tag. Tag obj has ref to make this a FK
	Name      string        // name for timeblock
	Days      *string       // days of the week timeblock reoccur
	Recur     bool          // whether timeblock reoccur
	StartDate time.Time     // start date for timeblock
	EndDate   *time.Time    // end date for timeblock
	TimeStamp types.Time    // timestamp for timeblock
	Duration  time.Duration // duration for a timestamp
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
