package auth

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	supaauth "github.com/Razikus/underbase/supaauth"
	"github.com/getkin/kin-openapi/openapi3filter"
	echom "github.com/oapi-codegen/echo-middleware"
)

type UserKey struct{}

const BearerAuth = "BearerAuth"

type HTTPAuthorizerConfig struct {
	Endpoint string
}

type HTTPAuthorizer struct {
	client *supaauth.ClientWithResponses
}

func NewHTTPAuthorizer(cfg HTTPAuthorizerConfig) *HTTPAuthorizer {
	client, err := supaauth.NewClientWithResponses(cfg.Endpoint)
	if err != nil {
		panic(err)
	}
	return &HTTPAuthorizer{client: client}
}

func (h *HTTPAuthorizer) VerifyToken(ctx context.Context, token string) (*supaauth.UserSchema, error) {
	resp, err := h.client.GetUserWithResponse(ctx, func(ctx context.Context, r *http.Request) error {
		if strings.HasPrefix(token, "Bearer ") {
			r.Header.Add("Authorization", token)
		} else {
			r.Header.Add("Authorization", "Bearer "+token)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("failed to get user: %s", resp.Status())
	}
	return resp.JSON200, nil
}

func (h *HTTPAuthorizer) Validate(ctx context.Context, openapiFilter *openapi3filter.AuthenticationInput) error {
	if openapiFilter.SecurityScheme == nil {
		return nil
	}
	if openapiFilter.SecuritySchemeName == BearerAuth {

		r := openapiFilter.RequestValidationInput.Request

		token := r.Header.Get("Authorization")
		if token == "" {
			return errors.New("unable to find auth header")
		}

		user, err := h.VerifyToken(ctx, token)
		if err != nil {
			return err
		}

		ctx = context.WithValue(ctx, UserKey{}, user)
		// This part is needed to make the context available in the echo request
		echoCtx := echom.GetEchoContext(ctx)
		echoCtx.SetRequest(echoCtx.Request().WithContext(ctx))

		// here a hack to make the context available in the echo request
		openapiFilter.RequestValidationInput.Request = echoCtx.Request()
		return nil
	}
	return errors.New("unsupported security scheme")
}

func ExtractUser(ctx context.Context) *supaauth.UserSchema {
	user, ok := ctx.Value(UserKey{}).(*supaauth.UserSchema)
	if !ok {
		return nil
	}
	return user
}
