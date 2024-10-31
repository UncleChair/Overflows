package main

import (
	"fmt"
	"os"
	"os/exec"
	_ "overflows/internal/packed"
	"runtime"
	"time"

	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"

	"overflows/internal/cmd"

	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	_ "github.com/gogf/gf/contrib/drivers/sqlite/v2"
)

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetTemplateIcon(icon.Data, icon.Data)
	systray.SetTitle("Overflows")
	systray.SetTooltip("Overflows")

	go func() {
		s := g.Server()
		mUrl := systray.AddMenuItem("Open UI", "my home")
		mQuit := systray.AddMenuItem("退出", "Quit the whole app")

		// Sets the icon of a menu item. Only available on Mac.
		// mQuit.SetIcon(icon.Data)

		systray.AddSeparator()

		for {
			select {
			case <-mUrl.ClickedCh:
				port := s.GetListenedPort()
				err := openURL(fmt.Sprintf("http://127.0.0.1:%d", port))
				if err != nil {
					fmt.Println("Failed to open URL:", err)
				}
			case <-mQuit.ClickedCh:
				systray.Quit()
				s.Shutdown()
				fmt.Println("Quit2 now...")
				return
			}
		}
	}()

	go func() {
		cmd.Main.Run(gctx.GetInitCtx())
	}()
}

func onExit() {
	now := time.Now()
	os.WriteFile(fmt.Sprintf(`on_exit_%d.txt`, now.UnixNano()), []byte(now.String()), 0644)
}

func openURL(url string) error {
	var err error
	switch runtime.GOOS {
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	default:
		return fmt.Errorf("unsupported platform")
	}
	return err
}
