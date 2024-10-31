package main

import (
	"context"
	"fmt"
	"os/exec"
	_ "overflows/internal/packed"
	"runtime"

	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/ncruces/zenity"

	"overflows/internal/cmd"

	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	_ "github.com/gogf/gf/contrib/drivers/sqlite/v2"
)

func main() {
	ctx := gctx.New()
	systray.Run(func() {
		onReady(ctx)
	}, func() {
		onExit(ctx)
	})
}

func onReady(ctx context.Context) {
	systray.SetTemplateIcon(icon.Data, icon.Data)
	systray.SetTitle("Overflows")
	systray.SetTooltip("Overflows")

	go func() {
		s := g.Server()
		mUrl := systray.AddMenuItem("Open UI", "Open Overflows Frontend")
		mQuit := systray.AddMenuItem("Exit", "Quit the whole app")

		// Sets the icon of a menu item. Only available on Mac.
		// mQuit.SetIcon(icon.Data)

		systray.AddSeparator()

		for {
			select {
			case <-mUrl.ClickedCh:
				port := s.GetListenedPort()
				openURL(ctx, fmt.Sprintf("http://127.0.0.1:%d", port))
			case <-mQuit.ClickedCh:
				systray.Quit()
				s.Shutdown()
				fmt.Println("Quit Overflows...")
				return
			}
		}
	}()

	go func() {
		cmd.Main.Run(ctx)
	}()
}

func onExit(ctx context.Context) {
	g.Log().Info(ctx, "ShutDown App")
}

func openURL(ctx context.Context, url string) {
	var err error
	switch runtime.GOOS {
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	}
	g.Log().Info(ctx, "openURL Failed: ", err.Error())
	if err != nil {
		zenity.Info(fmt.Sprintf("Please input the following URL in your browser:\n%s", url),
			zenity.Title("Failed to open UI"),
			zenity.InfoIcon)
	}
}
