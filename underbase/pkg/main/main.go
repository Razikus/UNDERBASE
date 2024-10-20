package main

import (
	"encoding/json"
	"os"

	"github.com/Razikus/underbase/pkg/api"
	"github.com/Razikus/underbase/pkg/auth"
	"github.com/Razikus/underbase/server"
	"github.com/getkin/kin-openapi/openapi3filter"
	openapimiddleware "github.com/go-openapi/runtime/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echom "github.com/oapi-codegen/echo-middleware"
)

func main() {

	authorizerEndpoint := os.Getenv("AUTHORIZER_ENDPOINT")
	if authorizerEndpoint == "" {
		panic("AUTHORIZER_ENDPOINT is required like http://authorizer:9999")
	}
	authorizer := auth.NewHTTPAuthorizer(auth.HTTPAuthorizerConfig{
		Endpoint: authorizerEndpoint,
	})

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	helloHandler := api.HelloHandler{}
	handler := api.NewAPI(&helloHandler)
	strictHandler := server.NewStrictHandler(handler, nil)

	apiV1 := e.Group("/api/v1")

	swaggerDoc, err := server.GetSwagger()
	if err != nil {
		panic(err)
	}

	apiV1.Use(echom.OapiRequestValidatorWithOptions(swaggerDoc, &echom.Options{
		Options: openapi3filter.Options{
			AuthenticationFunc: authorizer.Validate,
		},
		ErrorHandler: func(c echo.Context, err *echo.HTTPError) error {
			return c.String(err.Code, "ERROR: "+err.Error())
		},
	}))

	e.GET("/api/v1/swagger/doc.json", func(ctx echo.Context) error {
		json.NewEncoder(ctx.Response().Writer).Encode(swaggerDoc)
		return nil
	})
	e.Any("/api/v1/swagger", func(ctx echo.Context) error {
		openapimiddleware.SwaggerUI(openapimiddleware.SwaggerUIOpts{
			Path:    "/api/v1/swagger/",
			SpecURL: "/api/v1/swagger/doc.json",
		}, nil).ServeHTTP(ctx.Response().Writer, ctx.Request())
		return nil
	})
	server.RegisterHandlers(apiV1, strictHandler)

	e.Logger.Fatal(e.Start(":3000"))
}
