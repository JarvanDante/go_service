package backend

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"go-service/api/backendRoute"
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

	// 获取总数（在添加银行卡筛选之前）
	count, err := model.Count()
	if err != nil {
		return nil, err
	}

	// 银行卡号筛选（需要 JOIN user_bank 表）
	if req.CardNo != nil && *req.CardNo != "" {
		model = model.LeftJoin("user_bank", "user.id = user_bank.user_id")
		model = model.Where("user_bank.card_no", *req.CardNo)
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
	orderField := "user.created_at"
	if req.SortField != nil && *req.SortField != "" {
		orderField = s.getOrderByField(*req.SortField)
	}

	// 排序规则
	orderRule := "DESC"
	if req.SortRule != nil && *req.SortRule == 0 {
		orderRule = "ASC"
	}

	// 查询字段
	fields := "user.id, user.pay_times, user.username, user.grade_id, user.level_id, user.agent_id, " +
		"user.status, user.register_ip, user.register_time, user.last_login_ip, user.last_login_time, " +
		"user.last_login_address, user.realname, user.mobile, user.email, user.focus_level, " +
		"user.balance_status, user.register_url"

	// 关联 balance_user 表
	model = model.LeftJoin("jh_balance.balance_user", "user.id = balance_user.user_id")
	fields += ", balance_user.balance, balance_user.balance_frozen, balance_user.points"

	// 如果没有银行卡筛选，也需要关联 user_bank 表来获取银行卡号
	if req.CardNo == nil || *req.CardNo == "" {
		model = model.LeftJoin("user_bank", "user.id = user_bank.user_id")
	}
	fields += ", user_bank.card_no"

	// 执行查询
	var userList []map[string]interface{}
	err = model.Fields(fields).
		Page(page, size).
		Order(orderField + " " + orderRule).
		Scan(&userList)

	if err != nil {
		return nil, err
	}

	// 获取所有代理用户名（批量查询）
	agentIds := make([]int, 0)
	for _, user := range userList {
		if agentId, ok := user["agent_id"].(int); ok && agentId > 0 {
			agentIds = append(agentIds, agentId)
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
	for _, user := range userList {
		item := g.Map{
			"id":                 user["id"],
			"pay_times":          user["pay_times"],
			"username":           user["username"],
			"grade_id":           user["grade_id"],
			"level_id":           user["level_id"],
			"agent_id":           user["agent_id"],
			"status":             user["status"],
			"register_ip":        user["register_ip"],
			"register_time":      user["register_time"],
			"last_login_ip":      user["last_login_ip"],
			"last_login_time":    user["last_login_time"],
			"last_login_address": user["last_login_address"],
			"realname":           user["realname"],
			"mobile":             user["mobile"],
			"email":              user["email"],
			"focus_level":        user["focus_level"],
			"balance_status":     user["balance_status"],
			"balance":            user["balance"],
			"balance_frozen":     user["balance_frozen"],
			"points":             user["points"],
			"card_no":            user["card_no"],
			"register_url":       user["register_url"],
			"is_online":          0, // TODO: 实现在线状态检测
		}

		// 添加代理用户名
		if agentId, ok := user["agent_id"].(int); ok {
			item["agent_username"] = agentMap[agentId]
		}

		// 添加等级名称
		if gradeId, ok := user["grade_id"].(int); ok {
			item["grade_name"] = gradeMap[gradeId]
		}

		// 添加层级名称
		if levelId, ok := user["level_id"].(int); ok {
			item["level_name"] = levelMap[levelId]
		}

		// 手机号脱敏
		if mobile, ok := user["mobile"].(string); ok && len(mobile) >= 11 {
			item["mobile"] = mobile[:3] + "****" + mobile[7:]
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
		return "user.register_time"
	case "balance":
		return "balance_user.balance"
	case "balance_frozen":
		return "balance_user.balance_frozen"
	case "last_login_time":
		return "user.last_login_time"
	case "points":
		return "balance_user.points"
	case "first":
		return "user.created_at"
	default:
		return "user.created_at"
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
