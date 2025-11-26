package backend

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"go-service/api/backendRoute"
	daosite "go-service/internal/dao/jh_site/dao"
	entitysite "go-service/internal/dao/jh_site/model/entity"
	"go-service/internal/dao/jinhuang/model/entity"
	"go-service/internal/service/backend"
)

type sRebateRule struct {
}

func init() {
	backend.RegisterRebateRule(&sRebateRule{})
}

// LIndex 获取返水设置
func (s *sRebateRule) LIndex(ctx context.Context, req *backendRoute.RebateSettingReq) (interface{}, error) {
	site := g.RequestFromCtx(ctx).GetCtxVar("site").Val().(*entity.Site)
	siteId := site.Id

	// 查询返水规则列表
	var rules []entitysite.RebateRule
	err := daosite.RebateRule.Ctx(ctx).
		Where("site_id", siteId).
		Where("status", 1).
		Order("id ASC").
		Scan(&rules)

	if err != nil {
		return nil, err
	}

	// 组装返回数据
	ruleList := make([]g.Map, 0)
	for _, rule := range rules {
		item := g.Map{
			"id":   rule.Id,
			"name": rule.Name,
		}

		// 获取规则的条件列表
		optionList, err := s.getRuleOptionList(ctx, gconv.Int(siteId), int(rule.Id))
		if err != nil {
			return nil, err
		}
		item["option_list"] = optionList

		ruleList = append(ruleList, item)
	}

	return g.Map{"rule_list": ruleList}, nil
}

// getRuleOptionList 获取规则的条件列表
func (s *sRebateRule) getRuleOptionList(ctx context.Context, siteId int, ruleId int) ([]g.Map, error) {
	var options []entitysite.RebateRuleOption
	err := daosite.RebateRuleOption.Ctx(ctx).
		Where("site_id", siteId).
		Where("rule_id", ruleId).
		Order("id ASC").
		Scan(&options)

	if err != nil {
		return nil, err
	}

	optionList := make([]g.Map, 0)
	for _, option := range options {
		item := g.Map{
			"id":  option.Id,
			"min": fmt.Sprintf("%.2f", option.DailyValidBetMin),
			"max": fmt.Sprintf("%.2f", option.DailyValidBetMax),
		}

		// 获取游戏返水比例列表
		percentList, err := s.getGamePercentList(ctx, siteId, ruleId, int(option.Id))
		if err != nil {
			return nil, err
		}
		item["percent_list"] = percentList

		optionList = append(optionList, item)
	}

	return optionList, nil
}

// getGamePercentList 获取游戏返水比例列表
func (s *sRebateRule) getGamePercentList(ctx context.Context, siteId int, ruleId int, ruleOptionId int) ([]g.Map, error) {
	// 查询游戏返水比例
	var gamePercents []entitysite.RebateRuleOptionGame
	err := daosite.RebateRuleOptionGame.Ctx(ctx).
		Where("site_id", siteId).
		Where("rule_id", ruleId).
		Where("rule_option_id", ruleOptionId).
		Scan(&gamePercents)

	if err != nil {
		return nil, err
	}

	// 转换为map方便查找
	percentMap := make(map[int]float64)
	for _, gp := range gamePercents {
		percentMap[gp.GameId] = gp.Percent
	}

	// 获取站点游戏列表
	var games []entitysite.SiteGame
	err = daosite.SiteGame.Ctx(ctx).
		Where("site_id", siteId).
		Where("status", 1).
		Order("id ASC").
		Scan(&games)

	if err != nil {
		return nil, err
	}

	// 组装游戏返水比例列表
	percentList := make([]g.Map, 0)
	for _, game := range games {
		percent := 0.0
		if p, ok := percentMap[int(game.Id)]; ok {
			percent = p * 100 // 转换为百分比
		}

		item := g.Map{
			"id":      game.Id,
			"type":    game.Type,
			"name":    game.Name,
			"percent": percent,
		}
		percentList = append(percentList, item)
	}

	return percentList, nil
}

