package api

import (
	"context"

	"github.com/Razikus/underbase/pkg/auth"
	"github.com/Razikus/underbase/server"
)

type API struct {
	*HelloHandler
}

func NewAPI(helloHandler *HelloHandler) *API {
	return &API{
		HelloHandler: helloHandler,
	}
}

type HelloHandler struct {
}

func (a *HelloHandler) GetHello(ctx context.Context, request server.GetHelloRequestObject) (server.GetHelloResponseObject, error) {
	user := auth.ExtractUser(ctx)
	if user == nil {
		return server.GetHello401JSONResponse{}, nil
	}
	idOf := *user.Id
	message := "Hello " + idOf.String()
	return server.GetHello200JSONResponse{
		Message: &message,
	}, nil
}
