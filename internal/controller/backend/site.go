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
