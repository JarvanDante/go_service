# 菜单接口迁移总结

## 迁移内容

已将 b_service 中的菜单接口 `$api->get('menus', 'AdminController@menu')` 迁移到 go_service 项目中。

## 迁移详情

### 1. 路由定义

**文件**: `api/backendRoute/admin.go`

添加了菜单请求结构体：

```go
type MenusReq struct {
	g.Meta `path:"/menus" method:"get" summary:"获取菜单列表"`
}
```

### 2. 控制器方法

**文件**: `internal/controller/backend/admin.go`

添加了 `Menus` 方法：

```go
func (c *AdminController) Menus(ctx context.Context, req *backendRoute.MenusReq) (res *response.Response, err error)
```

该方法调用服务层的 `LMenus` 方法获取菜单数据。

### 3. 服务接口

**文件**: `internal/service/backend/admin.go`

在 `IAdmin` 接口中添加了方法签名：

```go
LMenus(ctx context.Context, req *backendRoute.MenusReq) (res interface{}, err error)
```

### 4. 服务实现

**文件**: `internal/logic/backend/admin.go`

实现了 `LMenus` 方法和 `BuildMenuTree` 辅助函数：

- `LMenus`: 获取当前登录用户的菜单列表
- `BuildMenuTree`: 构造菜单树结构

## 接口说明

### 请求

```
GET /api/backend/menus
```

**参数**:

- `token` (必选): 员工 token

### 响应

```json
{
  "code": 200,
  "status": 1,
  "message": "获取数据成功",
  "data": [
    {
      "id": 1,
      "type": 1,
      "name": "系统",
      "backend_url": "",
      "frontend_url": "sysSetting",
      "open": true,
      "checked": false,
      "children": [
        {
          "id": 2,
          "type": 1,
          "name": "全局设置",
          "backend_url": "basic-setting",
          "frontend_url": "sysSetting/basicSetting",
          "open": true,
          "checked": false,
          "children": [
            {
              "id": 3,
              "type": 1,
              "name": "基本信息",
              "backend_url": "user-basic-info",
              "frontend_url": "sysSetting/basicSetting/sysBasicSet",
              "open": true,
              "checked": false,
              "children": null
            }
          ]
        }
      ]
    }
  ]
}
```

## 菜单数据结构

每个菜单项包含以下字段：

| 字段         | 类型    | 说明                       |
| ------------ | ------- | -------------------------- |
| id           | int     | 菜单 ID                    |
| type         | int     | 菜单类型（1=页面，2=操作） |
| name         | string  | 菜单名称                   |
| backend_url  | string  | 后端 URL                   |
| frontend_url | string  | 前端路由路径               |
| open         | boolean | 是否展开                   |
| checked      | boolean | 是否选中                   |
| children     | array   | 子菜单列表                 |

## 权限控制

菜单列表根据当前登录用户的角色权限进行过滤：

1. 获取当前用户的角色
2. 获取角色对应的权限列表
3. 只返回类型为 1（页面）的权限项
4. 构造树形结构

## 与 b_service 的对比

### b_service 实现

```php
public function menu()
{
    $data = $this->admin->getRolePermissionList();
    return success('获取数据成功', $data);
}
```

### go_service 实现

```go
func (s *sAdmin) LMenus(ctx context.Context, req *backendRoute.MenusReq) (interface{}, error) {
    admin := g.RequestFromCtx(ctx).GetCtxVar("admin").Val().(*entitysite.Admin)
    role, err := daosite.GetRole(ctx, admin.AdminRoleId)
    if err != nil {
        return nil, err
    }
    permIds := daosite.GetPermissionIds(role)
    permList, _ := daosite.GetPermissions(ctx, permIds)
    menuTree := BuildMenuTree(permList, 0)
    return menuTree, nil
}
```

## 修改的文件

1. `api/backendRoute/admin.go` - 添加 MenusReq 结构体
2. `internal/controller/backend/admin.go` - 添加 Menus 方法
3. `internal/service/backend/admin.go` - 添加 LMenus 方法签名
4. `internal/logic/backend/admin.go` - 实现 LMenus 和 BuildMenuTree 方法

## 测试

可以通过以下方式测试菜单接口：

```bash
curl -X GET "http://localhost:8000/api/backend/menus" \
  -H "Authorization: Bearer {token}"
```

## 注意事项

1. 菜单接口需要有效的 token（通过认证中间件验证）
2. 只返回类型为 1 的菜单项（页面类型）
3. 菜单树根据权限进行过滤
4. 子菜单为空时返回 null 而不是空数组

## 后续工作

1. 前端需要更新菜单 API 调用地址为 `/api/backend/menus`
2. 确保前端菜单转换逻辑与返回的数据结构兼容
3. 测试不同角色的菜单权限显示
