package pgmodels

import (
	"time"

	"github.com/peterHoburg/go-date-and-time-extension/dtegorm"
	"gorm.io/gorm"
)

type TimeBlock struct {
	ID        uint           `json:"-" gorm:"primarykey"` // identifier for timeblock
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	TagID     uint          `doc:"The ID of the Tag"                                                                                                  required:"false"`                                                                              //nolint:lll
	Name      string        `doc:"Name for the timeblock"                                                                                             example:"SuperCoolName"                                                      required:"false"` //nolint:lll
	Days      *string       `doc:"Days of the week where the timeblock reoccurs in binary form. STARTING ON MONDAY! Example is every week on Tuesday" example:"0100000"                                                            required:"false"` //nolint:lll
	Recur     bool          `default:"false"                                                                                                          doc:"Does the timeblock reoccur. If this is set, days also needs to be set." required:"false"` //nolint:lll
	StartDate dtegorm.Date  `doc:"Date to start"                                                                                                      format:"date"                                                                required:"true"`  //nolint:lll
	EndDate   *dtegorm.Date `doc:"Date to end"                                                                                                        format:"date"                                                                required:"false"` //nolint:lll
	StartTime dtegorm.Time  `doc:"Time to start"                                                                                                      format:"time"                                                                required:"true"`  //nolint:lll
	EndTime   dtegorm.Time  `doc:"Time to end"                                                                                                        format:"time"                                                                required:"true"`  //nolint:lll
}
