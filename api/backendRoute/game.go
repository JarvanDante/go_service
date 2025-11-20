package backendRoute

import "github.com/gogf/gf/v2/frame/g"

type GamesReq struct {
	g.Meta `path:"/games" method:"get" summary:"获取游戏列表"`
}

type UpdateGamesReq struct {
	g.Meta `path:"/update-games" method:"post" summary:"修改站点游戏"`
	Data   []GameItem `json:"data"`
}

type GameItem struct {
	Id     int `json:"id"`
	Status int `json:"status"`
	Sort   int `json:"sort"`
}
