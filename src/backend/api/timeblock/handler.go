package timeblockapi

import (
	"context"

	"github.com/rs/zerolog/log"

	dbpg "github.com/Sourceware-Lab/realquick/database/postgres"
)

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
