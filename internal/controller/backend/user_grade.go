package backend

import (
	"context"
	"go-service/api/backendRoute"
	"go-service/internal/model/response"
	"go-service/internal/service/backend"
)

type UserGradeController struct{}

func NewUserGradeController() *UserGradeController {
	return &UserGradeController{}
}

// Index
/**
 * showdoc
 * @catalog 后台/会员
 * @title 获取会员等级列表
 * @description 获取会员等级列表
 * @method get
 * @url /app/user-grades
 * @param token 必选 string 员工token
 * @return {"code":200,"status":1,"message":"获取数据成功","data":[]}
 * @return_param code int 状态码
 * @return_param status int 成功/失败状态
 * @return_param message string 提示说明
 * @return_param data array 会员等级列表
 * @remark 备注
 * @number 1
 */
func (c *UserGradeController) Index(ctx context.Context, req *backendRoute.UserGradesReq) (res *response.Response, err error) {
	data, err := backend.ServiceUserGrade().LIndex(ctx, req)
	if err != nil {
		response.JsonErrCtx(ctx, err.Error())
		return
	}

	response.JsonOkCtx(ctx, data, "获取数据成功")
	return
}

// SaveUserGrades
/**
 * showdoc
 * @catalog 后台/会员
 * @title 保存会员等级
 * @description 保存会员等级
 * @method post
 * @url /app/save-user-grades
 * @param token 必选 string 员工token
 * @param data 必选 string 会员等级数据(JSON字符串)
 * @param fieldsDisable 可选 string 不显示的字段
 * @param autoProviding 可选 string 自动发放的业务
 * @return {"code":200,"status":1,"message":"保存成功","data":null}
 * @return_param code int 状态码
 * @return_param status int 成功/失败状态
 * @return_param message string 提示说明
 * @remark 备注
 * @number 2
 */
func (c *UserGradeController) SaveUserGrades(ctx context.Context, req *backendRoute.SaveUserGradesReq) (res *response.Response, err error) {
	err = backend.ServiceUserGrade().LSaveUserGrades(ctx, req)
	if err != nil {
		response.JsonErrCtx(ctx, err.Error())
		return
	}

	response.JsonOkCtx(ctx, nil, "保存成功")
	return
}

// DeleteUserGrade
/**
 * showdoc
 * @catalog 后台/会员
 * @title 删除会员等级
 * @description 删除会员等级
 * @method post
 * @url /app/delete-user-grades
 * @param token 必选 string 员工token
 * @param id 必选 int 会员等级ID
 * @return {"code":200,"status":1,"message":"删除成功","data":null}
 * @return_param code int 状态码
 * @return_param status int 成功/失败状态
 * @return_param message string 提示说明
 * @remark 备注
 * @number 3
 */
func (c *UserGradeController) DeleteUserGrade(ctx context.Context, req *backendRoute.DeleteUserGradeReq) (res *response.Response, err error) {
	err = backend.ServiceUserGrade().LDeleteUserGrade(ctx, req)
	if err != nil {
		response.JsonErrCtx(ctx, err.Error())
		return
	}

	response.JsonOkCtx(ctx, nil, "删除成功")
	return
}
