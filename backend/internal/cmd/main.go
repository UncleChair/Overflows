package cmd

import (
	"context"
	"errors"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
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
			port := parser.GetOpt("p", 8000)
			s := g.Server()
			s.SetPort(port.Int())
			mode := parser.GetOpt("m", "standalone")
			if mode.String() == "standalone" {
				gdb.SetDefaultGroup("standaloneMode")
				Standalone.Run(ctx)
			} else if mode.String() == "server" {
				gdb.SetDefaultGroup("serverMode")
				Server.Run(ctx)
			} else {
				return errors.New("invalid mode")
			}
			return nil
		},
	}
)
