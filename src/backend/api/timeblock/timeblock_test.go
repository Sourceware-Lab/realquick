package timeblockapi_test

import (
	"encoding/json"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/danielgtaylor/huma/v2/humatest"
	"github.com/google/go-cmp/cmp"
	"github.com/peterHoburg/go-date-and-time-extension/dte"
	"github.com/peterHoburg/go-date-and-time-extension/dtegorm"
	"gorm.io/gorm"

	"github.com/Sourceware-Lab/realquick/api"
	timeblockapi "github.com/Sourceware-Lab/realquick/api/timeblock"
	dbpg "github.com/Sourceware-Lab/realquick/database/postgres"
	pgmodels "github.com/Sourceware-Lab/realquick/database/postgres/models"
)

//nolint:funlen
func TestRoutes(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		basePath string
		want     timeblockapi.TimeblockPostInput
	}{
		{
			name:     "get",
			basePath: "/db_example/orm",
			want: timeblockapi.TimeblockPostInput{
				Body: timeblockapi.TimeblockPostBodyInput{
					TimeBlock: pgmodels.TimeBlock{
						TagID:     0,
						Name:      "",
						Days:      nil,
						Recur:     false,
						StartDate: dtegorm.Date{},
						EndDate:   &dtegorm.Date{},
						StartTime: dtegorm.Time{},
						EndTime:   dtegorm.Time{},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		func() {
			dbName := dbpg.Setup()
			defer dbpg.Teardown(dbName)

			_, apiInstance := humatest.New(t)
			api.AddRoutes(apiInstance)

			birthdayTime := time.Now().Add(time.Duration(-tt.want.Body.Age) * time.Hour * 24 * 365)
			birthday := birthdayTime.Format(time.DateOnly)
			tt.want.Body.Birthday = &birthday

			memberNumber := strconv.Itoa(rand.Intn(1000000)) //nolint:gosec
			tt.want.Body.MemberNumber = &memberNumber

			resp := apiInstance.Post(tt.basePath, tt.want.Body)

			postRespBody := dbexample.PostOutputDBExample{}.Body

			err := json.Unmarshal(resp.Body.Bytes(), &postRespBody)
			if err != nil {
				t.Fatalf("Failed to unmarshal response: %s", err.Error())
			}

			expectedPostBody := dbexample.PostOutputDBExample{}.Body
			expectedPostBody.ID = "1"

			if !cmp.Equal(postRespBody, expectedPostBody) {
				t.Fatalf("Unexpected response: %s", resp.Body.String())
			}

			getResp := apiInstance.Get(tt.basePath + "/1")
			getRespBody := dbexample.GetOutputDBExample{}

			err = json.Unmarshal(getResp.Body.Bytes(), &getRespBody)
			if err != nil {
				t.Fatalf("Failed to unmarshal response: %s", err.Error())
			}

			if !cmp.Equal(getRespBody.Body, tt.want.Body) {
				t.Fatalf("Unexpected response: %s", getResp.Body.String())
			}
		}()
	}
}
