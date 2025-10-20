package frontend

import (
	"context"
	"go-service/internal/service/frontend"
)

type sDemo struct{}

func init() {
	frontend.RegisterDemo(&sDemo{})
}

func (s *sDemo) Index(ctx context.Context) (string, error) {
	return "Access Denied", nil
}
