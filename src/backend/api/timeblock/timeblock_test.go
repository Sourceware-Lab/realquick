package timeblockapi_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/humatest"
	"github.com/peterHoburg/go-date-and-time-extension/dtegorm"

	"github.com/Sourceware-Lab/realquick/api"
	timeblockapi "github.com/Sourceware-Lab/realquick/api/timeblock"
	dbpg "github.com/Sourceware-Lab/realquick/database/postgres"
	pgmodels "github.com/Sourceware-Lab/realquick/database/postgres/models"
	"github.com/Sourceware-Lab/realquick/internal/utils"
)

//nolint:funlen
//nolint:gocognit
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

	tests := []struct {
		name          string
		basePath      string
		input         input
		expectErr     bool
		wantErrDetail string
		wantErrStatus int
	}{
		{
			name:     "mostly empty",
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
			expectErr: false,
		},

		{
			name:     "Full",
			basePath: "/timeblock",
			input: input{
				TagID:     1,
				Name:      "some dumb name",
				Days:      utils.MakePointer("0100000"),
				Recur:     true,
				StartDate: func() dtegorm.Date { t, _ := dtegorm.NewDate("2023-01-02"); return t }(),
				EndDate:   utils.MakePointer(func() dtegorm.Date { t, _ := dtegorm.NewDate("2023-01-03"); return t }()),
				StartTime: func() dtegorm.Time { t, _ := dtegorm.NewTime("15:10:04Z"); return t }(),
				EndTime:   func() dtegorm.Time { t, _ := dtegorm.NewTime("17:10:04Z"); return t }(),
			},
			expectErr: false,
		},
		{
			name:     "with days, missing recur",
			basePath: "/timeblock",
			input: input{
				TagID:     1,
				Name:      "some dumb name",
				Days:      utils.MakePointer("0100000"),
				Recur:     false,
				StartDate: func() dtegorm.Date { t, _ := dtegorm.NewDate("2023-01-02"); return t }(),
				EndDate:   nil,
				StartTime: dtegorm.Time{},
				EndTime:   dtegorm.Time{},
			},
			expectErr:     true,
			wantErrDetail: "validation failed",
			wantErrStatus: http.StatusUnprocessableEntity,
		},

		{
			name:     "start date after end date",
			basePath: "/timeblock",
			input: input{
				TagID:     1,
				Name:      "some dumb name",
				Days:      nil,
				Recur:     false,
				StartDate: func() dtegorm.Date { t, _ := dtegorm.NewDate("2023-01-02"); return t }(),
				EndDate:   utils.MakePointer(func() dtegorm.Date { t, _ := dtegorm.NewDate("2023-01-01"); return t }()),
				StartTime: dtegorm.Time{},
				EndTime:   dtegorm.Time{},
			},
			expectErr:     true,
			wantErrDetail: "validation failed",
			wantErrStatus: http.StatusUnprocessableEntity,
		},
		{
			name:     "mostly empty",
			basePath: "/timeblock",
			input: input{
				TagID:     1,
				Name:      "some dumb name",
				Days:      nil,
				Recur:     false,
				StartDate: dtegorm.Date{},
				EndDate:   nil,
				StartTime: func() dtegorm.Time { t, _ := dtegorm.NewTime("15:10:04Z"); return t }(),
				EndTime:   func() dtegorm.Time { t, _ := dtegorm.NewTime("14:10:04Z"); return t }(),
			},
			expectErr: true,
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

			if resp.Code != http.StatusCreated {
				if !tt.expectErr {
					t.Fatalf("Unexpected response: %s", resp.Body.String())
				}
				respErr := huma.ErrorModel{}

				err := json.Unmarshal(resp.Body.Bytes(), &respErr)
				if err != nil {
					t.Fatalf("Failed to unmarshal response: %s", err.Error())
				}

				if respErr.Detail != tt.wantErrDetail {
					t.Fatalf("Incorrect error detail: %s", respErr.Detail)
				}

				if respErr.Status != tt.wantErrStatus {
					t.Fatalf("Incorrect error status: %d", respErr.Status)
				}

				return
			} else if tt.expectErr {
				t.Fatalf("expected error, got %d", resp.Code)
			}
			postRespBody := timeblockapi.TimeblockPostOutput{}.Body

			err := json.Unmarshal(resp.Body.Bytes(), &postRespBody)
			if err != nil {
				t.Fatalf("Failed to unmarshal response: %s", err.Error())
			}

			getResp := apiInstance.Get(tt.basePath + "/" + fmt.Sprint(postRespBody.ID))
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

			if tt.input.StartDate.String() != getRespBody.StartDate.String() {
				t.Fatalf("Unexpected response: %s", resp.Body.String())
			}

			if tt.input.EndDate == nil {
				if getRespBody.EndDate != nil {
					t.Fatalf("Unexpected response: %s", resp.Body.String())
				}
			} else {
				if tt.input.EndDate.String() != getRespBody.EndDate.String() {
					t.Fatalf("Unexpected response: %s", resp.Body.String())
				}
			}

			if tt.input.StartTime.String() != getRespBody.StartTime.String() {
				t.Fatalf("Unexpected response: %s", resp.Body.String())
			}

			if tt.input.EndTime.String() != getRespBody.EndTime.String() {
				t.Fatalf("Unexpected response: %s", resp.Body.String())
			}

			if postRespBody.ID != getRespBody.ID {
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
