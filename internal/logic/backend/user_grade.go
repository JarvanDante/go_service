package backend

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"go-service/api/backendRoute"
	daosite "go-service/internal/dao/jh_site/dao"
	entitysite "go-service/internal/dao/jh_site/model/entity"
	"go-service/internal/dao/jinhuang/model/entity"
	"go-service/internal/service/backend"
	"strconv"
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

// LSaveUserGrades
/**
 * @desc：保存会员等级
 * @param ctx
 * @param req
 * @return err
 * @author : Carson
 */
func (s *sUserGrade) LSaveUserGrades(ctx context.Context, req *backendRoute.SaveUserGradesReq) error {
	site := g.RequestFromCtx(ctx).GetCtxVar("site").Val().(*entity.Site)
	siteId := site.Id

	// 解析JSON数据
	var gradeData []map[string]interface{}
	err := json.Unmarshal([]byte(req.Data), &gradeData)
	if err != nil {
		return errors.New("数据格式错误")
	}

	if len(gradeData) == 0 {
		return errors.New("数据不能为空")
	}

	// 验证数据
	for _, item := range gradeData {
		if err := s.validateGradeData(item); err != nil {
			return err
		}
	}

	// 开启事务
	err = daosite.UserGrade.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		now := gtime.Now()

		for _, item := range gradeData {
			gradeId := getIntValue(item, "id")

			// 准备数据
			data := g.Map{
				"site_id":                siteId,
				"name":                   getStringValue(item, "name"),
				"points_upgrade":         getIntValue(item, "points_upgrade"),
				"bonus_upgrade":          getFloatValue(item, "bonus_upgrade"),
				"bonus_birthday":         getFloatValue(item, "bonus_birthday"),
				"rebate_percent_sports":  percentToDecimal(getFloatValue(item, "rebate_percent_sports")),
				"rebate_percent_lottery": percentToDecimal(getFloatValue(item, "rebate_percent_lottery")),
				"rebate_percent_live":    percentToDecimal(getFloatValue(item, "rebate_percent_live")),
				"rebate_percent_egame":   percentToDecimal(getFloatValue(item, "rebate_percent_egame")),
				"fields_disable":         req.FieldsDisable,
				"auto_providing":         req.AutoProviding,
				"updated_at":             now,
			}

			if gradeId > 0 {
				// 更新
				_, err := tx.Model("user_grade").
					Where("id", gradeId).
					Where("site_id", siteId).
					Data(data).
					Update()
				if err != nil {
					return err
				}
			} else {
				// 新增
				data["created_at"] = now
				data["status"] = 1
				_, err := tx.Model("user_grade").Data(data).Insert()
				if err != nil {
					return err
				}
			}
		}

		return nil
	})

	return err
}

// LDeleteUserGrade
/**
 * @desc：删除会员等级
 * @param ctx
 * @param req
 * @return err
 * @author : Carson
 */
func (s *sUserGrade) LDeleteUserGrade(ctx context.Context, req *backendRoute.DeleteUserGradeReq) error {
	site := g.RequestFromCtx(ctx).GetCtxVar("site").Val().(*entity.Site)
	siteId := site.Id

	// 查询等级信息
	var grade entitysite.UserGrade
	err := daosite.UserGrade.Ctx(ctx).
		Where("id", req.Id).
		Where("site_id", siteId).
		Scan(&grade)

	if err != nil || grade.Id == 0 {
		return errors.New("没有找到该会员等级")
	}

	// 检查是否有会员使用该等级
	userCount, err := daosite.User.Ctx(ctx).
		Where("site_id", siteId).
		Where("grade_id", req.Id).
		Count()

	if err != nil {
		return err
	}

	if userCount > 0 {
		return errors.New("该等级下面有会员，请先将会员变更为其它等级，才能删除该等级")
	}

	// 删除等级
	_, err = daosite.UserGrade.Ctx(ctx).
		Where("id", req.Id).
		Where("site_id", siteId).
		Delete()

	if err != nil {
		return errors.New("删除失败")
	}

	return nil
}

// validateGradeData 验证等级数据
func (s *sUserGrade) validateGradeData(data map[string]interface{}) error {
	name := getStringValue(data, "name")
	if name == "" {
		return errors.New("名称不能为空")
	}

	pointsUpgrade := getIntValue(data, "points_upgrade")
	if pointsUpgrade < 0 {
		return errors.New("升级消耗积分必须是正整数")
	}

	bonusUpgrade := getFloatValue(data, "bonus_upgrade")
	if bonusUpgrade < 0 {
		return errors.New("升级赠送彩金必须是有效金额")
	}

	bonusBirthday := getFloatValue(data, "bonus_birthday")
	if bonusBirthday < 0 {
		return errors.New("生日彩金必须是有效金额")
	}

	// 验证返水百分比
	rebateFields := []string{"rebate_percent_sports", "rebate_percent_lottery", "rebate_percent_live", "rebate_percent_egame"}
	for _, field := range rebateFields {
		value := getFloatValue(data, field)
		if value < 0 || value > 99 {
			return fmt.Errorf("%s百分比必须在0-99之间", field)
		}
	}

	return nil
}

// formatMoney 格式化金额为两位小数
func formatMoney(value float64) string {
	return g.NewVar(value).String()
}

// decimalToPercent 将小数转换为百分比字符串
func decimalToPercent(value float64) string {
	return g.NewVar(value * 100).String()
}

// percentToDecimal 将百分比转换为小数
func percentToDecimal(value float64) float64 {
	return value / 100
}

// getStringValue 从map中获取字符串值
func getStringValue(data map[string]interface{}, key string) string {
	if val, ok := data[key]; ok {
		return g.NewVar(val).String()
	}
	return ""
}

// getIntValue 从map中获取整数值
func getIntValue(data map[string]interface{}, key string) int {
	if val, ok := data[key]; ok {
		return g.NewVar(val).Int()
	}
	return 0
}

// getFloatValue 从map中获取浮点数值
func getFloatValue(data map[string]interface{}, key string) float64 {
	if val, ok := data[key]; ok {
		// 如果是字符串，尝试转换
		if str, ok := val.(string); ok {
			if f, err := strconv.ParseFloat(str, 64); err == nil {
				return f
			}
		}
		return g.NewVar(val).Float64()
	}
	return 0
}
