package token

import (
	"context"
	"overflows/internal/service"
	"time"

	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/util/grand"
)

type sToken struct{}

func init() {
	service.RegisterToken(New())
}

func New() service.IToken {
	return &sToken{}
}

// Generate token and store it in cache.
// Key is the key in cache, style is the type of token, l is the length of token.
// Type can be choose from digit, letter or mixed.
func generate(ctx context.Context, key string, style string, length int, timeout int64) (token string) {

	if key == "" || length <= 0 {
		panic("Wrong input parameter")
	}
	if style == "digit" {
		token = grand.Digits(length)
	} else if style == "letter" {
		token = grand.Letters(length)
	} else if style == "mixed" {
		token = grand.S(length)
	} else {
		panic("Wrong token style")
	}
	var err error
	if timeout > 0 {
		err = gcache.Set(ctx, key, token, time.Duration(timeout)*time.Second)
	} else {
		err = gcache.Set(ctx, key, token, 0)
	}
	if err != nil {
		panic(err)
	}

	return
}

func (s *sToken) GenerateDigit(ctx context.Context, key string, length int) string {
	return generate(ctx, key, "digit", length, 0)
}
func (s *sToken) GenerateDigitWithEx(ctx context.Context, key string, length int, timeout int64) string {
	return generate(ctx, key, "digit", length, timeout)
}
func (s *sToken) GenerateLetter(ctx context.Context, key string, length int) string {
	return generate(ctx, key, "letter", length, 0)
}
func (s *sToken) GenerateLetterWithEx(ctx context.Context, key string, length int, timeout int64) string {
	return generate(ctx, key, "letter", length, timeout)
}
func (s *sToken) GenerateMixed(ctx context.Context, key string, length int) string {
	return generate(ctx, key, "mixed", length, 0)
}
func (s *sToken) GenerateMixedWithEx(ctx context.Context, key string, length int, timeout int64) string {
	return generate(ctx, key, "mixed", length, timeout)
}

func (s *sToken) Verify(ctx context.Context, key string, token string) (bool, error) {
	value, err := gcache.Get(ctx, key)
	if err != nil {
		return false, err
	}
	if value.String() != token {
		return false, nil
	}
	gcache.Remove(ctx, key)
	return true, nil
}

func (s *sToken) VerifyWithoutDel(ctx context.Context, key string, token string) (bool, error) {
	value, err := gcache.Get(ctx, key)
	if err != nil {
		return false, err
	}
	if value.String() != token {
		return false, nil
	}
	return true, nil
}

func (s *sToken) Get(ctx context.Context, key string) (value string, err error) {
	result, err := gcache.Get(ctx, key)
	if err == nil {
		value = result.String()
	}
	return
}

func (s *sToken) UserCSSAccessKey(userUid string, projectId string, versionId string) (key string) {
	key = "css_access:" + "u" + userUid + ":" + "p" + projectId + ":" + "v" + versionId
	return
}
