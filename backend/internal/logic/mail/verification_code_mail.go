package mail

import (
	"context"
	"overflows/internal/model/entity"
	"overflows/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/wneessen/go-mail"
)

func verificationCodeEmailContent(ctx context.Context, user *entity.Users) (result string, err error) {
	v := g.View()
	v.SetDefaultFile("verification_code_mail.html")
	// Five seconds expiry time
	token := service.Token().GenerateDigitWithEx(ctx, user.Uid, 6, 300)
	result, err = v.ParseDefault(ctx, g.Map{
		"Username": user.Username,
		"Code":     token,
	})
	return
}

func (s *sMailServer) SendVerificationCodeEmail(ctx context.Context, user *entity.Users) (bool, error) {
	m := mail.NewMsg()
	if err := m.From(s.From); err != nil {
		g.Log("mail").Print(ctx, "Failed to set From address: %s", err)
		return false, err
	}
	if err := m.To(user.Email); err != nil {
		g.Log("mail").Print(ctx, "Failed to set To address: %s", err)
		return false, err
	}
	m.Subject(g.I18n().T(ctx, `verificationCodeSubject`))
	content, err := verificationCodeEmailContent(ctx, user)
	if err != nil {
		g.Log("mail").Print(ctx, "Failed to generate varification code email: %s", err)
		return false, err
	}
	m.SetBodyString(mail.TypeTextHTML, content)
	err = s.SendEmail(ctx, m)
	if err != nil {
		return false, err
	}
	return true, nil
}
