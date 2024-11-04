// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	IVerification interface {
		CheckUsername(ctx context.Context, username string) error
		CheckUserEmail(ctx context.Context, email string) error
	}
)

var (
	localVerification IVerification
)

func Verification() IVerification {
	if localVerification == nil {
		panic("implement not found for interface IVerification, forgot register?")
	}
	return localVerification
}

func RegisterVerification(i IVerification) {
	localVerification = i
}
