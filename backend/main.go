package main

import (
	_ "overflows/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"overflows/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
