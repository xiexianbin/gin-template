package auth

import (
	casbin "github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v2"
	"github.com/go-kit/kit/endpoint"
	_ "github.com/go-sql-driver/mysql"

	"github.com/xiexianbin/golib/logger"
)

var _ endpoint.Middleware

func Casbin(dbtype, conn string) (*casbin.Enforcer, error) {
	Apter, err := gormadapter.NewAdapter(dbtype, conn, true)
	if err != nil {
		return nil, err
	}
	e, err := casbin.NewEnforcer("conf/rbac_model.conf", Apter)
	if err != nil {
		return nil, err
	}
	if err := e.LoadPolicy(); err == nil {
		return e, err
	} else {
		logger.Infof("casbin rbac_model or policy init error, message: %v \r\n", err.Error())
		return nil, err
	}
}