// LCreateRebateRule 添加返水规则
func (s *sRebateRule) LCreateRebateRule(ctx context.Context, req *backendRoute.CreateRebateRuleReq) error {
	site := g.RequestFromCtx(ctx).GetCtxVar("site").Val().(*entity.Site)
	siteId := site.Id

	// 检查规则名称是否已存在
	count, err := daosite.RebateRule.Ctx(ctx).
		Where("site_id", siteId).
		Where("name", req.Name).
		Count()

	if err != nil {
		return err
	}

	if count > 0 {
		return errors.New("规则名称已经被使用")
	}

	// 开启事务
	err = daosite.RebateRule.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		now := gtime.Now()

		// 添加返水规则
		ruleData := g.Map{
			"site_id":    siteId,
			"name":       req.Name,
			"status":     1,
			"created_at": now,
			"updated_at": now,
		}

		result, err := tx.Model("rebate_rule").Data(ruleData).Insert()
		if err != nil {
			return err
		}

		ruleId, err := result.LastInsertId()
		if err != nil {
			return err
		}

		// 添加默认条件
		optionData := g.Map{
			"site_id":             siteId,
			"rule_id":             ruleId,
			"daily_valid_bet_min": 0,
			"daily_valid_bet_max": 49999.99,
			"created_at":          now,
			"updated_at":          now,
		}

		_, err = tx.Model("rebate_rule_option").Data(optionData).Insert()
		if err != nil {
			return err
		}

		return nil
	})

	return err
}

// LUpdateRebateRule 编辑返水规则
func (s *sRebateRule) LUpdateRebateRule(ctx context.Context, req *backendRoute.UpdateRebateRuleReq) error {
	site := g.RequestFromCtx(ctx).GetCtxVar("site").Val().(*entity.Site)
	siteId := site.Id

	// 检查规则是否存在
	var rule entitysite.RebateRule
	err := daosite.RebateRule.Ctx(ctx).
		Where("id", req.Id).
		Where("site_id", siteId).
		Scan(&rule)

	if err != nil || rule.Id == 0 {
		return errors.New("没有找到返水规则")
	}

	// 检查规则名称是否已被其他规则使用
	count, err := daosite.RebateRule.Ctx(ctx).
		Where("site_id", siteId).
		Where("name", req.Name).
		WhereNot("id", req.Id).
		Count()

	if err != nil {
		return err
	}

	if count > 0 {
		return errors.New("规则名称已经被使用")
	}

	// 更新规则
	_, err = daosite.RebateRule.Ctx(ctx).
		Where("id", req.Id).
		Where("site_id", siteId).
		Data(g.Map{
			"name":       req.Name,
			"updated_at": gtime.Now(),
		}).
		Update()

	return err
}

// LDeleteRebateRule 删除返水规则
func (s *sRebateRule) LDeleteRebateRule(ctx context.Context, req *backendRoute.DeleteRebateRuleReq) error {
	site := g.RequestFromCtx(ctx).GetCtxVar("site").Val().(*entity.Site)
	siteId := site.Id

	// 检查规则是否存在
	var rule entitysite.RebateRule
	err := daosite.RebateRule.Ctx(ctx).
		Where("id", req.Id).
		Where("site_id", siteId).
		Scan(&rule)

	if err != nil || rule.Id == 0 {
		return errors.New("没有找到返水规则")
	}

	// 开启事务删除规则及其条件
	err = daosite.RebateRule.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 删除规则
		_, err := tx.Model("rebate_rule").
			Where("id", req.Id).
			Where("site_id", siteId).
			Delete()
		if err != nil {
			return err
		}

		// 删除规则条件
		_, err = tx.Model("rebate_rule_option").
			Where("site_id", siteId).
			Where("rule_id", req.Id).
			Delete()
		if err != nil {
			return err
		}

		// 删除游戏返水比例
		_, err = tx.Model("rebate_rule_option_game").
			Where("site_id", siteId).
			Where("rule_id", req.Id).
			Delete()
		if err != nil {
			return err
		}

		return nil
	})

	return err
}

// LCreateRebateOption 添加返水规则条件
func (s *sRebateRule) LCreateRebateOption(ctx context.Context, req *backendRoute.CreateRebateOptionReq) error {
	site := g.RequestFromCtx(ctx).GetCtxVar("site").Val().(*entity.Site)
	siteId := site.Id

	// 检查规则是否存在
	var rule entitysite.RebateRule
	err := daosite.RebateRule.Ctx(ctx).
		Where("id", req.Id).
		Where("site_id", siteId).
		Scan(&rule)

	if err != nil || rule.Id == 0 {
		return errors.New("没有找到返水规则")
	}

	// 获取最大的投注范围
	var maxOption entitysite.RebateRuleOption
	err = daosite.RebateRuleOption.Ctx(ctx).
		Where("site_id", siteId).
		Where("rule_id", req.Id).
		OrderDesc("daily_valid_bet_max").
		Scan(&maxOption)

	min := 0.0
	max := 49999.99

	if maxOption.Id > 0 {
		min = maxOption.DailyValidBetMax + 0.01
		max = min + 49999.99
	}

	// 添加条件
	now := gtime.Now()
	optionData := g.Map{
		"site_id":             siteId,
		"rule_id":             req.Id,
		"daily_valid_bet_min": min,
		"daily_valid_bet_max": max,
		"created_at":          now,
		"updated_at":          now,
	}

	_, err = daosite.RebateRuleOption.Ctx(ctx).Data(optionData).Insert()

	return err
}

