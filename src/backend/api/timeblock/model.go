package timeblockapi

import (
	"time"

	"github.com/peterHoburg/go-date-and-time-extension/dtegorm"
)

type TimeblockPostInput struct {
	Body TimeblockPostBodyInput `json:"body"`
}

type TimeblockPostBodyInput struct {
	TagID     uint          // ID of tag. Tag obj has ref to make this a FK
	Name      string        // name for timeblock
	Days      *string       // days of the week timeblock reoccur
	Recur     bool          // whether timeblock reoccur
	StartDate dtegorm.Date  `example:"2024-01-02"`
	EndDate   *dtegorm.Date `example:"2024-01-02"`
	TimeStamp dtegorm.Time  `example:"15:04:05Z"`
	Duration  time.Duration // duration for a timestamp
}

type TimeblockPostOutput struct {
	Body struct {
		ID string `doc:"Id for new user" example:"999" json:"id"`
	}
}
