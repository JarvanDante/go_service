package backend

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"go-service/api/backendRoute"
	daobalance "go-service/internal/dao/jh_balance/dao"
	daosite "go-service/internal/dao/jh_site/dao"
	entitysite "go-service/internal/dao/jh_site/model/entity"
	"go-service/internal/dao/jinhuang/model/entity"
	"go-service/internal/service/backend"
)

type sUser struct {
}

func init() {
	backend.RegisterUser(&sUser{})
}

//LIndex
/**
 * @desc：获取会员列表
 * @param ctx
 * @param req
 * @return res
 * @return err
 * @author : Carson
 */
func (s *sUser) LIndex(ctx context.Context, req *backendRoute.UsersReq) (interface{}, error) {
	site := g.RequestFromCtx(ctx).GetCtxVar("site").Val().(*entity.Site)
	siteId := site.Id

	// 构建查询条件
	model := daosite.User.Ctx(ctx).Where("user.site_id", siteId)

	// 会员等级
	if req.Grade != nil && *req.Grade > 0 {
		model = model.Where("user.grade_id", *req.Grade)
	}

	// 会员层级
	if req.Level != nil && *req.Level > 0 {
		model = model.Where("user.level_id", *req.Level)
	}

	// 状态
	if req.Status != nil {
		model = model.Where("user.status", *req.Status)
	}

	// 用户名
	if req.Username != nil && *req.Username != "" {
		model = model.Where("user.username", *req.Username)
	}

	// 真实姓名
	if req.Realname != nil && *req.Realname != "" {
		model = model.Where("user.realname", *req.Realname)
	}

	// 手机号
	if req.Mobile != nil && *req.Mobile != "" {
		model = model.Where("user.mobile", *req.Mobile)
	}

	// 代理用户名（需要关联查询 agent 表）
	if req.Agent != nil && *req.Agent != "" {
		// 先查询代理ID
		var agent *entitysite.Agent
		err := daosite.Agent.Ctx(ctx).
			Where("site_id", siteId).
			Where("username", *req.Agent).
			Fields("id").
			Scan(&agent)

		if err != nil {
			return g.Map{
				"list":               []interface{}{},
				"count":              0,
				"total_users":        0,
				"total_charge_users": 0,
			}, nil
		}

		// 如果找不到代理，返回空列表
		if agent == nil || agent.Id == 0 {
			return g.Map{
				"list":               []interface{}{},
				"count":              0,
				"total_users":        0,
				"total_charge_users": 0,
			}, nil
		}

		// 添加代理ID筛选条件
		model = model.Where("user.agent_id", agent.Id)
	}

	// 注册时间范围
	if req.StartDate != nil && *req.StartDate != "" {
		model = model.Where("user.created_at >= ?", *req.StartDate)
	}

	if req.EndDate != nil && *req.EndDate != "" {
		model = model.Where("user.created_at <= ?", *req.EndDate)
	}

	// 域名筛选
	if req.Domain != nil && *req.Domain != "" {
		model = model.WhereLike("user.register_url", "%"+*req.Domain+"%")
	}

	// 是否首存
	if req.Charge != nil && *req.Charge > 0 {
		model = model.Where("user.pay_times >= ?", 1)
	}

	// 银行卡号筛选
	if req.CardNo != nil && *req.CardNo != "" {
		// 先查询符合银行卡号的用户ID
		var userBanks []struct {
			UserId int `json:"user_id"`
		}
		err := daosite.UserBank.Ctx(ctx).
			Where("card_no", *req.CardNo).
			Fields("user_id").
			Scan(&userBanks)

		if err != nil || len(userBanks) == 0 {
			return g.Map{
				"list":               []interface{}{},
				"count":              0,
				"total_users":        0,
				"total_charge_users": 0,
			}, nil
		}

		userIds := make([]interface{}, 0)
		for _, ub := range userBanks {
			userIds = append(userIds, ub.UserId)
		}
		model = model.WhereIn("id", userIds)
	}

	// 获取总数
	count, err := model.Count()
	if err != nil {
		return nil, err
	}

	// 分页参数
	page := 1
	if req.Page != nil && *req.Page > 0 {
		page = *req.Page
	}

	size := 20
	if req.Num != nil && *req.Num > 0 {
		size = *req.Num
	}

	// 排序字段
	orderField := "created_at"
	if req.SortField != nil && *req.SortField != "" {
		orderField = s.getOrderByField(*req.SortField)
	}

	// 排序规则
	orderRule := "DESC"
	if req.SortRule != nil && *req.SortRule == 0 {
		orderRule = "ASC"
	}

	// 执行查询 - 只查询 user 表
	var users []entitysite.User
	err = model.Page(page, size).
		Order(orderField + " " + orderRule).
		Scan(&users)

	if err != nil {
		return nil, err
	}

	// 如果没有数据，直接返回
	if len(users) == 0 {
		return g.Map{
			"list":               []interface{}{},
			"count":              count,
			"total_users":        0,
			"total_charge_users": 0,
		}, nil
	}

	// 收集所有用户ID
	userIds := make([]interface{}, 0)
	for _, user := range users {
		userIds = append(userIds, user.Id)
	}

	// 批量查询余额信息
	var balanceUsers []struct {
		UserId        int     `json:"user_id"`
		Balance       float64 `json:"balance"`
		BalanceFrozen float64 `json:"balance_frozen"`
		Points        int     `json:"points"`
	}
	daobalance.BalanceUser.Ctx(ctx).
		WhereIn("user_id", userIds).
		Fields("user_id, balance, balance_frozen, points").
		Scan(&balanceUsers)

	balanceMap := make(map[int]map[string]interface{})
	for _, bu := range balanceUsers {
		balanceMap[bu.UserId] = map[string]interface{}{
			"balance":        bu.Balance,
			"balance_frozen": bu.BalanceFrozen,
			"points":         bu.Points,
		}
	}

	// 批量查询银行卡信息
	var userBanks []struct {
		UserId int    `json:"user_id"`
		CardNo string `json:"card_no"`
	}
	daosite.UserBank.Ctx(ctx).
		WhereIn("user_id", userIds).
		Fields("user_id, card_no").
		Scan(&userBanks)

	bankMap := make(map[int]string)
	for _, ub := range userBanks {
		bankMap[ub.UserId] = ub.CardNo
	}

	// 获取所有代理用户名（批量查询）
	agentIds := make([]interface{}, 0)
	for _, user := range users {
		if user.AgentId > 0 {
			agentIds = append(agentIds, user.AgentId)
		}
	}

	// 查询代理信息
	agentMap := make(map[int]string)
	if len(agentIds) > 0 {
		var agents []entitysite.Agent
		err = daosite.Agent.Ctx(ctx).
			Where("site_id", siteId).
			WhereIn("id", agentIds).
			Fields("id, username").
			Scan(&agents)

		if err == nil {
			for _, agent := range agents {
				agentMap[int(agent.Id)] = agent.Username
			}
		}
	}

	// 获取会员等级和层级列表
	gradeMap, _ := s.getUserGradeMap(ctx, siteId)
	levelMap, _ := s.getUserLevelMap(ctx, siteId)

	// 组装返回数据
	data := make([]g.Map, 0)
	for _, user := range users {
		// 获取余额信息
		balance := 0.0
		balanceFrozen := 0.0
		points := 0
		if balanceInfo, ok := balanceMap[int(user.Id)]; ok {
			balance = balanceInfo["balance"].(float64)
			balanceFrozen = balanceInfo["balance_frozen"].(float64)
			points = balanceInfo["points"].(int)
		}

		// 获取银行卡号
		cardNo := ""
		if cn, ok := bankMap[int(user.Id)]; ok {
			cardNo = cn
		}

		item := g.Map{
			"id":                 user.Id,
			"pay_times":          user.PayTimes,
			"username":           user.Username,
			"grade_id":           user.GradeId,
			"level_id":           user.LevelId,
			"agent_id":           user.AgentId,
			"status":             user.Status,
			"register_ip":        user.RegisterIp,
			"register_time":      user.RegisterTime,
			"last_login_ip":      user.LastLoginIp,
			"last_login_time":    user.LastLoginTime,
			"last_login_address": user.LastLoginAddress,
			"realname":           user.Realname,
			"mobile":             user.Mobile,
			"email":              user.Email,
			"focus_level":        user.FocusLevel,
			"balance_status":     user.BalanceStatus,
			"balance":            balance,
			"balance_frozen":     balanceFrozen,
			"points":             points,
			"card_no":            cardNo,
			"register_url":       user.RegisterUrl,
			"created_at":         user.CreatedAt,
			"is_online":          0, // TODO: 实现在线状态检测
		}

		// 添加代理用户名
		if user.AgentId > 0 {
			item["agent_username"] = agentMap[user.AgentId]
		} else {
			item["agent_username"] = ""
		}

		// 添加等级名称
		if user.GradeId > 0 {
			item["grade_name"] = gradeMap[user.GradeId]
		} else {
			item["grade_name"] = ""
		}

		// 添加层级名称
		if user.LevelId > 0 {
			item["level_name"] = levelMap[user.LevelId]
		} else {
			item["level_name"] = ""
		}

		// 手机号脱敏
		if len(user.Mobile) >= 11 {
			item["mobile"] = user.Mobile[:3] + "****" + user.Mobile[7:]
		}

		data = append(data, item)
	}

	// 获取总会员数
	totalUsers, _ := daosite.User.Ctx(ctx).Where("site_id", siteId).Count()

	// 获取充值会员数（pay_times >= 1）
	totalChargeUsers, _ := daosite.User.Ctx(ctx).
		Where("site_id", siteId).
		Where("pay_times >= ?", 1).
		Count()

	return g.Map{
		"list":               data,
		"count":              count,
		"total_users":        totalUsers,
		"total_charge_users": totalChargeUsers,
	}, nil
}

