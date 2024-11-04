// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

type (
	IBcrypt interface {
		Generate(password string) (hashedPassword string, err error)
		Compare(hashPassword string, password string) bool
	}
)

var (
	localBcrypt IBcrypt
)

func Bcrypt() IBcrypt {
	if localBcrypt == nil {
		panic("implement not found for interface IBcrypt, forgot register?")
	}
	return localBcrypt
}

func RegisterBcrypt(i IBcrypt) {
	localBcrypt = i
}
