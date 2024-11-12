package service

import (
	"context"
	"library-synapsis/model/web/create"
	"library-synapsis/model/web/response"
)

type LoginService interface {
	Login(ctx context.Context, loginRequest create.LoginCreateRequest) (response.LoginResponse, error)
}
