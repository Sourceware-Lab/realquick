package timeblockapi

import (
	"context"
	//"database/sql"
	//"fmt"
	//"strconv"
	//"time"
	//
	//"github.com/rs/zerolog/log"
	//
	//dbpg "github.com/Sourceware-Lab/realquick/database/postgres"
	//pgmodels "github.com/Sourceware-Lab/realquick/database/postgres/models"
)

func Post(_ context.Context, input *TimeblockPostInput) (*TimeblockPostOutput, error) {
	//resp := &PostOutputDBExample{}
	//user := pgmodels.User{
	//	Name:         input.Body.Name,
	//	Email:        nil,
	//	Birthday:     nil,
	//	MemberNumber: sql.NullString{},
	//	ActivatedAt:  sql.NullTime{},
	//	Age:          input.Body.Age,
	//}
	//
	//if input.Body.Email != "" {
	//	user.Email = &input.Body.Email
	//}
	//
	//if input.Body.Birthday != nil {
	//	birthday, err := time.Parse(time.DateOnly, *input.Body.Birthday)
	//	if err != nil {
	//		return nil, fmt.Errorf("error parsing birthday: %w", err)
	//	}
	//
	//	user.Birthday = &birthday
	//}
	//
	//if input.Body.MemberNumber != nil {
	//	user.MemberNumber = sql.NullString{
	//		String: *input.Body.MemberNumber,
	//		Valid:  true,
	//	}
	//}
	//
	//result := dbpg.DB.Create(&user) // NOTE. This is a POINTER!
	//if result.Error != nil {
	//	log.Error().Err(result.Error).Msg("Error creating user")
	//
	//	return nil, result.Error
	//}
	//
	//resp.Body.ID = strconv.Itoa(int(user.ID)) //nolint:gosec
	//
	//return resp, nil
	a := TimeblockPostOutput{}
	return &a, nil
}
