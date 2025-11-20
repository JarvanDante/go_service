package backend

import (
	"context"
	"go-service/api/backendRoute"
	"go-service/internal/model/response"
	"go-service/internal/service/backend"
)

type SiteController struct{}

func NewSiteController() *SiteController {
	return &SiteController{}
}

//BasicSetting
/**
 * showdoc
 * @catalog 后台/系统/全局设置
 * @title 获取站点基本信息
 * @description 获取站点基本信息
 * @method get
 * @url /app/basic-setting
 * @param token 必选 string 员工token
 * @return {"code":200,"status":1,"message":"获取数据成功","data":{"agent_register_url":"http://103.72.166.56:1002/reg/phone01","agent_url":"http://103.72.166.56:1002/agent","balance":"10000.0000","balance_reset":"12801296.0000","close_reason":"您好，系统正在维护中，估计需要1个小时左右，给您造成不便，敬请谅解！","default_agent_id":0,"default_agent_name":"","is_close":0,"max_withdraw":99999999,"min_withdraw":1,"mobile_url":"http://dev.testpangu.com/kf/mobileBet/","register_time_interval":1,"service_url":"http://szzero.livechatvalue.com/chat/chatClient/chatbox.jsp?companyID=86102\u0026configID=55643\u0026jid=28","switch_register":1}}
 * @return_param code int 状态码
 * @return_param status int 成功/失败状态
 * @return_param message string 提示说明
 * @return_param data array 数组
 * @return_param data.code string 应用代码
 * @return_param data.name string 应用名称
 * @return_param data.register_time_interval int 同一IP重复注册次数
 * @return_param data.switch_register int 是否开放注册
 * @return_param data.is_close int 是否关站
 * @return_param data.close_reason string 关站提示语
 * @return_param data.service_url string 客服链接
 * @return_param data.agent_url string 代理链接
 * @return_param data.mobile_url string APP下载地址
 * @return_param data.agent_register_url string 代理推广地址
 * @return_param data.min_withdraw string 默认单笔最小提现金额
 * @return_param data.max_withdraw string 默认单笔最大提现金额
 * @return_param data.game_free_play string 游戏试玩地址
 * @return_param data.mobile_logo string 手机logo地址
 * @return_param data.default_agent_id string 默认代理ID
 * @return_param data.default_agent_name string 默认代理名称
 * @return_param data.balance string 默认额度
 * @return_param data.balance_reset string 可用额度
 * @remark 备注
 * @number 1
 */
func (c *SiteController) BasicSetting(ctx context.Context, req *backendRoute.BasicSettingReq) (res *response.Response, err error) {

	data, err := backend.ServiceSite().LBasicSetting(ctx)
	if err != nil {
		return nil, err
	}

	response.JsonOkCtx(ctx, data, "获取数据成功")

	return
}

//UpdateBasicSetting
/**
 * showdoc
 * @catalog 后台/系统/全局设置
 * @title 修改站点基本信息
 * @description 修改站点基本信息
 * @method post
 * @url /app/update-basic-setting
 * @param token 必选 string 员工token
 * @return {"code":200,"status":1,"message":"设置成功","data":null}
 * @return_param code int 状态码
 * @return_param status int 成功/失败状态
 * @return_param message string 提示说明
 * @return_param data array 数组
 * @remark 备注
 * @number 1
 */
func (c *SiteController) UpdateBasicSetting(ctx context.Context, req *backendRoute.UpdateBasicSettingReq) (res *response.Response, err error) {

	err = backend.ServiceSite().LUpdateBasicSetting(ctx, req)
	if err != nil {
		return nil, err
	}

	response.JsonOkCtx(ctx, nil, "设置成功")

	return
}

//RegisterSetting
/**
 * showdoc
 * @catalog 后台/系统/全局设置
 * @title 获取会员注册设置
 * @description 获取会员注册设置
 * @method get
 * @url /app/register-setting
 * @param token 必选 string 员工token
 * @return {"code":200,"status":1,"message":"获取数据成功","data":[{"id":1,"type":1,"name":"用户名","field_name":"username","display":1,"required":1},{"id":2,"type":2,"name":"密码","field_name":"password","display":1,"required":1},{"id":3,"type":3,"name":"验证码","field_name":"code","display":1,"required":1},{"id":4,"type":4,"name":"资金密码","field_name":"pay_password","display":0,"required":0},{"id":5,"type":5,"name":"密保问题","field_name":"safe_question","display":0,"required":0},{"id":6,"type":6,"name":"真实姓名","field_name":"realname","display":0,"required":0},{"id":7,"type":7,"name":"手机","field_name":"mobile","display":1,"required":0},{"id":8,"type":8,"name":"邮箱","field_name":"email","display":0,"required":0},{"id":9,"type":9,"name":"手机验证码","field_name":"phone_code","display":1,"required":1}]}
 * @return_param code int 状态码
 * @return_param status int 成功/失败状态
 * @return_param message string 提示说明
 * @return_param data array 数组
 * @return_param data.id int 数据id
 * @return_param data.type int 类型
 * @return_param data.name string 中文名称
 * @return_param data.field_name string 英文名称
 * @return_param data.display int 是否显示
 * @return_param data.required int 是否必需
 * @remark 备注
 * @number 1
 */
func (c *SiteController) RegisterSetting(ctx context.Context, req *backendRoute.RegisterSettingReq) (res *response.Response, err error) {

	data, err := backend.ServiceSite().LRegisterSetting(ctx)
	if err != nil {
		return nil, err
	}

	response.JsonOkCtx(ctx, data, "获取数据成功")

	return
}

//UpdateRegisterSetting
/**
 * showdoc
 * @catalog 后台/系统/全局设置
 * @title 修改会员注册设置
 * @description 修改会员注册设置
 * @method post
 * @url /app/update-register-setting
 * @param token 必选 string 员工token
 * @param type 必选 string 类型
 * @param display 必选 int 是否显示
 * @param required 必选 int 是否必须
 * @return {"code":200,"status":1,"message":"设置成功","data":null}
 * @return_param code int 状态码
 * @return_param status int 成功/失败状态
 * @return_param message string 提示说明
 * @return_param data array 数组
 * @remark 备注
 * @number 1
 */
func (c *SiteController) UpdateRegisterSetting(ctx context.Context, req *backendRoute.UpdateRegisterSettingReq) (res *response.Response, err error) {

	err = backend.ServiceSite().LUpdateRegisterSetting(ctx, req)
	if err != nil {
		return nil, err
	}

	response.JsonOkCtx(ctx, nil, "设置成功")

	return
}
