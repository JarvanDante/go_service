package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "go-service/internal/logic/frontend"
	_ "go-service/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"go-service/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
