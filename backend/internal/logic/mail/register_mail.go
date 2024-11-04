package mail

import (
	"context"
	"overflows/internal/model/entity"
	"overflows/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/wneessen/go-mail"
)

func registerEmailContent(ctx context.Context, user *entity.Users) (result string, err error) {
	v := g.View()
	v.SetDefaultFile("register_mail.html")
	// Five minutes expiry time
	url, _ := g.Cfg().Get(ctx, "frontendURL")
	// Not checking token generation, would be bug here
	verifyLink := url.String() + "/api/v1/auth/email?uid=" + user.Uid + "&token=" + service.Token().GenerateMixedWithEx(ctx, user.Uid+"_register", 12, 300)
	result, err = v.ParseDefault(ctx, g.Map{
		"Username":        user.Username,
		"EmailVerifyLink": verifyLink,
	})
	return
}

func (s *sMailServer) SendRegisterEmail(ctx context.Context, user *entity.Users) (bool, error) {
	m := mail.NewMsg()
	if err := m.From(s.From); err != nil {
		g.Log("mail").Print(ctx, "Failed to set From address: %s", err)
		return false, err
	}
	if err := m.To(user.Email); err != nil {
		g.Log("mail").Print(ctx, "Failed to set To address: %s", err)
		return false, err
	}
	m.Subject(g.I18n().T(ctx, `registerSubject`))
	content, err := registerEmailContent(ctx, user)
	if err != nil {
		g.Log("mail").Print(ctx, "Failed to generate email verification email: %s", err)
		return false, err
	}
	m.SetBodyString(mail.TypeTextHTML, content)
	err = s.SendEmail(ctx, m)
	if err != nil {
		return false, err
	}
	return true, nil
}
