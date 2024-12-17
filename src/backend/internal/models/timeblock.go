package models

import (
	"errors"
	"time"

	"github.com/Sourceware-Lab/go-huma-gin-postgres-template/internal/utils"
)

type TimeBlock struct {
	ID        int             // identifier for timeblock
	Tag       *Tag            // label and color for timeblock
	Name      string          // name for timeblock
	Days      *string         // days of the week timeblock recurs
	Recur     bool            // whether timeblock recurs
	StartDate time.Time       // start date for timeblock
	EndDate   time.Time       // end date for timeblock
	TS        utils.TimeStamp // timestamp for timeblock
	Dur       time.Duration   // duration for a timestamp (not stored in DB)
}

// creates a new timeblock while ensuring name is non-empty
// days must hold a 7 digit string possessing only 1s and 0s representing the days of the week
func NewRecurTB(id int, name string, days *string, start, end time.Time, ts utils.TimeStamp, tag *Tag) (*TimeBlock, error) {
	tb := new(TimeBlock)
	tb.ID = id
	tb.Recur = true
	if err := tb.SetName(name); err != nil {
		return nil, err
	}
	if err := tb.SetDays(days); err != nil {
		return nil, err
	}
	if err := tb.SetDates(start, end); err != nil {
		return nil, err
	}
	tb.SetTimeStamp(ts)
	if err := tb.SetDays(tb.Days); err != nil {
		return nil, err
	}
	tb.SetTag(tag)
	return tb, nil
}

// creates a new timeblock while ensuring name is non-empty, timestamp is for same day (days must be nil if recur is false)
func NewIndivTB(id int, name string, start, end time.Time, ts utils.TimeStamp, tag *Tag) (*TimeBlock, error) {
	tb := new(TimeBlock)
	tb.ID = id
	tb.Recur = false
	if err := tb.SetName(name); err != nil {
		return nil, err
	}
	if err := tb.SetDates(start, end); err != nil {
		return nil, err
	}
	tb.SetTimeStamp(ts)
	if err := tb.SetDays(tb.Days); err != nil {
		return nil, err
	}
	tb.SetTag(tag)
	return tb, nil
}

// sets Tag for timeblock
func (t *TimeBlock) SetTag(tag *Tag) {
	t.Tag = tag
}

// sets name for timeblock (cannot be empty)
func (t *TimeBlock) SetName(name string) error {
	if name == "" {
		return errors.New("name for timeblock cannot be empty")
	}
	t.Name = name
	return nil
}

// sets recurring days of week for timeblock,
// days must hold a 7 digit string possessing only 1s and 0s representing the days of the week
func (t *TimeBlock) SetDays(days *string) error {
	if !t.Recur {
		return errors.New("recurring days cannot be provided for individual timeblock")
	}
	if days == nil {
		return errors.New("recurring days cannot be nil")
	}
	if len(*days) != 7 {
		return errors.New("recurring days must have a length of 7")
	}
	for _, r := range *days {
		if r != '0' && r != '1' {
			return errors.New("recurring days in string must be a '0' or '1'")
		}
	}
	t.Days = days
	return nil
}

// need to make sure that only date is passed by handler during creation, no time
func (t *TimeBlock) SetDates(start, end time.Time) error {
	if !t.Recur && !utils.IsSameDate(start, end) {
		return errors.New("start and end dates for individual timeblocks must be the same")
	}
	if end.Before(start) {
		return errors.New("end date must be after start date for timeblocks")
	}
	t.StartDate = start
	t.EndDate = end
	return nil
}

// sets timestamp for timeblock (timestamp will validate that it is same day)
func (t *TimeBlock) SetTimeStamp(ts utils.TimeStamp) {
	t.TS = ts
	t.Dur = ts.End.Sub(ts.Start)
}
