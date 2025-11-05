package backend

import (
	"context"
	"go-service/api/backendRoute"
)

type (
	IPublic interface {
		LLogin(ctx context.Context, req *backendRoute.PublicReq) (string, error)
	}
)

var localPublic IPublic

func ServicePublic() IPublic {
	return localPublic
}

func RegisterPublic(p IPublic) {
	localPublic = p
}
