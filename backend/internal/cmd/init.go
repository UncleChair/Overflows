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
	Init = gcmd.Command{
		Name:  "Init",
		Brief: "Init project settings",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			mode := parser.GetOpt("m", "standalone")

			port := parser.GetOpt("p", 8000)
			s := g.Server()
			s.SetPort(port.Int())

			if mode.String() == "standalone" {
				gdb.SetDefaultGroup("standaloneMode")
			} else if mode.String() == "server" {
				gdb.SetDefaultGroup("serverMode")
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
			if mode.String() == "standalone" {
				service.InitCasbin(ctx, "standaloneMode")
			} else if mode.String() == "server" {
				service.InitCasbin(ctx, "serverMode")
			}

			// Check database connection
			db, err := gdb.Instance()
			if err != nil {
				g.Log().Error(ctx, "Database configuration error:", err)
				return err
			}
			_, err = db.Raw("SELECT 1").All()
			if err != nil {
				g.Log().Error(ctx, "Database connection failed:", err)
				return err
			}
			return nil
		},
	}
)
