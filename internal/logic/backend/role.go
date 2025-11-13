package backend

import (
	"context"
	"go-service/api/backendRoute"
	daosite "go-service/internal/dao/jh_site/dao"
	"go-service/internal/dao/jh_site/model/entity"
	daojh "go-service/internal/dao/jinhuang/dao"
	entityJh "go-service/internal/dao/jinhuang/model/entity"
	"go-service/internal/model/response"
	"go-service/internal/service/backend"
	"strings"
)

type sRole struct {
}

func init() {
	backend.RegisterRole(&sRole{})
}

//LIndex
/**
 * @desc：角色列表
 * @param ctx
 * @param req
 * @return role
 * @return err
 * @author : Carson
 */
func (s *sRole) LIndex(ctx context.Context, req *backendRoute.RolesReq) (role []*entity.AdminRole, err error) {

	site, err := daojh.GetSiteObject()
	if err != nil {
		return nil, err
	}

	role, err = daosite.GetRoleList(site.Id)
	if err != nil {
		return nil, err
	}

	return

}

//LPermissions
/**
 * @desc：权限列表和角色包含权限
 * @param ctx
 * @param req
 * @return res
 * @return err
 * @author : Carson
 */
func (s *sRole) LPermissions(ctx context.Context, req *backendRoute.PermissionsReq) (res *response.RolePermissionsRes, err error) {
	// 1️⃣ 获取站点
	site, err := daojh.GetSiteObject()
	if err != nil {
		return nil, err
	}

	// 2️⃣ 查询所有权限
	permissionEntities, err := daojh.GetPermissionsBySite(ctx)
	if err != nil {
		return nil, err
	}

	// 3️⃣ 将权限表转为树结构
	permissionTree := buildPermissionTree(permissionEntities, 0)

	// 4️⃣ 查询角色及其权限
	roleEntities, err := daosite.GetRoleList(site.Id)
	if err != nil {
		return nil, err
	}

	var roleList []*response.RoleWithPermissions
	for _, r := range roleEntities {
		roleList = append(roleList, &response.RoleWithPermissions{
			ID:             int(r.Id),
			Name:           r.Name,
			PermissionList: parsePermissionIDs(r.Permissions), // 如: "1,2,3"
		})
	}

	// 5️⃣ 返回
	res = &response.RolePermissionsRes{
		PermissionList: permissionTree,
		RoleList:       roleList,
	}
	return res, nil
}

func (s *sRole) LCreate(ctx context.Context, req *backendRoute.CreateReq) (err error) {

	site, err := daojh.GetSiteObject()
	if err != nil {
		return err
	}

	err = daosite.AddRole(ctx, int(site.Id), req)
	if err != nil {
		return err
	}

	return

}

//-----以下私有方法-------

//buildPermissionTree
/**
 * @desc：构建树形结构
 * @param list
 * @param parentID
 * @return []*response.PermissionTree
 * @author : Carson
 */
func buildPermissionTree(list []*entityJh.AdminPermission, parentID int) []*response.PermissionTree {
	var tree []*response.PermissionTree
	for _, item := range list {
		if item.ParentId == parentID {
			node := &response.PermissionTree{
				ID:          int(item.Id),
				Type:        item.Type,
				Name:        item.Name,
				BackendURL:  item.BackendUrl,
				FrontendURL: item.FrontendUrl,
				Open:        true,
				Checked:     false,
				Children:    buildPermissionTree(list, int(item.Id)),
			}
			tree = append(tree, node)
		}
	}
	return tree
}

//parsePermissionIDs
/**
 * @desc："1,2,3" 转为 []string
 * @param csv
 * @return []string
 * @author : Carson
 */
func parsePermissionIDs(csv string) []string {
	if csv == "" {
		return []string{}
	}
	parts := strings.Split(csv, ",")
	return parts
}
