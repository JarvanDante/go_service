package backend

import (
	"context"
	"go-service/api/backendRoute"
	"go-service/internal/model/response"
	"go-service/internal/service/backend"
)

type RoleController struct{}

func NewRoleController() *RoleController {
	return &RoleController{}
}

//Index
/**
 * showdoc
 * @catalog 后台/员工
 * @title 职务列表
 * @description 职务列表
 * @method get
 * @url /app/roles
 * @param token 必选 string 员工token
 * @return {"code":200,"status":1,"message":"success","data":[{"id":1,"site_id":0,"name":"管理员","status":0,"created_at":null,"updated_at":null,"permissions":""},{"id":2,"site_id":0,"name":"客服","status":0,"created_at":null,"updated_at":null,"permissions":""},{"id":3,"site_id":0,"name":"财务","status":0,"created_at":null,"updated_at":null,"permissions":""},{"id":4,"site_id":0,"name":"运营","status":0,"created_at":null,"updated_at":null,"permissions":""}]}
 * @return_param code int 状态码
 * @return_param status int 成功/失败状态
 * @return_param message string 提示说明
 * @return_param data string 主要数据
 * @return_param data.id string 角色id
 * @return_param data.name string 角色名称
 * @remark 备注
 * @number 1
 */
func (c *RoleController) Index(ctx context.Context, req *backendRoute.RolesReq) (res *response.Response, err error) {
	data, err := backend.ServiceRole().LIndex(ctx, req)
	if err != nil {
		return nil, err
	}

	response.JsonOkCtx(ctx, data)

	return
}

func (c *RoleController) Permissions(ctx context.Context, req *backendRoute.PermissionsReq) (res *response.Response, err error) {
	data, err := backend.ServiceRole().LPermissions(ctx, req)
	if err != nil {
		return nil, err
	}

	response.JsonOkCtx(ctx, data)

	return
}
