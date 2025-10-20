package main

import (
	_ "go-service/internal/logic"
	_ "go-service/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"go-service/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
