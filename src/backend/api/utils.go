package api

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"

	timeblockapi "github.com/Sourceware-Lab/realquick/api/timeblock"
)

// AddRoutes This is to make testing easier. We can pass a testing API interface.
//
//nolint:funlen
func AddRoutes(api huma.API) {
	//huma.Register(api, huma.Operation{
	//	OperationID: "healthcheck",
	//	Method:      http.MethodGet,
	//	Path:        "/healthcheck",
	//	Summary:     "healthcheck",
	//	Description: "healthcheck returns a 200 if the server is running.",
	//	Tags:        []string{"Healthcheck"},
	//},
	//	healthcheck.Get,
	//)
	//
	//huma.Register(api, huma.Operation{
	//	OperationID: "get-greeting",
	//	Method:      http.MethodGet,
	//	Path:        "/greeting/{name}",
	//	Summary:     "Get a greeting",
	//	Description: "Get a greeting for a person by name.",
	//	Tags:        []string{"Greetings"},
	//},
	//	greeting.Get,
	//)
	//
	//huma.Register(api, huma.Operation{
	//	OperationID:   "post-greeting",
	//	Method:        http.MethodPost,
	//	Path:          "/greeting",
	//	Summary:       "Post a greeting",
	//	Tags:          []string{"Greetings"},
	//	DefaultStatus: http.StatusCreated,
	//},
	//	greeting.Post,
	//)

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
