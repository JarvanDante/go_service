package backend

import (
	"context"
	"go-service/api/backendRoute"
	"go-service/internal/model/response"
	"go-service/internal/service/backend"
)

type RebateRuleController struct{}

func NewRebateRuleController() *RebateRuleController {
	return &RebateRuleController{}
}

// Index
/**
 * showdoc
 * @catalog 后台/系统/返水设置
 * @title 获取返水设置
 * @description 获取返水设置
 * @method get
 * @url /app/rebate-setting
 * @param token 必选 string 员工token
 * @return {"code":200,"status":1,"message":"获取数据成功","data":{"rule_list":[]}}
 * @return_param code int 状态码
 * @return_param status int 成功/失败状态
 * @return_param message string 提示说明
 * @return_param data object 返水规则列表
 * @remark 备注
 * @number 1
 */
func (c *RebateRuleController) Index(ctx context.Context, req *backendRoute.RebateSettingReq) (res *response.Response, err error) {
	data, err := backend.ServiceRebateRule().LIndex(ctx, req)
	if err != nil {
		response.JsonErrCtx(ctx, err.Error())
		return
	}

	response.JsonOkCtx(ctx, data, "获取数据成功")
	return
}

// CreateRebateRule
/**
 * showdoc
 * @catalog 后台/系统/返水设置
 * @title 添加返水规则
 * @description 添加返水规则
 * @method post
 * @url /app/create-rebate-rule
 * @param token 必选 string 员工token
 * @param name 必选 string 规则名称
 * @return {"code":200,"status":1,"message":"添加成功","data":null}
 * @return_param code int 状态码
 * @return_param status int 成功/失败状态
 * @return_param message string 提示说明
 * @remark 备注
 * @number 2
 */
func (c *RebateRuleController) CreateRebateRule(ctx context.Context, req *backendRoute.CreateRebateRuleReq) (res *response.Response, err error) {
	err = backend.ServiceRebateRule().LCreateRebateRule(ctx, req)
	if err != nil {
		response.JsonErrCtx(ctx, err.Error())
		return
	}

	response.JsonOkCtx(ctx, nil, "添加成功")
	return
}

// UpdateRebateRule
/**
 * showdoc
 * @catalog 后台/系统/返水设置
 * @title 编辑返水规则
 * @description 编辑返水规则
 * @method post
 * @url /app/update-rebate-rule
 * @param token 必选 string 员工token
 * @param id 必选 int 规则ID
 * @param name 必选 string 规则名称
 * @return {"code":200,"status":1,"message":"编辑成功","data":null}
 * @return_param code int 状态码
 * @return_param status int 成功/失败状态
 * @return_param message string 提示说明
 * @remark 备注
 * @number 3
 */
func (c *RebateRuleController) UpdateRebateRule(ctx context.Context, req *backendRoute.UpdateRebateRuleReq) (res *response.Response, err error) {
	err = backend.ServiceRebateRule().LUpdateRebateRule(ctx, req)
	if err != nil {
		response.JsonErrCtx(ctx, err.Error())
		return
	}

	response.JsonOkCtx(ctx, nil, "编辑成功")
	return
}

// DeleteRebateRule
/**
 * showdoc
 * @catalog 后台/系统/返水设置
 * @title 删除返水规则
 * @description 删除返水规则
 * @method post
 * @url /app/delete-rebate-rule
 * @param token 必选 string 员工token
 * @param id 必选 int 规则ID
 * @return {"code":200,"status":1,"message":"删除成功","data":null}
 * @return_param code int 状态码
 * @return_param status int 成功/失败状态
 * @return_param message string 提示说明
 * @remark 备注
 * @number 4
 */
func (c *RebateRuleController) DeleteRebateRule(ctx context.Context, req *backendRoute.DeleteRebateRuleReq) (res *response.Response, err error) {
	err = backend.ServiceRebateRule().LDeleteRebateRule(ctx, req)
	if err != nil {
		response.JsonErrCtx(ctx, err.Error())
		return
	}

	response.JsonOkCtx(ctx, nil, "删除成功")
	return
}

// CreateRebateOption
/**
 * showdoc
 * @catalog 后台/系统/返水设置
 * @title 添加返水规则条件
 * @description 添加返水规则条件
 * @method post
 * @url /app/create-rebate-option
 * @param token 必选 string 员工token
 * @param id 必选 int 规则ID
 * @return {"code":200,"status":1,"message":"添加成功","data":null}
 * @return_param code int 状态码
 * @return_param status int 成功/失败状态
 * @return_param message string 提示说明
 * @remark 备注
 * @number 5
 */
func (c *RebateRuleController) CreateRebateOption(ctx context.Context, req *backendRoute.CreateRebateOptionReq) (res *response.Response, err error) {
	err = backend.ServiceRebateRule().LCreateRebateOption(ctx, req)
	if err != nil {
		response.JsonErrCtx(ctx, err.Error())
		return
	}

	response.JsonOkCtx(ctx, nil, "添加成功")
	return
}

// UpdateRebateOption
/**
 * showdoc
 * @catalog 后台/系统/返水设置
 * @title 编辑返水规则条件
 * @description 编辑返水规则条件
 * @method post
 * @url /app/update-rebate-option
 * @param token 必选 string 员工token
 * @param id 必选 int 条件ID
 * @param min 必选 float 最小金额
 * @param max 必选 float 最大金额
 * @param content 可选 string 游戏返水比例JSON
 * @return {"code":200,"status":1,"message":"保存成功","data":null}
 * @return_param code int 状态码
 * @return_param status int 成功/失败状态
 * @return_param message string 提示说明
 * @remark 备注
 * @number 6
 */
func (c *RebateRuleController) UpdateRebateOption(ctx context.Context, req *backendRoute.UpdateRebateOptionReq) (res *response.Response, err error) {
	err = backend.ServiceRebateRule().LUpdateRebateOption(ctx, req)
	if err != nil {
		response.JsonErrCtx(ctx, err.Error())
		return
	}

	response.JsonOkCtx(ctx, nil, "保存成功")
	return
}

// DeleteRebateOption
/**
 * showdoc
 * @catalog 后台/系统/返水设置
 * @title 删除返水规则条件
 * @description 删除返水规则条件
 * @method post
 * @url /app/delete-rebate-option
 * @param token 必选 string 员工token
 * @param id 必选 int 条件ID
 * @return {"code":200,"status":1,"message":"删除成功","data":null}
 * @return_param code int 状态码
 * @return_param status int 成功/失败状态
 * @return_param message string 提示说明
 * @remark 备注
 * @number 7
 */
func (c *RebateRuleController) DeleteRebateOption(ctx context.Context, req *backendRoute.DeleteRebateOptionReq) (res *response.Response, err error) {
	err = backend.ServiceRebateRule().LDeleteRebateOption(ctx, req)
	if err != nil {
		response.JsonErrCtx(ctx, err.Error())
		return
	}

	response.JsonOkCtx(ctx, nil, "删除成功")
	return
}
