package response

type GetInfoRes struct {
	Roles        []string    `json:"roles"`
	Name         string      `json:"name"`
	Avatar       string      `json:"avatar"`
	Introduction string      `json:"introduction"`
	Menus        interface{} `json:"menus"`
}
