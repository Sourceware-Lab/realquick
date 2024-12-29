package timeblockapi_test

import (
	"encoding/json"
	"testing"

	"github.com/danielgtaylor/huma/v2/humatest"
	"github.com/peterHoburg/go-date-and-time-extension/dtegorm"

	"github.com/Sourceware-Lab/realquick/api"
	timeblockapi "github.com/Sourceware-Lab/realquick/api/timeblock"
	dbpg "github.com/Sourceware-Lab/realquick/database/postgres"
	pgmodels "github.com/Sourceware-Lab/realquick/database/postgres/models"
)

//nolint:funlen
func TestRoutes(t *testing.T) {
	t.Parallel()

	type input struct {
		TagID     uint
		Name      string
		Days      *string
		Recur     bool
		StartDate dtegorm.Date
		EndDate   *dtegorm.Date
		StartTime dtegorm.Time
		EndTime   dtegorm.Time
	}

	type want struct {
		ID uint
	}

	tests := []struct {
		name     string
		basePath string
		input    input
		want     want
	}{
		{
			name:     "post",
			basePath: "/timeblock",
			input: input{
				TagID:     1,
				Name:      "",
				Days:      nil,
				Recur:     false,
				StartDate: dtegorm.Date{},
				EndDate:   nil,
				StartTime: dtegorm.Time{},
				EndTime:   dtegorm.Time{},
			},
			want: want{
				ID: 1,
			},
		},
	}

	for _, tt := range tests {
		func() {
			dbName := dbpg.Setup()
			defer func() { dbpg.Teardown(dbName) }()

			result := dbpg.DB.Create(&pgmodels.Tag{
				Name:  "test",
				Color: "test",
			})
			if result.Error != nil {
				t.Fatalf("Failed to create tag: %s", result.Error.Error())
			}

			_, apiInstance := humatest.New(t)
			api.AddRoutes(apiInstance)

			resp := apiInstance.Post(tt.basePath, tt.input)

			postRespBody := timeblockapi.TimeblockPostOutput{}.Body

			err := json.Unmarshal(resp.Body.Bytes(), &postRespBody)
			if err != nil {
				t.Fatalf("Failed to unmarshal response: %s", err.Error())
			}

			// if !cmp.Equal(postRespBody, tt.want) { TODO why is this not working?
			//	t.Fatalf("Unexpected response: %s", resp.Body.String())
			//}

			if postRespBody.ID != tt.want.ID {
				t.Fatalf("Unexpected response: %s", resp.Body.String())
			}

			//	getResp := apiInstance.Get(tt.basePath + "/1")
			//	getRespBody := dbexample.GetOutputDBExample{}
			//
			//	err = json.Unmarshal(getResp.Body.Bytes(), &getRespBody)
			//	if err != nil {
			//		t.Fatalf("Failed to unmarshal response: %s", err.Error())
			//	}
			//
			//	if !cmp.Equal(getRespBody.Body, tt.want.Body) {
			//		t.Fatalf("Unexpected response: %s", getResp.Body.String())
			//	}
		}() //nolint:wsl
	}
}
