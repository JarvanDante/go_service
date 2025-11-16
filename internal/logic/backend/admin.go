package backend

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"go-service/api/backendRoute"
	daosite "go-service/internal/dao/jh_site/dao"
	"go-service/internal/dao/jh_site/model/entity"
	entityjh "go-service/internal/dao/jinhuang/model/entity"
	"go-service/internal/model/response"
	"go-service/internal/service/backend"
	"go-service/utility/helpers"
)

type sAdmin struct {
}

func init() {
	backend.RegisterAdmin(&sAdmin{})
}

//LGetInfo
/**
 * @desc：获取管理员信息
 * @param ctx
 * @param req
 * @return res
 * @return err
 * @author : Carson
 */
func (s *sAdmin) LGetInfo(ctx context.Context, req *backendRoute.GetInfoReq) (res *response.GetInfoRes, err error) {
	//获取当前admin信息
	admin := g.RequestFromCtx(ctx).GetCtxVar("admin").Val().(*entity.Admin)

	helpers.Print(admin.Id)

	// 获取角色
	role, err := daosite.GetRole(ctx, admin.AdminRoleId)
	if err != nil {
		return nil, err
	}
	permIds := daosite.GetPermissionIds(role)
	permList, _ := daosite.GetPermissions(ctx, permIds)

	menuTree := BuildTree(permList, 0)

	return &response.GetInfoRes{
		Roles:        []string{role.Name},
		Name:         admin.Username,
		Avatar:       "https://wpimg.wallstcn.com/577965b9-bb9e-4e02-9f0c-095b41417191",
		Introduction: "",
		Menus:        menuTree,
	}, nil

}

//BuildTree
/**
 * @desc：构造菜单树
 * @param list
 * @param parentId
 * @return []map[string]interface{}
 * @author : Carson
 */
func BuildTree(list []*entityjh.AdminPermission, parentId int) []map[string]interface{} {
	var tree []map[string]interface{}

	for _, item := range list {
		if item.ParentId == parentId && item.Type == 1 { // 只构建菜单
			node := map[string]interface{}{
				"id":       item.Id,
				"name":     item.Name,
				"path":     item.FrontendUrl,
				"type":     item.Type,
				"sort":     item.Sort,
				"children": BuildTree(list, gconv.Int(item.Id)),
			}
			tree = append(tree, node)
		}
	}

	return tree
}