// LUpdateRebateOption 编辑返水规则条件
func (s *sRebateRule) LUpdateRebateOption(ctx context.Context, req *backendRoute.UpdateRebateOptionReq) error {
	site := g.RequestFromCtx(ctx).GetCtxVar("site").Val().(*entity.Site)
	siteId := site.Id

	// 检查条件是否存在
	var option entitysite.RebateRuleOption
	err := daosite.RebateRuleOption.Ctx(ctx).
		Where("id", req.Id).
		Where("site_id", siteId).
		Scan(&option)

	if err != nil || option.Id == 0 {
		return errors.New("没有找到返水规则条件")
	}

	// 验证投注范围
	if req.Min >= req.Max {
		return errors.New("有效投注范围设置不合理")
	}

	// 开启事务
	err = daosite.RebateRuleOption.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 更新条件
		_, err := tx.Model("rebate_rule_option").
			Where("id", req.Id).
			Where("site_id", siteId).
			Data(g.Map{
				"daily_valid_bet_min": req.Min,
				"daily_valid_bet_max": req.Max,
				"updated_at":          gtime.Now(),
			}).
			Update()

		if err != nil {
			return err
		}

		// 如果有游戏返水比例数据，更新它们
		if req.Content != "" {
			var gameList []map[string]interface{}
			err = json.Unmarshal([]byte(req.Content), &gameList)
			if err != nil {
				return errors.New("游戏返水比例数据格式错误")
			}

			for _, game := range gameList {
				gameId := int(g.NewVar(game["id"]).Int())
				percent := g.NewVar(game["percent"]).Float64() / 100 // 转换为小数

				// 检查是否已存在
				count, err := tx.Model("rebate_rule_option_game").
					Where("site_id", siteId).
					Where("rule_id", option.RuleId).
					Where("rule_option_id", req.Id).
					Where("game_id", gameId).
					Count()

				if err != nil {
					return err
				}

				now := gtime.Now()
				if count > 0 {
					// 更新
					_, err = tx.Model("rebate_rule_option_game").
						Where("site_id", siteId).
						Where("rule_id", option.RuleId).
						Where("rule_option_id", req.Id).
						Where("game_id", gameId).
						Data(g.Map{
							"percent":    percent,
							"updated_at": now,
						}).
						Update()
				} else {
					// 插入
					_, err = tx.Model("rebate_rule_option_game").
						Data(g.Map{
							"site_id":        siteId,
							"rule_id":        option.RuleId,
							"rule_option_id": req.Id,
							"game_id":        gameId,
							"percent":        percent,
							"created_at":     now,
							"updated_at":     now,
						}).
						Insert()
				}

				if err != nil {
					return err
				}
			}
		}

		return nil
	})

	return err
}

// LDeleteRebateOption 删除返水规则条件
func (s *sRebateRule) LDeleteRebateOption(ctx context.Context, req *backendRoute.DeleteRebateOptionReq) error {
	site := g.RequestFromCtx(ctx).GetCtxVar("site").Val().(*entity.Site)
	siteId := site.Id

	// 检查条件是否存在
	var option entitysite.RebateRuleOption
	err := daosite.RebateRuleOption.Ctx(ctx).
		Where("id", req.Id).
		Where("site_id", siteId).
		Scan(&option)

	if err != nil || option.Id == 0 {
		return errors.New("没有找到返水规则条件")
	}

	// 检查是否是最后一条
	count, err := daosite.RebateRuleOption.Ctx(ctx).
		Where("site_id", siteId).
		Where("rule_id", option.RuleId).
		Count()

	if err != nil {
		return err
	}

	if count <= 1 {
		return errors.New("默认最后一条请勿删除")
	}

	// 开启事务删除条件及游戏返水比例
	err = daosite.RebateRuleOption.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 删除条件
		_, err := tx.Model("rebate_rule_option").
			Where("id", req.Id).
			Where("site_id", siteId).
			Delete()
		if err != nil {
			return err
		}

		// 删除游戏返水比例
		_, err = tx.Model("rebate_rule_option_game").
			Where("site_id", siteId).
			Where("rule_id", option.RuleId).
			Where("rule_option_id", req.Id).
			Delete()
		if err != nil {
			return err
		}

		return nil
	})

	return err
}
