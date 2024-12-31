package tagapi

import (
	pgmodels "github.com/Sourceware-Lab/realquick/database/postgres/models"
)

type TagGetInput struct {
	ID uint64 `doc:"Id for the tag you want to get" example:"999" path:"id"`
}

type TagGetOutput struct {
	Body struct {
		TagPostBodyInput
		ID uint64 `json:"id"`
	}
}

type TagPostInput struct {
	Body TagPostBodyInput `json:"body"`
}

type TagPostOutput struct {
	Body struct {
		ID uint64 `doc:"Id for new tag" example:"999" path:"id"`
	}
}

type TagPostBodyInput struct {
	pgmodels.Tag
}

func (i *TagPostBodyInput) TableName() string {
	return "tags"
}
