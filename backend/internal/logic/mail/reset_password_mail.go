package mail

import (
	"context"
	"overflows/internal/model/entity"
	"overflows/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/wneessen/go-mail"
)

func resetPasswordEmailContent(ctx context.Context, user *entity.Users) (result string, err error) {
	v := g.View()
	v.SetDefaultFile("reset_password_mail.html")
	// Five minutes expiry time
	url, _ := g.Cfg().Get(ctx, "frontendURL")
	// Not checking token generation, would be bug here
	link := url.String() + "/api/v1/auth/password?uid=" + user.Uid + "&token=" + service.Token().GenerateMixedWithEx(ctx, user.Uid+"_password", 32, 300)
	result, err = v.ParseDefault(ctx, g.Map{
		"Username":           user.Username,
		"ChangePasswordLink": link,
	})
	return
}

func (s *sMailServer) SendResetPasswordEmail(ctx context.Context, user *entity.Users) (bool, error) {
	m := mail.NewMsg()
	if err := m.From(s.From); err != nil {
		g.Log("mail").Print(ctx, "Failed to set From address: %s", err)
		return false, err
	}
	if err := m.To(user.Email); err != nil {
		g.Log("mail").Print(ctx, "Failed to set To address: %s", err)
		return false, err
	}
	m.Subject(g.I18n().T(ctx, `resetPasswordSubject`))
	content, err := resetPasswordEmailContent(ctx, user)
	if err != nil {
		g.Log("mail").Print(ctx, "Failed to generate reset password email: %s", err)
		return false, err
	}
	m.SetBodyString(mail.TypeTextHTML, content)
	err = s.SendEmail(ctx, m)
	if err != nil {
		return false, err
	}
	return true, nil
}
