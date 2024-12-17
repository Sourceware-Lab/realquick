package models

import (
	"errors"
	"regexp"
)

const hex = `^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$`

type Tag struct {
	ID    int    // identifier for Tag
	Name  string // name for Tag
	Color string // hex string
}

// creates a new Tag while ensuring name is non-empty and hex string is valid
func NewTag(id int, name string, color string) (*Tag, error) {
	var err error
	t := new(Tag)
	err = t.SetName(name)
	if err != nil {
		return nil, err
	}
	err = t.SetColor(color)
	if err != nil {
		return nil, err
	}
	return t, nil
}

// sets name for Tag (cannot be empty)
func (t *Tag) SetName(name string) error {
	if name == "" {
		return errors.New("name for tag cannot be empty")
	}
	t.Name = name
	return nil
}

// sets color for Tag (cannot be empty)
func (t *Tag) SetColor(color string) error {
	s := validateyHex(color)
	if !s {
		return errors.New("color for tag is not valid")
	}
	t.Color = color
	return nil
}

// validates the hex code for a tag
func validateyHex(color string) bool {
	if color == "" {
		return false
	}
	re := regexp.MustCompile(hex)
	return re.MatchString(color)
}
