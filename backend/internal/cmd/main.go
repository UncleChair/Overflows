package cmd

import (
	"context"
	"errors"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
)

type MainInput struct {
	Port       int  `name:"port" short:"p" default:"8000" description:"port number" arg:"true"`
	Standalone bool `name:"standalone" short:"s" default:"true" description:"enable standalone mode" arg:"true"`
}

var (
	Main = gcmd.Command{
		Name:  "overflows",
		Brief: "start Overflows",
		Arguments: []gcmd.Argument{
			{Name: "port", Short: "p", Brief: "port number", IsArg: false, Orphan: true},
			{Name: "mode", Short: "m", Brief: "mode", IsArg: false, Orphan: true},
		},
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			port := parser.GetOpt("p", 8000)
			s := g.Server()
			s.SetPort(port.Int())
			mode := parser.GetOpt("m", "standalone")
			if mode.String() == "standalone" {
				Standalone.Run(ctx)
			} else if mode.String() == "server" {
				Server.Run(ctx)
			} else {
				return errors.New("invalid mode")
			}
			return nil
		},
	}
)
