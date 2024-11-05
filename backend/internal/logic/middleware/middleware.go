package middleware

import (
	"mime"
	"net/http"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/net/ghttp"

	"overflows/api"
	"overflows/internal/consts"
	"overflows/internal/model"
	"overflows/internal/service"
)

type sMiddleware struct{}

const (
	contentTypeEventStream  = "text/event-stream"
	contentTypeOctetStream  = "application/octet-stream"
	contentTypeMixedReplace = "multipart/x-mixed-replace"
)

var (
	// streamContentType is the content types for stream response.
	streamContentType = []string{contentTypeEventStream, contentTypeOctetStream, contentTypeMixedReplace}
)

func init() {
	service.RegisterMiddleware(New())
}

func New() service.IMiddleware {
	return &sMiddleware{}
}

// Auth validates the request to allow only signed-in users visit.
func (s *sMiddleware) Auth(r *ghttp.Request) {
	service.JWTAuth().MiddlewareFunc()(r)
	r.Middleware.Next()
}

// CORS allows Cross-origin resource sharing.
func (s *sMiddleware) CORS(r *ghttp.Request) {
	corsOptions := r.Response.DefaultCORSOptions()
	// Set allow all cors when dev or staging (only in docker)
	if consts.AppEnv == "prod" {
		corsOptions.AllowDomain = []string{"unclechair.vip", "silicious.net"}
		corsOptions.AllowOrigin = r.Request.Header.Get("Origin")
		corsOptions.ExposeHeaders = "Trace-Id"
	} else {
		corsOptions.AllowDomain = []string{"localhost"}
		corsOptions.AllowOrigin = "http://localhost:3000"
		corsOptions.ExposeHeaders = "Trace-Id"
	}
	if !r.Response.CORSAllowedOrigin(corsOptions) {
		r.Response.WriteStatus(403)
	}
	r.Response.CORS(corsOptions)
	r.Middleware.Next()
}

// i18n language setting.
func (s *sMiddleware) Language(r *ghttp.Request) {
	lang := r.GetQuery("lang", "zh-CN").String()
	ctx := gi18n.WithLanguage(r.Context(), lang)
	r.SetCtx(ctx)
	r.Middleware.Next()
}

// Prevent context canceled error.
func (s *sMiddleware) NeverDoneCtx(r *ghttp.Request) {
	r.SetCtx(r.GetNeverDoneCtx())
	r.Middleware.Next()
}

// Only admin could pass
func (s *sMiddleware) IsAdmin(r *ghttp.Request) {
	uid := service.Context().Use(r.Context()).User.Uid
	e := service.Casbin().DefaultEnforcer()
	isAdmin, _ := e.HasRoleForUser(uid, "admin")
	if !isAdmin {
		r.Response.Status = 401
		r.Response.WriteJson(g.Map{
			"code":    1,
			"message": "Do not try it if you are not admin",
		})
		r.ExitAll()
	}
	r.Middleware.Next()
}

// ResponseHandler is the middleware handling response object and its error.
func (s *sMiddleware) ResponseHandler(r *ghttp.Request) {
	r.Middleware.Next()

	// There's custom buffer content, it then exits current handler.
	if r.Response.BufferLength() > 0 {
		return
	}

	// It does not output common response content if it is stream response.
	mediaType, _, _ := mime.ParseMediaType(r.Response.Header().Get("Content-Type"))
	for _, ct := range streamContentType {
		if mediaType == ct {
			return
		}
	}

	r.Response.Status = service.Context().Use(r.Context()).HttpStatus

	var (
		msg  string
		err  = r.GetError()
		res  = r.GetHandlerResponse()
		code = gerror.Code(err)
	)

	if err != nil {
		if code == gcode.CodeNil {
			code = gcode.CodeInternalError
		}
		msg = err.Error()
	} else {
		if r.Response.Status >= 300 || r.Response.Status < 200 {
			msg = http.StatusText(r.Response.Status)
			switch r.Response.Status {
			case http.StatusNotFound:
				code = gcode.CodeNotFound
			case http.StatusForbidden:
				code = gcode.CodeNotAuthorized
			default:
				code = gcode.CodeUnknown
			}
			// It creates error as it can be retrieved by other middlewares.
			err = gerror.NewCode(code, msg)
			r.SetError(err)
		} else {
			code = gcode.CodeOK
		}
	}
	r.Response.WriteJson(api.CommonRes{
		Code:    code.Code(),
		Message: msg,
		Data:    res,
	})
}

func (s *sMiddleware) Ctx(r *ghttp.Request) {
	customCtx := &model.Context{
		HttpStatus: 200,
		User: &model.ContextUser{
			Uid: "",
		},
		Data: make(g.Map),
	}
	service.Context().Init(r, customCtx)

	r.Assigns(g.Map{
		"Context": customCtx,
	})
	r.Middleware.Next()
}
