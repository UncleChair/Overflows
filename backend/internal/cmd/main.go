package cmd

import (
	"context"
	"errors"
	"overflows/internal/service"

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
			mode := parser.GetOpt("m", "standalone")
			port := parser.GetOpt("p", 8000)
			s := g.Server()
			s.SetPort(port.Int())

			if mode.String() == "standalone" || mode.String() == "server" {
				gdb.SetDefaultGroup(mode.String() + "Mode")
			} else {
				return errors.New("invalid mode")
			}

			// Do migration for standalone mode only
			if mode.String() == "standalone" {
				err = service.SQLiteMigration(ctx)
				if err != nil {
					g.Log().Error(ctx, "Database migration failed:", err)
					return err
				}
			}

			// Init Casbin
			if mode.String() == "standalone" || mode.String() == "server" {
				service.InitCasbin(ctx, mode.String()+"Mode")
			}

			if mode.String() == "standalone" {
				Standalone.Run(ctx)
			} else if mode.String() == "server" {
				Server.Run(ctx)
			}
			return nil
		},
	}
)
