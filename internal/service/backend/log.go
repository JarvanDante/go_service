package backend

import (
	"context"
	"go-service/api/backendRoute"
)

type (
	ILog interface {
		LAdminLogs(ctx context.Context, req *backendRoute.AdminLogsReq) (res interface{}, err error)
	}
)

var localLog ILog

func ServiceLog() ILog {
	return localLog
}

func RegisterLog(p ILog) {
	localLog = p
}
