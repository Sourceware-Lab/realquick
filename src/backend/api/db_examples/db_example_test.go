package dbexample_test

import (
	"encoding/json"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/danielgtaylor/huma/v2/humatest"
	"github.com/google/go-cmp/cmp"

	"github.com/Sourceware-Lab/realquick/api"
	dbexample "github.com/Sourceware-Lab/realquick/api/db_examples"
	dbpg "github.com/Sourceware-Lab/realquick/database/postgres"
)

//nolint:funlen
func TestRoutes(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		basePath string
		want     dbexample.PostInputDBExample
	}{
		{
			name:     "get",
			basePath: "/db_example/orm",
			want: dbexample.PostInputDBExample{
				Body: dbexample.PostBodyInputDBExampleBody{
					Name:         "jo",
					Age:          25,
					Email:        "jo@example.com",
					Birthday:     nil,
					MemberNumber: nil,
				},
			},
		},
		{
			name:     "get",
			basePath: "/db_example/raw_sql",
			want: dbexample.PostInputDBExample{
				Body: dbexample.PostBodyInputDBExampleBody{
					Name:         "jo1",
					Age:          26,
					Email:        "jo1@example.com",
					Birthday:     nil,
					MemberNumber: nil,
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
