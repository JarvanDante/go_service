package backend

import (
	"context"
	"go-service/api/backendRoute"
	"go-service/internal/model/response"
	"go-service/internal/service/backend"
)

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

//Index
/**
 * showdoc
 * @catalog 后台/会员
 * @title 会员列表
 * @description 获取会员列表
 * @method get
 * @url /app/users
 * @param token 必选 string 员工token
 * @param grade 可选 int 会员等级
 * @param level 可选 int 会员层级
 * @param status 可选 int 状态
 * @param username 可选 string 用户名
 * @param realname 可选 string 真实姓名
 * @param agent 可选 string 代理用户名
 * @param mobile 可选 string 手机号
 * @param card_no 可选 string 银行卡号
 * @param domain 可选 string 域名
 * @param start_date 可选 string 开始日期
 * @param end_date 可选 string 结束日期
 * @param charge 可选 int 是否首存
 * @param sort_field 可选 string 排序字段
 * @param sort_rule 可选 int 排序规则
 * @param page 可选 int 页数
 * @param num 可选 int 每页数量
 * @return {"code":200,"status":1,"message":"获取数据成功","data":{"list":[],"count":0,"total_users":0,"total_charge_users":0}}
 * @return_param code int 状态码
 * @return_param status int 成功/失败状态
 * @return_param message string 提示说明
 * @return_param data object 数据
 * @return_param data.list array 会员列表
 * @return_param data.count int 总条数
 * @return_param data.total_users int 总会员数
 * @return_param data.total_charge_users int 总充值会员数
 * @remark 备注
 * @number 1
 */
func (c *UserController) Index(ctx context.Context, req *backendRoute.UsersReq) (res *response.Response, err error) {

	data, err := backend.ServiceUser().LIndex(ctx, req)
	if err != nil {
		response.JsonErrCtx(ctx, err.Error())
		return
	}

	response.JsonOkCtx(ctx, data, "获取数据成功")

	return
}
