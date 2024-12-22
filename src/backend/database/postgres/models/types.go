package pgmodels

import (
	"errors"
	"gorm.io/gorm"
	"regexp"
	"strings"
	"time"
)

var (
	ErrMissingName  = errors.New("name is required")
	ErrInvalidColor = errors.New("valid hex string for color is required")
)

var regex = regexp.MustCompile("^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$")

type Tag struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string
	Color     string
}

func (t Tag) Verify() error {
	if strings.TrimSpace(t.Name) == "" {
		return ErrMissingName
	}

	if !t.isValidHex() {
		return ErrInvalidColor
	}

	return nil
}

func (t Tag) isValidHex() bool {
	return regex.MatchString(t.Color)
}
