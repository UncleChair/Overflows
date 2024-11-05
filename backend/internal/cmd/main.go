package cmd

import (
	"context"

	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gres"
)

var (
	Main = gcmd.Command{
		Name:  "overflows",
		Brief: "start Overflows",
		Arguments: []gcmd.Argument{
			{Name: "port", Short: "p", Brief: "port number", IsArg: false},
			{Name: "mode", Short: "m", Brief: "run mode, choose from [standalone, server]", IsArg: false},
		},
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			gres.Dump()
			Init.Run(ctx)
			mode := parser.GetOpt("m", "standalone")
			if mode.String() == "standalone" {
				Standalone.Run(ctx)
			} else if mode.String() == "server" {
				Server.Run(ctx)
			}
			return nil
		},
	}
)
