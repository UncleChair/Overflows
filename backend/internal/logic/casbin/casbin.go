package casbin

import (
	"overflows/internal/service"

	casbin "github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
)

type sCasbin struct {
	Adapter  interface{}
	Model    model.Model
	Enforcer *casbin.Enforcer
}

func init() {
	// ctx := gctx.GetInitCtx()
	c := New()
	c.Model = c.defaultModel()
	c.Adapter = nil
	c.Enforcer = nil
	service.RegisterCasbin(c)
}

func New() *sCasbin {
	return &sCasbin{}
}

func (s *sCasbin) defaultModel() (m model.Model) {
	m, _ = model.NewModelFromString(`
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[role_definition]
g = _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act
`)
	return
}

func (s *sCasbin) DefaultEnforcer() *casbin.Enforcer {
	return s.Enforcer
}
