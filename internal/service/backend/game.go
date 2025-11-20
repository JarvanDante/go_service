package backend

import (
	"context"
	"go-service/api/backendRoute"
)

type (
	IGame interface {
		LGames(ctx context.Context) (interface{}, error)
		LUpdateGames(ctx context.Context, req *backendRoute.UpdateGamesReq) error
	}
)

var localGame IGame

func ServiceGame() IGame {
	return localGame
}

func RegisterGame(p IGame) {
	localGame = p
}
