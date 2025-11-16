package response

// PermissionTree 权限树结构
type PermissionTree struct {
	ID          int               `json:"id"`
	Type        int               `json:"type"`
	Name        string            `json:"name"`
	BackendURL  string            `json:"backend_url"`
	FrontendURL string            `json:"frontend_url"`
	Open        bool              `json:"open"`
	Checked     bool              `json:"checked"`
	Children    []*PermissionTree `json:"children"`
}

// RoleWithPermissions 职务及权限列表
type RoleWithPermissions struct {
	ID             int      `json:"id"`
	Name           string   `json:"name"`
	PermissionList []string `json:"permission_list"`
}

// RolePermissionsRes 返回数据结构
type RolePermissionsRes struct {
	PermissionList []*PermissionTree      `json:"permission_list"`
	RoleList       []*RoleWithPermissions `json:"role_list"`
}
