package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "go-service/internal/logic/backend"
	_ "go-service/internal/logic/frontend"
	_ "go-service/internal/logic/middleware"
	_ "go-service/internal/packed"
	"go-service/utility/global"

	"github.com/gogf/gf/v2/os/gctx"

	"go-service/internal/cmd"
)

func main() {

	// 初始化全局变量
	if err := global.Init(); err != nil {
		panic(err)
	}

	cmd.Main.Run(gctx.GetInitCtx())
}
