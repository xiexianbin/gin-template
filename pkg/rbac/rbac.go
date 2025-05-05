// Copyright 2025 xiexianbin<me@xiexianbin.cn>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// ref https://blog.csdn.net/LeoForBest/article/details/133610889

package rbac

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

var basicModelConf = `[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[role_definition]
g = _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = g(r.sub, p.sub) && keyMatch2(r.obj,p.obj) && r.act == p.act
`

var tenantModelConf = `
[request_definition]
r = tenant, sub, obj, act

[policy_definition]
p = tenant, sub, obj, act, eft

[role_definition]
g = _, _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = (g(r.sub, "super_admin", "*") || r.tenant == p.tenant) &&
    g(r.sub, p.sub, p.tenant) &&
    keyMatch3(r.obj, p.obj) &&
    r.act == p.act
`

type RBAC struct {
	enforcer *casbin.Enforcer
	adapter  *gormadapter.Adapter
}

func New(db *gorm.DB) (*RBAC, error) {
	a, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		return nil, err
	}
	m, err := model.NewModelFromString(tenantModelConf)
	if err != nil {
		return nil, err
	}
	e, err := casbin.NewEnforcer(m, a)
	if err != nil {
		return nil, err
	}
	return &RBAC{adapter: a, enforcer: e}, nil
}

type User struct {
	UserName string
	Roles    []string
}

// GetUsers retrieves users along with their assigned roles.
func (r *RBAC) GetUsers() (users []User, err error) {
	p, err := r.enforcer.GetGroupingPolicy()
	if err != nil {
		return
	}
	usernameUser := make(map[string]*User, 0)
	for _, _p := range p {
		username, usergroup := _p[0], _p[1]
		if v, ok := usernameUser[username]; ok {
			usernameUser[username].Roles = append(v.Roles, usergroup)
		} else {
			usernameUser[username] = &User{UserName: username, Roles: []string{usergroup}}
		}
	}
	for _, v := range usernameUser {
		users = append(users, *v)
	}
	return
}

// GetRoles retrieves all defined roles.
func (r *RBAC) GetRoles() ([]string, error) {
	return r.enforcer.GetAllRoles()
}

// UpdateUserRole adds a user to a role group; if the role doesn't exist, it will be created.
func (r *RBAC) UpdateUserRole(username, role string) error {
	_, err := r.enforcer.AddGroupingPolicy(username, role)
	if err != nil {
		return err
	}
	return r.enforcer.SavePolicy()
}

// DeleteUserRole removes a user from a role group.
func (r *RBAC) DeleteUserRole(username, role string) error {
	_, err := r.enforcer.RemoveGroupingPolicy(username, role)
	if err != nil {
		return err
	}
	return r.enforcer.SavePolicy()
}

// RolePolicy represents (Role, Url, Method) corresponding to (v0, v1, v2) in the `CasbinRule` table.
type RolePolicy struct {
	Role   string `gorm:"column:v0"` // Role name
	Url    string `gorm:"column:v1"` // API URL
	Method string `gorm:"column:v2"` // HTTP REST method
}

// GetRolePolicy retrieves all role permissions.
func (r *RBAC) GetRolePolicy() (roles []RolePolicy, err error) {
	err = r.adapter.GetDb().Model(&CasbinRule{}).Where("ptype = 'p'").Find(&roles).Error
	if err != nil {
		return nil, err
	}
	return
}

// CreateRolePolicy creates a new role permission; existing ones will be ignored.
func (r *RBAC) CreateRolePolicy(rpolicy RolePolicy) error {
	// Do not operate directly on the database; use enforcer to simplify operations.
	err := r.enforcer.LoadPolicy()
	if err != nil {
		return err
	}
	_, err = r.enforcer.AddPolicy(rpolicy.Role, rpolicy.Url, rpolicy.Method)
	if err != nil {
		return err
	}
	return r.enforcer.SavePolicy()
}

// UpdateRolePolicy updates an existing role permission.
func (r *RBAC) UpdateRolePolicy(old, new RolePolicy) error {
	_, err := r.enforcer.UpdatePolicy([]string{old.Role, old.Url, old.Method},
		[]string{new.Role, new.Url, new.Method})
	if err != nil {
		return err
	}
	return r.enforcer.SavePolicy()
}

// DeleteRolePolicy deletes a role permission.
func (r *RBAC) DeleteRolePolicy(rpolicy RolePolicy) error {
	_, err := r.enforcer.RemovePolicy(rpolicy.Role, rpolicy.Url, rpolicy.Method)
	if err != nil {
		return err
	}
	return r.enforcer.SavePolicy()
}

// CanAccess checks whether a user has access to a specific policy.
func (r *RBAC) CanAccess(username, url, method string) (ok bool, err error) {
	return r.enforcer.Enforce(username, url, method)
}
