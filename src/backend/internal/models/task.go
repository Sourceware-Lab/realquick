package models

import "time"

type Task struct {
	Id         int       // identifier for a task
	Name       string    // name for a task
	Due        time.Time // due date for a task
	TotalHours float64   // total hours for a task
	Hours      float64   // hours remaining for a task
	SubHours   float64   // hours scheduled for a task
	Tag        *Tag      // label and color for a task
}

type SubTask struct {
	Task     *Task   // pointer to parent task
	Hours    float64 // hours for a scheduled subtask
	Overflow bool    // whether scheduled subtask overflows
	//TS       TimeStamp // timestamp for scheduled subtask
}

// creates a new Task while ensuring name is non-empty and hours are 0 or greater (only 0.5 decimal is allowed)
func NewTask(id int, name string, due time.Time, hours float64, tag *Tag) (*Task, error)

// creates a new SubTask while ensuring hours are greater than 0 and within Task's remaining hours
//func (t *Task) NewSubTask(hours float64, overflow bool, ts TimeStamp) *SubTask

// sets name for Task (cannot be empty)
func (t *Task) SetName(name string) error

// sets due date for Task
func (t *Task) SetDue(due time.Time)

// sets total hours for Task
func (t *Task) SetTotalHours(hours float64) error

// sets hours for Task
func (t *Task) SetHours(hours float64) error

// gets subtotal hours remaining for Task
func (t Task) GetSubHoursRemaining() float64

// sets Tag for Task
func (t *Task) SetTag(tag *Tag)

// gets formatted date stamp for Task
func (t Task) GetDateStamp() string

// sets Task SubHours to 0
func (t *Task) Reset()
