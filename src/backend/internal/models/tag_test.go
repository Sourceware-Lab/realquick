package models

import (
	"reflect"
	"testing"
)

func TestNewTagValid(t *testing.T) {
	tag, err := NewTag(0, "MA", "#40e0d0")
	exp := Tag{
		ID:    0,
		Name:  "MA",
		Color: "#40e0d0",
	}
	if err != nil {
		t.Fatalf("newTag(0, \"MA\", BLUE) error = %v; want nil", err)
	}

	if reflect.DeepEqual(exp, tag) {
		t.Errorf("newCard(0, \"MA\", BLUE) = %v; want %v", tag, exp)
	}
}

func TestNewTagInvalid(t *testing.T) {
	tests := []struct {
		id      int
		name    string
		color   string
		wantErr string
	}{
		{0, "MA", "fail", "color for tag is not valid"},
		{0, "", "#40e0d0", "name for tag cannot be empty"},
	}

	for _, tt := range tests {
		t.Run(tt.wantErr, func(t *testing.T) {
			tag, err := NewTag(tt.id, tt.name, tt.color)
			if err == nil || err.Error() != tt.wantErr {
				t.Errorf("newCard(%d, %q, %s) = %v; want error %q, got %v", tt.id, tt.name, tt.color, tag, tt.wantErr, err)
			}
		})
	}
}
