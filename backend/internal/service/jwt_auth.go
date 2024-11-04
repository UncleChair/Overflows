package service

import (
	"context"
	"overflows/internal/consts"
	"overflows/internal/dao"
	"time"

	jwt "github.com/gogf/gf-jwt/v2"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

var authService *jwt.GfJWTMiddleware

func JWTAuth() *jwt.GfJWTMiddleware {
	return authService
}

func init() {
	auth := jwt.New(&jwt.GfJWTMiddleware{
		Realm:           "user zone",
		Key:             []byte(consts.JWTKey),
		Timeout:         time.Hour * 24,
		MaxRefresh:      time.Hour * 48,
		IdentityKey:     "uid",
		TokenLookup:     "header: Authorization, query: token, cookie: jwt",
		TokenHeadName:   "Bearer",
		TimeFunc:        time.Now,
		Authenticator:   Authenticator,
		Unauthorized:    Unauthorized,
		PayloadFunc:     PayloadFunc,
		IdentityHandler: IdentityHandler,
	})
	authService = auth
}

// PayloadFunc is a callback function that will be called during login.
// Using this function it is possible to add additional payload data to the webtoken.
// The data is then made available during requests via c.Get("JWT_PAYLOAD").
// Note that the payload is not encrypted.
// The attributes mentioned on jwt.io can't be used as keys for the map.
// Optional, by default no additional data will be set.
func PayloadFunc(data interface{}) jwt.MapClaims {
	claims := jwt.MapClaims{}
	params := data.(map[string]interface{})
	if len(params) > 0 {
		for k, v := range params {
			claims[k] = v
		}
	}
	return claims
}

// IdentityHandler get the identity from JWT and set the identity for every request
// Using this function, by r.GetParam(identityKey) get identity
// The user identity is also injected into customized context here
func IdentityHandler(ctx context.Context) interface{} {
	// ToDo: Handle user access level here
	claims := jwt.ExtractClaims(ctx)
	Context().SetUserUid(ctx, claims[authService.IdentityKey].(string))
	return claims[authService.IdentityKey]
}

// Unauthorized is used to define customized Unauthorized callback function.
func Unauthorized(ctx context.Context, code int, message string) {
	r := g.RequestFromCtx(ctx)
	r.Response.Status = code
	originalError := gerror.New(message)
	// error handling here
	if gerror.Equal(originalError, AuthError.InternalError) {
		code = 1
		r.Response.Status = 500
	}
	if gerror.Equal(originalError, AuthError.NoUser) {
		code = 1
		r.Response.Status = 401
	}
	if gerror.Equal(originalError, AuthError.WrongCredential) {
		code = 2
		r.Response.Status = 401
	}
	if gerror.Equal(originalError, AuthError.TooManyAttempts) {
		code = 1
		r.Response.Status = 403
	}
	r.Response.WriteJson(g.Map{
		"code":    code,
		"message": message,
	})
	r.ExitAll()
}

// Authenticator is used to validate login parameters.
// It must return user data as user identifier, it will be stored in Claim Array.
// if your identityKey is 'id', your user data must have 'id'
// Check error (e) to determine the appropriate error message.
func Authenticator(ctx context.Context) (interface{}, error) {
	r := g.RequestFromCtx(ctx)
	var user *gdb.Model
	identity := r.Get("username").String()
	if identity == "" {
		identity = r.Get("email").String()
		user = dao.Users.Ctx(ctx).Where(g.Map{
			dao.Users.Columns().Email: identity,
		})
	} else {
		user = dao.Users.Ctx(ctx).Where(g.Map{
			dao.Users.Columns().Username: identity,
		})
	}
	exist, _ := user.Count()
	if exist == 0 {
		return nil, AuthError.NoUser
	}

	result, err := user.Fields(
		dao.Users.Columns().Password,
		dao.Users.Columns().Uid,
		dao.Users.Columns().Username,
		dao.Users.Columns().Lock,
		dao.Users.Columns().LoginAttempts).One()
	if err != nil {
		return nil, AuthError.InternalError
	}

	pass := result[dao.Users.Columns().Password]
	uid := result[dao.Users.Columns().Uid]
	username := result[dao.Users.Columns().Username]
	lock := result[dao.Users.Columns().Lock]
	loginAttempts := result[dao.Users.Columns().LoginAttempts]
	if lock.Bool() {
		return nil, AuthError.TooManyAttempts
	}
	if !Bcrypt().Compare(pass.String(), r.Get("password").String()) {
		if loginAttempts.Int() > 3 {
			user.Data(g.Map{
				dao.Users.Columns().LastLogin: gtime.Now(),
				dao.Users.Columns().Lock:      true,
			}).Update()
		} else {
			user.Data(g.Map{
				dao.Users.Columns().LastLogin:     gtime.Now(),
				dao.Users.Columns().LoginAttempts: loginAttempts.Int() + 1,
			}).Update()
		}
		return nil, AuthError.WrongCredential
	} else {
		user.Data(g.Map{
			dao.Users.Columns().LastLogin:     gtime.Now(),
			dao.Users.Columns().LoginAttempts: 0,
		}).Update()
		return g.Map{
			"uid":      uid,
			"username": username,
		}, nil
	}
}

var AuthError = struct {
	InternalError   error
	NoUser          error
	WrongCredential error
	TooManyAttempts error
}{
	InternalError:   gerror.New("Something wrong with server"),
	NoUser:          gerror.New("No related users"),
	WrongCredential: gerror.New("Wrong username or password"),
	TooManyAttempts: gerror.New("Too many wrong attempts, please try later"),
}
