package cmd

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/getlantern/systray"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/ncruces/zenity"
)

var (
	Standalone = gcmd.Command{
		Name: "standalone",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			go func() {
				err = zenity.Info("App is running in background",
					zenity.Title("Overflows"),
					zenity.InfoIcon)
				if err != nil {
					panic(err)
				}
			}()
			systray.Run(func() {
				onReady(ctx)
			}, func() {
				onExit(ctx)
			})
			return nil
		},
	}
)

func onReady(ctx context.Context) {
	iconPath := "resource/static/logo.ico"
	iconData, err := os.ReadFile(iconPath)
	if err != nil {
		panic(err)
	}
	systray.SetIcon(iconData)
	systray.SetTitle("Overflows")
	systray.SetTooltip("Overflows")

	go func() {
		mUrl := systray.AddMenuItem("Open UI", "Open Overflows Frontend")
		mQuit := systray.AddMenuItem("Exit", "Quit the whole app")
		systray.AddSeparator()
		for {
			select {
			case <-mUrl.ClickedCh:
				s := g.Server()
				port := s.GetListenedPort()
				openURL(fmt.Sprintf("http://127.0.0.1:%d", port))
			case <-mQuit.ClickedCh:
				s := g.Server()
				systray.Quit()
				s.Shutdown()
				fmt.Println("Quit Overflows...")
				return
			}
		}
	}()

	go func() {
		Server.Run(ctx)
	}()
}

func onExit(ctx context.Context) {
	g.Log().Info(ctx, "ShutDown App")
}

func openURL(url string) {
	var err error
	switch runtime.GOOS {
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	}
	if err != nil {
		zenity.Info(fmt.Sprintf("Please input the following URL in your browser:\n%s", url),
			zenity.Title("Failed to open UI"),
			zenity.InfoIcon)
	}
}
