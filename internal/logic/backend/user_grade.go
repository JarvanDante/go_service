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

type sUserGrade struct {
}

func init() {
	backend.RegisterUserGrade(&sUserGrade{})
}

// LIndex
/**
 * @desc：获取会员等级列表
 * @param ctx
 * @param req
 * @return res
 * @return err
 * @author : Carson
 */
func (s *sUserGrade) LIndex(ctx context.Context, req *backendRoute.UserGradesReq) (interface{}, error) {
	site := g.RequestFromCtx(ctx).GetCtxVar("site").Val().(*entity.Site)
	siteId := site.Id

	// 查询会员等级列表
	var grades []entitysite.UserGrade
	err := daosite.UserGrade.Ctx(ctx).
		Where("site_id", siteId).
		Where("status", 1). // STATUS_ACTIVE = 1
		Order("id ASC").
		Scan(&grades)

	if err != nil {
		return nil, err
	}

	// 组装返回数据
	data := make([]g.Map, 0)
	for i, grade := range grades {
		item := g.Map{
			"id":                     grade.Id,
			"name":                   grade.Name,
			"points_upgrade":         grade.PointsUpgrade,
			"bonus_upgrade":          formatMoney(grade.BonusUpgrade),
			"bonus_birthday":         formatMoney(grade.BonusBirthday),
			"rebate_percent_sports":  decimalToPercent(grade.RebatePercentSports),
			"rebate_percent_lottery": decimalToPercent(grade.RebatePercentLottery),
			"rebate_percent_live":    decimalToPercent(grade.RebatePercentLive),
			"rebate_percent_egame":   decimalToPercent(grade.RebatePercentEgame),
		}

		// 获取该等级的会员数量
		userCount, _ := daosite.User.Ctx(ctx).
			Where("site_id", siteId).
			Where("grade_id", grade.Id).
			Count()
		item["user_count"] = userCount

		// 只在第一个等级返回字段配置
		if i == 0 {
			item["fields_disable"] = grade.FieldsDisable
			item["auto_providing"] = grade.AutoProviding
		}

		data = append(data, item)
	}

	return data, nil
}

// formatMoney 格式化金额为两位小数
func formatMoney(value float64) string {
	return g.NewVar(value).String()
}

// decimalToPercent 将小数转换为百分比字符串
func decimalToPercent(value float64) string {
	return g.NewVar(value * 100).String()
}
