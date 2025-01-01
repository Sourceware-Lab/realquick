package tagapi_test

import (
	"encoding/json"
	"github.com/Sourceware-Lab/realquick/api"
	tagapi "github.com/Sourceware-Lab/realquick/api/tag"
	dbpg "github.com/Sourceware-Lab/realquick/database/postgres"
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/humatest"
	"net/http"
	"strconv"
	"testing"
)

//nolint:funlen, gocognit, cyclop, tparallel
func TestRoutes(t *testing.T) {
	t.Parallel()

	type input struct {
		Name  string
		Color string
	}

	tests := []struct {
		name          string
		basePath      string
		input         input
		expectedErr   bool
		wantErrDetail string
		wantErrStatus int
	}{
		{
			name:     "tag post request",
			basePath: "/tag",
			input: input{
				Name:  "MATH",
				Color: "#00e4ff",
			},
			expectedErr: false,
		},
	}

	for _, tt := range tests { //nolint:paralleltest
		t.Run(tt.name, func(t *testing.T) {
			dbName := dbpg.Setup()
			defer func() { dbpg.Teardown(dbName) }()

			_, apiInstance := humatest.New(t)
			api.AddRoutes(apiInstance)

			resp := apiInstance.Post(tt.basePath, tt.input)
			if resp.Code != http.StatusCreated { //nolint:nestif
				if !tt.expectedErr {
					t.Fatalf("Unexpected response: %s", resp.Body.String())
				}

				respErr := huma.ErrorModel{}

				err := json.Unmarshal(resp.Body.Bytes(), &respErr)
				if err != nil {
					t.Fatalf("Error unmarshalling JSON: %s", err.Error())
				}

				if respErr.Detail != tt.wantErrDetail {
					t.Fatalf("Incorrect error detail: %s", respErr.Detail)
				}

				if respErr.Status != tt.wantErrStatus {
					t.Fatalf("Incorrect error status: %d", respErr.Status)
				}

				return
			} else if tt.expectedErr {
				t.Fatalf("expected error, got %d", resp.Code)
			}

			postRespBody := tagapi.TagPostOutput{}.Body

			err := json.Unmarshal(resp.Body.Bytes(), &postRespBody) //nolint:musttag
			if err != nil {
				t.Fatalf("Failed to unmarshal response: %s", err.Error())
			}

			getResp := apiInstance.Get(tt.basePath + "/" + strconv.FormatUint(postRespBody.ID, 10))
			getRespBody := tagapi.TagGetOutput{}.Body

			err = json.Unmarshal(getResp.Body.Bytes(), &getRespBody) //nolint:musttag
			if err != nil {
				t.Fatalf("Failed to unmarshal response: %s", err.Error())
			}

			if tt.input.Name != getRespBody.Name {
				t.Fatalf("Incorrect response name: %s", resp.Body.String())
			}

			if tt.input.Color != getRespBody.Color {
				t.Fatalf("Incorrect response color: %s", resp.Body.String())
			}

			if postRespBody.ID != getRespBody.ID {
				t.Fatalf("Incorrect response ID: %d", postRespBody.ID)
			}
		})
	}
}
