package auth

import (
	"context"

	supaauth "github.com/Razikus/underbase/supaauth"
	"github.com/getkin/kin-openapi/openapi3filter"
)

type Authorizer interface {
	VerifyToken(ctx context.Context, token string) (*supaauth.UserSchema, error)
	Validate(ctx context.Context, o *openapi3filter.AuthenticationInput) error
}
