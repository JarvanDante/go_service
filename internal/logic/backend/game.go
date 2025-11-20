package backend

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"go-service/api/backendRoute"
	daosite "go-service/internal/dao/jh_site/dao"
	entitysite "go-service/internal/dao/jh_site/model/entity"
	"go-service/internal/dao/jinhuang/dao"
	"go-service/internal/dao/jinhuang/model/entity"
	"go-service/internal/service/backend"
)

type sGame struct {
}

func init() {
	backend.RegisterGame(&sGame{})
}

func (s *sGame) LGames(ctx context.Context) (interface{}, error) {
	site := g.RequestFromCtx(ctx).GetCtxVar("site").Val().(*entity.Site)

	siteId := site.Id

	// 先查 game 表所有可用游戏
	var gameList []*entity.Game
	err := dao.Game.Ctx(ctx).
		Where(dao.Game.Columns().Status, 1).
		Scan(&gameList)
	if err != nil {
		return nil, err
	}

	// 查 site_game
	var siteGameList []*entitysite.SiteGame
	err = daosite.SiteGame.Ctx(ctx).
		Where(g.Map{
			"site_id":      siteId,
			"is_available": 1,
		}).
		Order("type asc").
		Order("sort asc").
		Scan(&siteGameList)
	if err != nil {
		return nil, err
	}

	// 拼装数据并按 type 分组
	data := make(map[int][]map[string]interface{})

	for _, sg := range siteGameList {
		for _, game := range gameList {
			if sg.GameId != gconv.Int(game.Id) {
				continue
			}

			item := map[string]interface{}{
				"id":     game.Id,
				"name":   game.Name,
				"status": sg.Status,
				"sort":   sg.Sort,
			}

			data[game.Type] = append(data[game.Type], item)
		}
	}

	return data, nil
}

func (s *sGame) LUpdateGames(ctx context.Context, req *backendRoute.UpdateGamesReq) error {
	site := g.RequestFromCtx(ctx).GetCtxVar("site").Val().(*entity.Site)
	siteId := site.Id

	for _, gitem := range req.Data {
		_, err := daosite.SiteGame.Ctx(ctx).
			Data(g.Map{
				"status": gitem.Status,
				"sort":   gitem.Sort,
			}).
			Where(g.Map{
				"site_id":      siteId,
				"game_id":      gitem.Id,
				"is_available": 1,
			}).
			Update()

		if err != nil {
			return err
		}
	}

	return nil
}