// getOrderByField 获取排序字段
func (s *sUser) getOrderByField(sortField string) string {
	switch sortField {
	case "register_time":
		return "register_time"
	case "balance":
		return "created_at" // 余额排序暂时用创建时间代替
	case "balance_frozen":
		return "created_at" // 冻结余额排序暂时用创建时间代替
	case "last_login_time":
		return "last_login_time"
	case "points":
		return "created_at" // 积分排序暂时用创建时间代替
	case "first":
		return "created_at"
	default:
		return "created_at"
	}
}

// getUserGradeMap 获取会员等级映射
func (s *sUser) getUserGradeMap(ctx context.Context, siteId uint) (map[int]string, error) {
	var grades []entitysite.UserGrade
	err := daosite.UserGrade.Ctx(ctx).
		Where("site_id", siteId).
		Fields("id, name").
		Scan(&grades)

	if err != nil {
		return nil, err
	}

	gradeMap := make(map[int]string)
	for _, grade := range grades {
		gradeMap[int(grade.Id)] = grade.Name
	}

	return gradeMap, nil
}

// getUserLevelMap 获取会员层级映射
func (s *sUser) getUserLevelMap(ctx context.Context, siteId uint) (map[int]string, error) {
	var levels []entitysite.UserLevel
	err := daosite.UserLevel.Ctx(ctx).
		Where("site_id", siteId).
		Fields("id, name").
		Scan(&levels)

	if err != nil {
		return nil, err
	}

	levelMap := make(map[int]string)
	for _, level := range levels {
		levelMap[int(level.Id)] = level.Name
	}

	return levelMap, nil
}
