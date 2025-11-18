package backend

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	daobalance "go-service/internal/dao/jh_balance/dao"
	daosite "go-service/internal/dao/jh_site/dao"
	"go-service/internal/dao/jinhuang/model/entity"
	"go-service/internal/service/backend"
)

type sSite struct {
}

func init() {
	backend.RegisterSiten(&sSite{})
}

//LBasicSetting
/**
 * @desc：获取站点基本配置信息
 * @param ctx
 * @return map[string]interface{}
 * @return error
 * @author : Carson
 */
func (s *sSite) LBasicSetting(ctx context.Context) (map[string]interface{}, error) {

	//从 ctx 中获取 site
	site := g.RequestFromCtx(ctx).GetCtxVar("site").Val().(*entity.Site)

	cfg, err := daosite.SiteConfig.Ctx(ctx).Where("site_id", site.Id).One()
	if err != nil {
		return nil, err
	}
	if cfg == nil {
		return nil, nil
	}

	model := cfg.Map()

	msg := "您好，系统正在维护中，估计需要1个小时左右，给您造成不便，敬请谅解！"

	data := map[string]interface{}{
		"register_time_interval": model["register_time_interval"],
		"switch_register":        model["switch_register"],
		"is_close":               model["is_close"],
		"close_reason":           IfEmpty(model["close_reason"], msg),
		"service_url":            model["url_service"],
		"agent_url":              model["url_agent_pc"],
		"mobile_url":             model["url_mobile"],
		"agent_register_url":     model["url_agent_register"],
		"min_withdraw":           model["min_withdraw"],
		"max_withdraw":           model["max_withdraw"],
		"default_agent_id":       0,
		"default_agent_name":     "",
	}

	// 默认代理
	if model["default_agent_id"] != nil {
		agent, _ := daosite.Agent.Ctx(ctx).Where("id", model["default_agent_id"]).One()
		if agent != nil {
			data["default_agent_id"] = agent["id"]
			data["default_agent_name"] = agent["username"]
		}
	}

	// 站点余额
	siteBalance, _ := daobalance.BalanceSite.Ctx(ctx).Where("site_id", site.Id).One()
	if siteBalance != nil {
		data["balance"] = siteBalance["balance_default"]
		data["balance_reset"] = siteBalance["balance_available"]
	}

	return data, nil

}

func IfEmpty(v interface{}, def string) string {
	if v == nil || v == "" {
		return def
	}
	return v.(string)
}
