package backend

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"go-service/api/backendRoute"
	daobalance "go-service/internal/dao/jh_balance/dao"
	daosite "go-service/internal/dao/jh_site/dao"
	entitysite "go-service/internal/dao/jh_site/model/entity"
	"go-service/internal/dao/jinhuang/model/entity"

	entitynew "go-service/internal/model/entity"
	"go-service/internal/service/backend"
)

type sSite struct {
}

func init() {
	backend.RegisterSite(&sSite{})
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
		"site_name":              site.Name,
		"site_code":              site.Code,
		"register_time_interval": model["register_time_interval"],
		"switch_register":        model["switch_register"],
		"is_close":               model["is_close"],
		"close_reason":           IfEmpty(model["close_reason"], msg),
		"service_url":            model["url_service"],
		"agent_url":              model["url_agent_pc"],
		"mobile_url":             model["url_mobile"],
		"mobile_logo":            model["mobile_logo"],
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

//LUpdateBasicSetting
/**
 * @desc：修改站点基本信息
 * @param ctx
 * @param req
 * @return error
 * @author : Carson
 */
func (s *sSite) LUpdateBasicSetting(ctx context.Context, req *backendRoute.UpdateBasicSettingReq) error {
	//从 ctx 中获取 site
	site := g.RequestFromCtx(ctx).GetCtxVar("site").Val().(*entity.Site)

	_, err := daosite.SiteConfig.Ctx(ctx).
		Data(g.Map{
			"register_time_interval": req.RegisterTimeInterval,
			"switch_register":        req.SwitchRegister,
			"is_close":               req.IsClose,
			"agent_url":              req.AgentUrl,
			"agent_register_url":     req.AgentRegisterUrl,
			"service_url":            req.ServiceUrl,
			"mobile_url":             req.MobileUrl,
			"min_withdraw":           req.MinWithdraw,
			"max_withdraw":           req.MaxWithdraw,
			"mobile_logo":            req.MobileLogo,
		}).
		Where("site_id", site.Id).
		Update()
	if err != nil {
		return err
	}

	return nil
}

//LRegisterSetting
/**
 * @desc：获取会员注册设置
 * @param ctx
 * @return interface{}
 * @return error
 * @author : Carson
 */
func (s *sSite) LRegisterSetting(ctx context.Context) (interface{}, error) {

	site := g.RequestFromCtx(ctx).GetCtxVar("site").Val().(*entity.Site)

	fields := daosite.SiteRegister.Columns()

	var list []*entitynew.SiteRegister
	where := map[string]interface{}{}
	where["site_id"] = site.Id

	err := daosite.SiteRegister.Ctx(ctx).
		Where(where).
		Fields(fields.Id, fields.Type, fields.Name, fields.FieldName, fields.Display, fields.Required).
		Scan(&list)

	if err != nil {
		return nil, err
	}

	return list, nil

}

func (s *sSite) LUpdateRegisterSetting(ctx context.Context, req *backendRoute.UpdateRegisterSettingReq) error {
	//从 ctx 中获取 site
	site := g.RequestFromCtx(ctx).GetCtxVar("site").Val().(*entity.Site)

	siteId := site.Id

	fields := daosite.SiteRegister.Columns()

	// 是否存在记录
	var model entitysite.SiteRegister
	err := daosite.SiteRegister.Ctx(ctx).
		Where(fields.SiteId, siteId).
		Where(fields.Type, req.Type).
		Scan(&model)
	if err != nil {
		return err
	}

	data := g.Map{
		fields.SiteId:   siteId,
		fields.Type:     req.Type,
		fields.Display:  req.Display,
		fields.Required: req.Required,
	}

	if model.Id > 0 {
		// 更新
		_, err = daosite.SiteRegister.Ctx(ctx).
			Data(data).
			Where(fields.Id, model.Id).
			Update()
	} else {
		// 插入
		return gerror.New("配置不存在")
	}

	return err
}
