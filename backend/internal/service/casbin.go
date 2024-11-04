// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	casbin "github.com/casbin/casbin/v2"
)

type (
	ICasbin interface {
		DefaultEnforcer() *casbin.Enforcer
	}
)

var (
	localCasbin ICasbin
)

func Casbin() ICasbin {
	if localCasbin == nil {
		panic("implement not found for interface ICasbin, forgot register?")
	}
	return localCasbin
}

func RegisterCasbin(i ICasbin) {
	localCasbin = i
}
