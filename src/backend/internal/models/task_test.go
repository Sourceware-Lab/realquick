package models

import (
	"reflect"
	"testing"
	"time"
)

func TestNewTaskValid(t *testing.T) {
	due := time.Now()
	t1, err := NewTask(0, "HW3", due, 4.5, nil)
	exp := &Task{
		ID:         0,
		Name:       "HW3",
		Due:        due,
		TotalHours: 4.5,
		Hours:      4.5,
		Tag:        nil,
	}
	if err != nil {
		t.Fatalf("NewTask(0, \"HW3\", time.Now(), 4.5, nil) error = %v; want nil", err)
	}

	if !reflect.DeepEqual(exp, t1) {
		t.Errorf("NewTask(0, \"HW3\", time.Now(), 4.5, nil) = %v; want %v", t1, exp)
	}
}

func TestNewTaskInvalid(t *testing.T) {
	tests := []struct {
		test       string
		id         int
		name       string
		due        time.Time
		totalHours float64
		tag        *Tag
		wantErr    string
	}{
		{"empty name", 0, "", time.Now(), 2.0, nil, "name for task cannot be empty"},
		{"negative hours", 1, "Ch3", time.Now(), -1, nil, "hours for task cannot be negative"},
		{"invalid decimal", 2, "HW4", time.Now(), 3.2, nil, "hours for task cannot have a decimal besides 0.5"},
	}

	for _, tt := range tests {
		t.Run(tt.test, func(t *testing.T) {
			t1, err := NewTask(tt.id, tt.name, tt.due, tt.totalHours, tt.tag)
			if err == nil || err.Error() != tt.wantErr {
				t.Errorf("NewTask(%d, %q, %v, %f, %v) = %v; want error %q, got %v", tt.id, tt.name, tt.due, tt.totalHours, tt.tag, t1, tt.tag, tt.wantErr)
			}
		})
	}
}

func TestSetNameValid(t *testing.T) {
	t1, _ := NewTask(0, "HW3", time.Now(), 4.5, nil)

	err := t1.SetName("Finish HW3")
	if err != nil {
		t.Fatalf("t1.SetName(\"Finish HW3\") error = %v; want nil", err)
	}

	if "Finish HW3" != t1.Name {
		t.Errorf("t1.SetName(\"Finish HW3\") = %v; want \"Finish HW3\"", t1)
	}
}

func TestSetNameInvalid(t *testing.T) {
	t1, _ := NewTask(0, "HW3", time.Now(), 4.5, nil)

	err := t1.SetName("")
	if err == nil {
		t.Error("t1.SetName(\"\") error = nil; want \"name for task cannot be empty\"")
	}
}

func TestSetDue(t *testing.T) {
	t1, _ := NewTask(0, "HW3", time.Now(), 4.5, nil)

	due := time.Now().Add(10000)
	t1.SetDue(due)

	if due != t1.Due {
		t.Errorf("t1.SetDue(due) = %v; want %v", t1.Due, due)
	}
}

func TestSetTotalHoursValid(t *testing.T) {
	t1, _ := NewTask(0, "HW3", time.Now(), 4.5, nil)

	err := t1.SetTotalHours(6)
	if err != nil {
		t.Errorf("t1.SetTotalHours(6) = %v; want nil", err)
	}

	if t1.TotalHours != 6 || t1.Hours != 6 {
		t.Errorf("t1.SetTotalHours(6) = %v; want nil", err)
	}
}

func TestSetTotalHoursInvalid(t *testing.T) {
	t1, _ := NewTask(0, "HW3", time.Now(), 4.5, nil)

	err := t1.SetTotalHours(0)
	if err == nil {
		t.Errorf("t1.SetTotalHours(0) = nil; want %v", err)
	}

	err = t1.SetTotalHours(3.2)
	if err == nil {
		t.Errorf("t1.SetTotalHours(3.2) = nil; want %v", err)
	}
}

func TestGetSubHoursRemaining(t *testing.T) {
	t1, _ := NewTask(0, "HW3", time.Now(), 4.5, nil)
	t1.SubHours = 2.5
	if t1.GetSubHoursRemaining() != 2 {
		t.Errorf("t1.GetSubHoursRemaining() = %f; want 2", t1.GetSubHoursRemaining())
	}
}

func TestSetTag(t *testing.T) {

}

func TestRest(t *testing.T) {

}
