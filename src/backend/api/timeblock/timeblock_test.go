package timeblockapi_test

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/danielgtaylor/huma/v2/humatest"
	"github.com/peterHoburg/go-date-and-time-extension/dtegorm"

	"github.com/Sourceware-Lab/realquick/api"
	timeblockapi "github.com/Sourceware-Lab/realquick/api/timeblock"
	dbpg "github.com/Sourceware-Lab/realquick/database/postgres"
	pgmodels "github.com/Sourceware-Lab/realquick/database/postgres/models"
	"github.com/Sourceware-Lab/realquick/internal/utils"
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
	// TODO write test that expects error from put
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

		{
			name:     "post",
			basePath: "/timeblock",
			input: input{
				TagID:     1,
				Name:      "some dumb name",
				Days:      utils.MakePointer("0100000"),
				Recur:     true,
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
		t.Run(tt.name, func(t *testing.T) {
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

			getResp := apiInstance.Get(tt.basePath + "/1")
			getRespBody := timeblockapi.TimeblockGetOutput{}.Body

			err = json.Unmarshal(getResp.Body.Bytes(), &getRespBody)
			if err != nil {
				t.Fatalf("Failed to unmarshal response: %s", err.Error())
			}

			if tt.input.TagID != getRespBody.TagID {
				t.Fatalf("Unexpected response: %s", resp.Body.String())
			}

			if tt.input.Name != getRespBody.Name {
				t.Fatalf("Unexpected response: %s", resp.Body.String())
			}

			if tt.input.Days == nil {
				if getRespBody.Days != nil {
					t.Fatalf("Unexpected response: %s", resp.Body.String())
				}
			} else {
				if *tt.input.Days != *getRespBody.Days {
					t.Fatalf("Unexpected response: %s", resp.Body.String())
				}
			}

			if tt.input.Recur != getRespBody.Recur {
				t.Fatalf("Unexpected response: %s", resp.Body.String())
			}

			if tt.input.StartDate != getRespBody.StartDate {
				t.Fatalf("Unexpected response: %s", resp.Body.String())
			}

			if tt.input.EndDate != getRespBody.EndDate {
				t.Fatalf("Unexpected response: %s", resp.Body.String())
			}

			if tt.input.StartTime.String() != getRespBody.StartTime.String() {
				t.Fatalf("Unexpected response: %s", resp.Body.String())
			}

			if tt.input.EndTime.String() != getRespBody.EndTime.String() {
				t.Fatalf("Unexpected response: %s", resp.Body.String())
			}

			if tt.want.ID != getRespBody.ID {
				t.Fatalf("Unexpected response: %s", resp.Body.String())
			}

		}) //nolint:wsl
	}
}

func compareStructs(input, wanted interface{}) bool {
	inputValue := reflect.ValueOf(input)
	wantedValue := reflect.ValueOf(wanted)

	// Ensure both are structs
	if inputValue.Kind() != reflect.Struct || wantedValue.Kind() != reflect.Struct {
		fmt.Println("Both input and wanted must be structs")
		return false
	}

	inputType := inputValue.Type()

	// Iterate over the fields of the input struct
	for i := 0; i < inputValue.NumField(); i++ {
		inputField := inputValue.Field(i)
		fieldName := inputType.Field(i).Name
		wantedField := wantedValue.FieldByName(fieldName)

		// Check if the field exists in wanted and compare values
		if !wantedField.IsValid() {
			fmt.Printf("Field %s not found in wanted\n", inputType.Field(i).Name)
			return false
		}

		if inputField.Interface() != wantedField.Interface() {
			fmt.Printf("Mismatch in field %s: input=%v, wanted=%v\n",
				inputType.Field(i).Name, inputField.Interface(), wantedField.Interface())
			return false
		}
	}

	return true
}
