package backend

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"go-service/api/backendRoute"
	daosite "go-service/internal/dao/jh_site/dao"
	"go-service/internal/dao/jh_site/model/do"
	entitysite "go-service/internal/dao/jh_site/model/entity"
	"go-service/internal/dao/jinhuang/model/entity"
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
	admin := g.RequestFromCtx(ctx).GetCtxVar("admin").Val().(*entitysite.Admin)

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
func BuildTree(list []*entity.AdminPermission, parentId int) []map[string]interface{} {
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

func (s *sAdmin) LAdmins(ctx context.Context, req *backendRoute.AdminsReq) (interface{}, error) {
	site := g.RequestFromCtx(ctx).GetCtxVar("site").Val().(*entity.Site)
	siteId := site.Id

	where := map[string]interface{}{}
	where["site_id"] = siteId
	where["delete_at"] = 0

	if req.Username != nil && *req.Username != "" {
		where["username like ?"] = "%" + *req.Username + "%"

	}
	if req.Status != nil {
		where["status"] = *req.Status
	}
	m := daosite.Admin.Ctx(ctx).Where(where)

	total, _ := m.Count()

	page := *req.Page
	size := *req.Size

	var list []*entitysite.Admin
	err := m.OrderDesc("id").Page(page, size).Scan(&list)
	if err != nil {
		return nil, err
	}

	// 获取角色名
	var roles []*entitysite.AdminRole
	err = daosite.AdminRole.Ctx(ctx).
		Where("site_id", siteId).
		Scan(&roles)
	if err != nil {
		return nil, err
	}

	roleMap := map[int]string{}
	for _, r := range roles {

		roleMap[gconv.Int(r.Id)] = r.Name
	}

	var result []map[string]interface{}
	for _, a := range list {
		result = append(result, g.Map{
			"id":              a.Id,
			"username":        a.Username,
			"admin_role_id":   a.AdminRoleId,
			"admin_role_name": roleMap[a.AdminRoleId],
			"nickname":        a.Nickname,
			"status":          a.Status,
			"last_login_ip":   a.LastLoginIp,
			"last_login_time": a.LastLoginTime,
		})
	}

	return g.Map{
		"list":  result,
		"count": total,
	}, nil
}

func (s *sAdmin) LCreateAdmin(ctx context.Context, req *backendRoute.CreateAdminReq) (err error) {
	site := g.RequestFromCtx(ctx).GetCtxVar("site").Val().(*entity.Site)

	// 判断是否已存在
	exist, err := daosite.Admin.
		Ctx(ctx).
		Where(do.Admin{
			SiteId:   site.Id,
			Username: req.Username,
			DeleteAt: 0,
		}).One()

	if err != nil {
		return err
	}
	if !exist.IsEmpty() {
		return gerror.New("用户名已经被使用")
	}

	bcrypt, err := helpers.Bcrypt(req.Password)
	if err != nil {
		return err
	}
	// 新增
	_, err = daosite.Admin.Ctx(ctx).Data(do.Admin{
		SiteId:      site.Id,
		Username:    req.Username,
		Nickname:    req.Nickname,
		Password:    bcrypt,
		AdminRoleId: req.Role,
		Status:      req.Status,
		DeleteAt:    0,
		CreatedAt:   gtime.Now(),
	}).Insert()

	if err != nil {
		return err
	}

	// 日志
	err = daosite.AddAdminLog(ctx, "添加员工："+req.Username)
	if err != nil {
		return err
	}

	return nil
}
