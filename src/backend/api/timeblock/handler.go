package timeblockapi

import (
	"context"

	"github.com/rs/zerolog/log"

	dbpg "github.com/Sourceware-Lab/realquick/database/postgres"
)

func Get(_ context.Context, input *TimeblockGetInput) (*TimeblockGetOutput, error) {
	resp := &TimeblockGetOutput{}

	result := dbpg.DB.First(&resp.Body, input.ID)
	if result.Error != nil {
		log.Error().Err(result.Error).Msg("Error getting timeblock")

		return nil, result.Error
	}

	return resp, nil
}

func Post(_ context.Context, input *TimeblockPostInput) (*TimeblockPostOutput, error) {
	resp := &TimeblockPostOutput{}

	result := dbpg.DB.Create(&input.Body) // NOTE. This is a POINTER!
	if result.Error != nil {
		log.Error().Err(result.Error).Msg("Error creating timeblock")

		return nil, result.Error
	}

	resp.Body.ID = input.Body.ID

	return resp, nil
}
