package backend

import (
	"context"
	"go-service/api/backendRoute"
	"go-service/internal/service/backend"
)

type sPublic struct {
}

func init() {
	backend.RegisterPublic(&sPublic{})
}

func (s *sPublic) LLogin(ctx context.Context, req *backendRoute.PublicReq) (string, error) {

	return "token", nil
}
