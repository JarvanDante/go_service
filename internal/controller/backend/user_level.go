package backend

import (
	"context"
	"go-service/api/backendRoute"
	"go-service/internal/model/response"
	"go-service/internal/service/backend"
)

type UserLevelController struct{}

func NewUserLevelController() *UserLevelController {
	return &UserLevelController{}
}

// Index
/**
 * showdoc
 * @catalog 后台/会员
 * @title 获取会员层级列表
 * @description 获取会员层级列表
 * @method get
 * @url /app/user-levels
 * @param token 必选 string 员工token
 * @return {"code":200,"status":1,"message":"获取数据成功","data":[]}
 * @return_param code int 状态码
 * @return_param status int 成功/失败状态
 * @return_param message string 提示说明
 * @return_param data array 会员层级列表
 * @remark 备注
 * @number 1
 */
func (c *UserLevelController) Index(ctx context.Context, req *backendRoute.UserLevelsReq) (res *response.Response, err error) {
	data, err := backend.ServiceUserLevel().LIndex(ctx, req)
	if err != nil {
		response.JsonErrCtx(ctx, err.Error())
		return
	}

	response.JsonOkCtx(ctx, data, "获取数据成功")
	return
}
