package mail

import (
	"context"
	"overflows/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/wneessen/go-mail"
)

type sMailServer struct {
	host     string
	port     int
	username string
	password string
	From     string
}

func init() {
	service.RegisterMailServer(New())
}

func New() service.IMailServer {
	ctx := gctx.GetInitCtx()
	host, _ := g.Cfg().Get(ctx, "mail.host")
	port, _ := g.Cfg().Get(ctx, "mail.port")
	username, _ := g.Cfg().Get(ctx, "mail.username")
	password, _ := g.Cfg().Get(ctx, "mail.password")
	from, _ := g.Cfg().Get(ctx, "mail.from")
	return &sMailServer{
		host:     host.String(),
		port:     port.Int(),
		username: username.String(),
		password: password.String(),
		From:     from.String(),
	}
}

func (s *sMailServer) NewClient() (client *mail.Client, err error) {
	client, err = mail.NewClient(
		s.host,
		mail.WithPort(s.port),
		mail.WithSMTPAuth(mail.SMTPAuthPlain),
		mail.WithUsername(s.username),
		mail.WithPassword(s.password),
	)
	return
}

func (s *sMailServer) SendEmail(ctx context.Context, m *mail.Msg) (err error) {
	client, err := s.NewClient()
	if err != nil {
		g.Log("mail").Print(ctx, "Failed to create new mail client: %s", err)
		return err
	}
	if err := client.DialAndSend(m); err != nil {
		g.Log("mail").Print(ctx, "Failed to send mail: %s", err)
		return err
	}
	return nil
}
