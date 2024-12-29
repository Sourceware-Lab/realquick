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

	TagID     uint          `doc:"The ID of the Tag"                                                                                                  required:"false"`                                                                              //nolint:lll
	Name      string        `doc:"Name for the timeblock"                                                                                             example:"SuperCoolName"                                                      required:"false"` //nolint:lll
	Days      *string       `doc:"Days of the week where the timeblock reoccurs in binary form. STARTING ON MONDAY! Example is every week on Tuesday" example:"0100000"                                                            required:"false"` //nolint:lll
	Recur     bool          `default:"false"                                                                                                          doc:"Does the timeblock reoccur. If this is set, days also needs to be set." required:"false"` //nolint:lll
	StartDate dtegorm.Date  `doc:"Date to start"                                                                                                      format:"date"                                                                required:"true"`  //nolint:lll
	EndDate   *dtegorm.Date `doc:"Date to end"                                                                                                        format:"date"                                                                required:"false"` //nolint:lll
	StartTime dtegorm.Time  `doc:"Time to start"                                                                                                      format:"time"                                                                required:"true"`  //nolint:lll
	EndTime   dtegorm.Time  `doc:"Time to end"                                                                                                        format:"time"                                                                required:"true"`  //nolint:lll
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
