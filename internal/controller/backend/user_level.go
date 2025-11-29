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
 * @catalog 后台/会员/会员层级
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

// CreateUserLevel
/**
 * showdoc
 * @catalog 后台/会员/会员层级
 * @title 新增会员层级
 * @description 新增会员层级
 * @method post
 * @url /app/create-user-level
 * @param token 必选 string 员工token
 * @param name 必选 string 层级名称
 * @param is_rebate 必选 int 是否返水
 * @param rebate_rule_id 可选 int 返水规则ID
 * @param daily_withdraw_times 必选 int 单日提款次数上限
 * @param transfer_account_list 可选 string 转账接口列表JSON
 * @param payment_account_list 可选 string 支付接口列表JSON
 * @return {"code":200,"status":1,"message":"添加会员层级成功","data":null}
 * @return_param code int 状态码
 * @return_param status int 成功/失败状态
 * @return_param message string 提示说明
 * @remark 备注
 * @number 2
 */
func (c *UserLevelController) CreateUserLevel(ctx context.Context, req *backendRoute.CreateUserLevelReq) (res *response.Response, err error) {
	err = backend.ServiceUserLevel().LCreateUserLevel(ctx, req)
	if err != nil {
		response.JsonErrCtx(ctx, err.Error())
		return
	}

	response.JsonOkCtx(ctx, nil, "添加会员层级成功")
	return
}

// GetUpdateUserLevel
/**
 * showdoc
 * @catalog 后台/会员/会员层级
 * @title 获取会员层级信息用于编辑
 * @description 获取会员层级信息用于编辑
 * @method get
 * @url /app/update-user-level
 * @param token 必选 string 员工token
 * @param id 必选 int 层级ID
 * @return {"code":200,"status":1,"message":"获取数据成功","data":{}}
 * @return_param code int 状态码
 * @return_param status int 成功/失败状态
 * @return_param message string 提示说明
 * @return_param data object 层级信息
 * @remark 备注
 * @number 3
 */
func (c *UserLevelController) GetUpdateUserLevel(ctx context.Context, req *backendRoute.GetUpdateUserLevelReq) (res *response.Response, err error) {
	data, err := backend.ServiceUserLevel().LGetUpdateUserLevel(ctx, req)
	if err != nil {
		response.JsonErrCtx(ctx, err.Error())
		return
	}

	response.JsonOkCtx(ctx, data, "获取数据成功")
	return
}

// UpdateUserLevel
/**
 * showdoc
 * @catalog 后台/会员/会员层级
 * @title 编辑会员层级
 * @description 编辑会员层级
 * @method post
 * @url /app/update-user-level
 * @param token 必选 string 员工token
 * @param id 必选 int 层级ID
 * @param name 必选 string 层级名称
 * @param is_rebate 必选 int 是否返水
 * @param rebate_rule_id 可选 int 返水规则ID
 * @param daily_withdraw_times 必选 int 单日提款次数上限
 * @param transfer_account_list 可选 string 转账接口列表JSON
 * @param payment_account_list 可选 string 支付接口列表JSON
 * @return {"code":200,"status":1,"message":"编辑成功","data":null}
 * @return_param code int 状态码
 * @return_param status int 成功/失败状态
 * @return_param message string 提示说明
 * @remark 备注
 * @number 4
 */
func (c *UserLevelController) UpdateUserLevel(ctx context.Context, req *backendRoute.UpdateUserLevelReq) (res *response.Response, err error) {
	err = backend.ServiceUserLevel().LUpdateUserLevel(ctx, req)
	if err != nil {
		response.JsonErrCtx(ctx, err.Error())
		return
	}

	response.JsonOkCtx(ctx, nil, "编辑成功")
	return
}

// DeleteUserLevel
/**
 * showdoc
 * @catalog 后台/会员/会员层级
 * @title 删除会员层级
 * @description 删除会员层级
 * @method post
 * @url /app/delete-user-level
 * @param token 必选 string 员工token
 * @param id 必选 int 层级ID
 * @return {"code":200,"status":1,"message":"删除成功","data":null}
 * @return_param code int 状态码
 * @return_param status int 成功/失败状态
 * @return_param message string 提示说明
 * @remark 备注
 * @number 5
 */
func (c *UserLevelController) DeleteUserLevel(ctx context.Context, req *backendRoute.DeleteUserLevelReq) (res *response.Response, err error) {
	err = backend.ServiceUserLevel().LDeleteUserLevel(ctx, req)
	if err != nil {
		response.JsonErrCtx(ctx, err.Error())
		return
	}

	response.JsonOkCtx(ctx, nil, "删除成功")
	return
}
