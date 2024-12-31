package tagapi

import (
	"context"
	dbpg "github.com/Sourceware-Lab/realquick/database/postgres"
	"github.com/rs/zerolog/log"
)

func Get(_ context.Context, input *TagGetInput) (*TagGetOutput, error) {
	resp := &TagGetOutput{}

	result := dbpg.DB.First(&resp.Body, input.ID)
	if result.Error != nil {
		log.Error().Err(result.Error).Msg("Error getting tag")

		return nil, result.Error
	}

	return resp, nil
}

func Post(_ context.Context, input *TagPostInput) (*TagPostOutput, error) {
	resp := &TagPostOutput{}

	result := dbpg.DB.Create(&input.Body)
	if result.Error != nil {
		log.Error().Err(result.Error).Msg("Error creating tag")

		return nil, result.Error
	}

	return resp, nil
}
