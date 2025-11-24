package backend

import (
	"context"
	"go-service/api/backendRoute"
	"go-service/internal/model/response"
	"go-service/internal/service/backend"
)

type LogController struct{}

func NewLogController() *LogController {
	return &LogController{}
}

//AdminLogs
/**
 * showdoc
 * @catalog 后台/系统
 * @title 后台操作日志
 * @description 后台操作日志列表
 * @method get
 * @url /app/admin-logs
 * @param token 必选 string 员工token
 * @param username 可选 string 用户名
 * @param start 可选 string 开始时间
 * @param end 可选 string 结束时间
 * @param page 可选 int 页数
 * @param size 可选 int 页码
 * @return {"code":200,"status":1,"message":"获取数据成功","data":{"count":100,"list":[{"username":"admin","ip":"127.0.0.1","remark":"登录系统","created_at":"2024-01-01 12:00:00"}]}}
 * @return_param code int 状态码
 * @return_param status int 成功/失败状态
 * @return_param message string 提示说明
 * @return_param data object 数据
 * @return_param data.count int 总条数
 * @return_param data.list array 日志列表
 * @return_param data.list.username string 用户名
 * @return_param data.list.ip string IP地址
 * @return_param data.list.remark string 操作说明
 * @return_param data.list.created_at string 创建时间
 * @remark 备注
 * @number 1
 */
func (c *LogController) AdminLogs(ctx context.Context, req *backendRoute.AdminLogsReq) (res *response.Response, err error) {

	data, err := backend.ServiceLog().LAdminLogs(ctx, req)
	if err != nil {
		response.JsonErrCtx(ctx, err.Error())
		return
	}

	response.JsonOkCtx(ctx, data, "获取数据成功")

	return
}
