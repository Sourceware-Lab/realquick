package api

import (
	tagapi "github.com/Sourceware-Lab/realquick/api/tag"
	"net/http"

	"github.com/danielgtaylor/huma/v2"

	"github.com/Sourceware-Lab/realquick/api/healthcheck"
	timeblockapi "github.com/Sourceware-Lab/realquick/api/timeblock"
)

// AddRoutes This is to make testing easier. We can pass a testing API interface.
func AddRoutes(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "healthcheck",
		Method:      http.MethodGet,
		Path:        "/healthcheck",
		Summary:     "healthcheck",
		Description: "healthcheck returns a 200 if the server is running.",
		Tags:        []string{"Healthcheck"},
	},
		healthcheck.Get,
	)

	huma.Register(api, huma.Operation{
		OperationID:   "get-tag",
		Method:        http.MethodGet,
		Path:          "/tag/{id}",
		Summary:       "Get tag",
		Tags:          []string{"Tag"},
		DefaultStatus: http.StatusOK,
	},
		tagapi.Get,
	)

	huma.Register(api, huma.Operation{
		OperationID:   "post-tag",
		Method:        http.MethodPost,
		Path:          "/tag",
		Summary:       "Create new tag",
		Tags:          []string{"Tag"},
		DefaultStatus: http.StatusCreated,
	},
		tagapi.Post,
	)

	huma.Register(api, huma.Operation{
		OperationID:   "get-timeblock",
		Method:        http.MethodGet,
		Path:          "/timeblock/{id}",
		Summary:       "Get timeblock",
		Tags:          []string{"timeblock"},
		DefaultStatus: http.StatusOK,
	},
		timeblockapi.Get,
	)

	huma.Register(api, huma.Operation{
		OperationID:   "post-timeblock",
		Method:        http.MethodPost,
		Path:          "/timeblock",
		Summary:       "Create new timeblock",
		Tags:          []string{"timeblock"},
		DefaultStatus: http.StatusCreated,
	},
		timeblockapi.Post,
	)
}
