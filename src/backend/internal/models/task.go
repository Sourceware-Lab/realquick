package models

import (
	"errors"
	"math"
	"time"

	"github.com/Sourceware-Lab/go-huma-gin-postgres-template/internal/utils"
)

type Task struct {
	ID         int       // identifier for a task
	Name       string    // name for a task
	Due        time.Time // due date for a task
	TotalHours float64   // total hours for a task
	Hours      float64   // current hours remaining for a task
	SubHours   float64   // hours scheduled for a task
	Tag        *Tag      // label and color for a task
}

type SubTask struct {
	Task     *Task           // pointer to parent task
	Hours    float64         // hours for a scheduled subtask
	Overflow bool            // whether scheduled subtask overflows
	TS       utils.TimeStamp // timestamp for scheduled subtask
}

// creates a new Task while ensuring name is non-empty and hours are 0 or greater (only 0.5 decimal is allowed)
func NewTask(id int, name string, due time.Time, hours float64, tag *Tag) (*Task, error) {
	t := new(Task)
	t.ID = id
	if err := t.SetName(name); err != nil {
		return nil, err
	}
	t.SetDue(due)
	if err := t.SetTotalHours(hours); err != nil {
		return nil, err
	}
	t.SetTag(tag)
	return t, nil
}

// creates a new SubTask while ensuring hours are greater than 0 and within Task's remaining hours
func (t *Task) NewSubTask(hours float64, overflow bool, ts utils.TimeStamp) (*SubTask, error) {
	st := new(SubTask)
	if hours <= 0 {
		return nil, errors.New("hours for subtask cannot be zero or negative")
	}
	dec := hours - math.Floor(hours)
	if dec != 0.0 && dec != 0.5 {
		return nil, errors.New("hours for subtask cannot have a decimal besides 0.5")
	}
	if hours > t.GetSubHoursRemaining() {
		return nil, errors.New("hours for subtask cannot be greater than current hours remaining")
	}
	st.Hours = hours
	st.Overflow = overflow
	st.TS = ts
	t.SubHours += hours
	return st, nil
}

// sets name for Task (cannot be empty)
func (t *Task) SetName(name string) error {
	if name == "" {
		return errors.New("name for task cannot be empty")
	}
	t.Name = name
	return nil
}

// sets due date for Task
func (t *Task) SetDue(due time.Time) {
	t.Due = due
}

// sets total hours for Task (does not decrement over days, it stays the same)
func (t *Task) SetTotalHours(hours float64) error {
	if hours <= 0 {
		return errors.New("hours for task cannot be negative")
	}
	dec := hours - math.Floor(hours)
	if dec != 0.0 && dec != 0.5 {
		return errors.New("hours for task cannot have a decimal besides 0.5")
	}
	t.TotalHours = hours
	t.Hours = hours
	return nil
}

// gets subtotal hours remaining for Task
func (t Task) GetSubHoursRemaining() float64 {
	return t.Hours - t.SubHours
}

// sets Tag for Task
func (t *Task) SetTag(tag *Tag) {
	t.Tag = tag
}

// sets Task SubHours to 0
func (t *Task) Reset() {
	t.SubHours = 0
}
