package pgmodels

import (
	"errors"
	"regexp"
	"strings"
	"time"

	"gorm.io/gorm"
)

var (
	ErrMissingTagName = errors.New("name is required")
	ErrInvalidColor   = errors.New("valid hex string for color is required")
	regex             = regexp.MustCompile("^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$")
)

type Tag struct {
	ID        uint64         `gorm:"primarykey" json:"-"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index"      json:"-"`

	Name  string `doc:"Name for the tag"  example:"MATH"    required:"true"`
	Color string `doc:"Color for the tag" example:"#ff6100" required:"true"`
}

func (t Tag) Verify() error {
	if strings.TrimSpace(t.Name) == "" {
		return ErrMissingTagName
	}

	if !t.isValidHex() {
		return ErrInvalidColor
	}

	return nil
}

func (t Tag) isValidHex() bool {
	return regex.MatchString(t.Color)
}
