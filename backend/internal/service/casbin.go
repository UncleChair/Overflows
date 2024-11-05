package service

import (
	"context"

	casbin "github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gdbadapter "github.com/jxo-me/gdb-adapter"
)

var CasbinInstance *sCasbin

type sCasbin struct {
	Adapter  interface{}
	Model    model.Model
	Enforcer *casbin.Enforcer
}

func InitCasbin(ctx context.Context, groupname string) {
	c := &sCasbin{}
	c.Model = c.defaultModel()
	c.Adapter, _ = gdbadapter.NewAdapter(ctx, groupname)
	c.Enforcer, _ = casbin.NewEnforcer(c.Model, c.Adapter)
	CasbinInstance = c
}

func Casbin() *sCasbin {
	return CasbinInstance
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
