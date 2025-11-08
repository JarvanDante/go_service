package backend

import (
	"context"
	"go-service/api/backendRoute"
	daosite "go-service/internal/dao/jh_site/dao"
	"go-service/internal/dao/jh_site/model/entity"
	daojh "go-service/internal/dao/jinhuang/dao"
	"go-service/internal/service/backend"
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
func (s *sRole) LIndex(ctx context.Context, req *backendRoute.RoleReq) (role []*entity.AdminRole, err error) {

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
