package backend

import (
	"context"
	"go-service/api/backendRoute"
)

type (
	IRebateRule interface {
		LIndex(ctx context.Context, req *backendRoute.RebateSettingReq) (res interface{}, err error)
		LCreateRebateRule(ctx context.Context, req *backendRoute.CreateRebateRuleReq) (err error)
		LUpdateRebateRule(ctx context.Context, req *backendRoute.UpdateRebateRuleReq) (err error)
		LDeleteRebateRule(ctx context.Context, req *backendRoute.DeleteRebateRuleReq) (err error)
		LCreateRebateOption(ctx context.Context, req *backendRoute.CreateRebateOptionReq) (err error)
		LUpdateRebateOption(ctx context.Context, req *backendRoute.UpdateRebateOptionReq) (err error)
		LDeleteRebateOption(ctx context.Context, req *backendRoute.DeleteRebateOptionReq) (err error)
	}
)

var localRebateRule IRebateRule

func ServiceRebateRule() IRebateRule {
	return localRebateRule
}

func RegisterRebateRule(p IRebateRule) {
	localRebateRule = p
}
