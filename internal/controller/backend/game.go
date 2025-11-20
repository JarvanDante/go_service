package backend

import (
	"context"
	"go-service/api/backendRoute"
	"go-service/internal/model/response"
	"go-service/internal/service/backend"
)

type GameController struct{}

func NewGameController() *GameController {
	return &GameController{}
}

//Games
/**
 * showdoc
 * @catalog 后台/游戏管理
 * @title 获取游戏管理列表
 * @description 获取游戏管理列表
 * @method get
 * @url /app/games
 * @param token 必选 string 员工token
 * @return {"code":200,"status":1,"message":"获取数据成功","data":{"1":[{"id":1,"name":"AG体育","sort":1,"status":1},{"id":2,"name":"沙巴体育","sort":2,"status":1}],"2":[{"id":3,"name":"BBIN彩票","sort":1,"status":1}],"3":[{"id":4,"name":"AG国际厅","sort":1,"status":1},{"id":5,"name":"OG视讯","sort":2,"status":1},{"id":6,"name":"EBET视讯","sort":3,"status":1},{"id":7,"name":"HG视讯","sort":7,"status":1}],"4":[{"id":9,"name":"AG电游","sort":1,"status":1},{"id":10,"name":"PT电游","sort":2,"status":1},{"id":13,"name":"AG捕鱼","sort":3,"status":1},{"id":14,"name":"PT捕鱼","sort":4,"status":1}],"5":[{"id":11,"name":"LY棋牌","sort":1,"status":1},{"id":12,"name":"KY棋牌","sort":2,"status":1}]}}
 * @return_param code int 状态码
 * @return_param status int 成功/失败状态
 * @return_param message string 提示说明
 * @return_param data array 数组
 * @return_param data.id int ID
 * @return_param data.name string 游戏名称
 * @return_param data.sort int 游戏排序
 * @return_param data.status int 状态
 * @remark 备注
 * @number 1
 */
func (c *GameController) Games(ctx context.Context, req *backendRoute.GamesReq) (res *response.Response, err error) {

	data, err := backend.ServiceGame().LGames(ctx)
	if err != nil {
		return nil, err
	}

	response.JsonOkCtx(ctx, data, "获取数据成功")

	return
}

//UpdateGames
/**
 * showdoc
 * @catalog 后台/游戏管理
 * @title 修改游戏开关
 * @description 修改游戏开关
 * @method post
 * @url /app/update-games
 * @param token 必选 string 员工token
 * @param data 必选 string json如:[{"id":1,"status":0,"sort":0},{"id":2,"status":1,"sort":0}]
 * @return {"code":200,"status":1,"message":"获取数据成功","data":{"1":[{"id":1,"name":"AG体育","sort":1,"status":1},{"id":2,"name":"沙巴体育","sort":2,"status":1}],"2":[{"id":3,"name":"BBIN彩票","sort":1,"status":1}],"3":[{"id":4,"name":"AG国际厅","sort":1,"status":1},{"id":5,"name":"OG视讯","sort":2,"status":1},{"id":6,"name":"EBET视讯","sort":3,"status":1},{"id":7,"name":"HG视讯","sort":7,"status":1}],"4":[{"id":9,"name":"AG电游","sort":1,"status":1},{"id":10,"name":"PT电游","sort":2,"status":1},{"id":13,"name":"AG捕鱼","sort":3,"status":1},{"id":14,"name":"PT捕鱼","sort":4,"status":1}],"5":[{"id":11,"name":"LY棋牌","sort":1,"status":1},{"id":12,"name":"KY棋牌","sort":2,"status":1}]}}
 * @return_param code int 状态码
 * @return_param status int 成功/失败状态
 * @return_param message string 提示说明
 * @return_param data array 数组
 * @remark 备注
 * @number 1
 */
func (c *GameController) UpdateGames(ctx context.Context, req *backendRoute.UpdateGamesReq) (res *response.Response, err error) {

	err = backend.ServiceGame().LUpdateGames(ctx, req)
	if err != nil {
		return nil, err
	}

	response.JsonOkCtx(ctx, nil, "修改成功")

	return
}
