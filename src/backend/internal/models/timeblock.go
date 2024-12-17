package models

import (
	"time"

	"github.com/Sourceware-Lab/go-huma-gin-postgres-template/internal/utils"
)

type TimeBlock struct {
	ID        int       // identifier for timeblock
	Tag       *Tag      // label and color for timeblock
	Name      string    // name for timeblock
	Days      *string   // days of the week timeblock recurs
	Recur     bool      // whether timeblock recurs
	StartDate time.Time // start date for timeblock
	EndDate   time.Time // end date for timeblock
	//TS        timestamp // timestamp for timeblock
	Dur time.Duration
}

// creates a new timeblock while ensuring name is non-empty, timestamp is for same day (days must be non-nil if recur is false)
// days must hold a 7 digit string possessing only 1s and 0s representing the days of the week
func NewRecurTB(id int, name string, ts utils.TimeStamp, days []int, tag *Tag) (*TimeBlock, error)

// creates a new timeblock while ensuring name is non-empty, timestamp is for same day (days must be nil if recur is false)
func NewIndivTB(id int, name string, ts utils.TimeStamp, tag *Tag) (*TimeBlock, error)

// sets name for timeblock (cannot be empty)
func (t *TimeBlock) SetName(name string) error

// sets timestamp for timeblock
func (t *TimeBlock) SetTimeStamp(ts utils.TimeStamp) error

// sets recurring days of week for timeblock
// days must hold a 7 digit string possessing only 1s and 0s representing the days of the week
func (t *TimeBlock) SetDays(days string) error

// sets Tag for timeblock
func (t *TimeBlock) SetTag(tag *Tag)
