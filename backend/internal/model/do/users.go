// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Users is the golang structure of table users for DAO operations like Where/Data.
type Users struct {
	g.Meta        `orm:"table:users, do:true"`
	Uid           interface{} //
	AvatarUrl     interface{} //
	Username      interface{} //
	Email         interface{} //
	Password      interface{} //
	CreatedAt     *gtime.Time //
	UpdatedAt     *gtime.Time //
	DeletedAt     *gtime.Time //
	LastLogin     *gtime.Time //
	LoginAttempts interface{} //
	Lock          interface{} //
	LockAt        *gtime.Time //
}
