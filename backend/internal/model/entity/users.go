// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Users is the golang structure for table users.
type Users struct {
	Uid           string      `json:"uid"           orm:"uid"            ` //
	AvatarUrl     string      `json:"avatarUrl"     orm:"avatar_url"     ` //
	Username      string      `json:"username"      orm:"username"       ` //
	Email         string      `json:"email"         orm:"email"          ` //
	Password      string      `json:"password"      orm:"password"       ` //
	CreatedAt     *gtime.Time `json:"createdAt"     orm:"created_at"     ` //
	UpdatedAt     *gtime.Time `json:"updatedAt"     orm:"updated_at"     ` //
	DeletedAt     *gtime.Time `json:"deletedAt"     orm:"deleted_at"     ` //
	LastLogin     *gtime.Time `json:"lastLogin"     orm:"last_login"     ` //
	LoginAttempts int         `json:"loginAttempts" orm:"login_attempts" ` //
	Lock          bool        `json:"lock"          orm:"lock"           ` //
	LockAt        *gtime.Time `json:"lockAt"        orm:"lock_at"        ` //
}
