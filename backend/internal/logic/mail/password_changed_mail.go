package mail

import (
	"context"
	"overflows/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/wneessen/go-mail"
)

func passwordChangedEmailContent(ctx context.Context, user *entity.Users) (result string, err error) {
	v := g.View()
	v.SetDefaultFile("password_changed_mail.html")
	result, err = v.ParseDefault(ctx, g.Map{
		"Username":    user.Username,
		"ChangedTime": gtime.Now().String(),
	})
	return
}

func (s *sMailServer) SendPasswordChangedEmail(ctx context.Context, user *entity.Users) (bool, error) {
	m := mail.NewMsg()
	if err := m.From(s.From); err != nil {
		g.Log("mail").Print(ctx, "Failed to set From address: %s", err)
		return false, err
	}
	if err := m.To(user.Email); err != nil {
		g.Log("mail").Print(ctx, "Failed to set To address: %s", err)
		return false, err
	}
	m.Subject(g.I18n().T(ctx, `passwordChangedSubject`))
	content, err := passwordChangedEmailContent(ctx, user)
	if err != nil {
		g.Log("mail").Print(ctx, "Failed to generate password changed email: %s", err)
		return false, err
	}
	m.SetBodyString(mail.TypeTextHTML, content)
	err = s.SendEmail(ctx, m)
	if err != nil {
		return false, err
	}
	return true, nil
}
