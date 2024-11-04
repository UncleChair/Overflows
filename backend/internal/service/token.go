// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	IToken interface {
		GenerateDigit(ctx context.Context, key string, length int) string
		GenerateDigitWithEx(ctx context.Context, key string, length int, timeout int64) string
		GenerateLetter(ctx context.Context, key string, length int) string
		GenerateLetterWithEx(ctx context.Context, key string, length int, timeout int64) string
		GenerateMixed(ctx context.Context, key string, length int) string
		GenerateMixedWithEx(ctx context.Context, key string, length int, timeout int64) string
		Verify(ctx context.Context, key string, token string) (bool, error)
		VerifyWithoutDel(ctx context.Context, key string, token string) (bool, error)
		Get(ctx context.Context, key string) (value string, err error)
		UserCSSAccessKey(userUid string, projectId string, versionId string) (key string)
	}
)

var (
	localToken IToken
)

func Token() IToken {
	if localToken == nil {
		panic("implement not found for interface IToken, forgot register?")
	}
	return localToken
}

func RegisterToken(i IToken) {
	localToken = i
}
