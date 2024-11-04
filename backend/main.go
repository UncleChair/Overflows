package main

import (
	_ "overflows/internal/packed"

	_ "overflows/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"overflows/internal/cmd"

	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	_ "github.com/gogf/gf/contrib/drivers/sqlite/v2"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
