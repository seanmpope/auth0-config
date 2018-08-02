package main

import (
	"io/ioutil"

	"github.com/zenoss/go-auth0/auth0/authz"
	"gopkg.in/yaml.v2"
)

type Globals map[string]interface{}

type Auth0Config struct {
	Globals       Globals
	Authorization AuthorizationConfig
	Management    ManagementConfig
}

// func (cfg *Auth0Config) UnmarshalYAML(unmarshal func(interface{}) error) error {
// 	g := make(map[string]interface{})
// 	err := unmarshal(g)
// 	if err != nil {
// 		return err
// 	}
// 	cfg.Globals = g
// 	fmt.Printf("\n%+v\n", g)
// 	return nil
// }

func NewAuth0Config(fileName string) (Auth0Config, error) {
	var cfg Auth0Config
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		return cfg, err
	}
	err = yaml.Unmarshal(bs, &cfg)
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}

type ManagementConfig struct {
}

type AuthorizationConfig struct {
	Groups      []authz.Group
	Roles       []authz.Role
	Permissions []authz.Permission
}

type RolesClient interface {
	Create(authz.Role) (authz.Role, error)
}

type PermissionsClient interface {
	Create(authz.Permission) (authz.Permission, error)
}

type OperationResult struct {
	Request interface{}
	Result  interface{}
}

// func NewAuthorizationConfig(fileName string) (AuthorizationConfig, error) {
// 	var cfg AuthorizationConfig
// 	bs, err := ioutil.ReadFile(fileName)
// 	if err != nil {
// 		return cfg, err
// 	}
// 	err = yaml.Unmarshal(bs, &cfg)
// 	if err != nil {
// 		return cfg, err
// 	}

// 	return cfg, nil
// }

func (cfg *AuthorizationConfig) CreateRoles(rc RolesClient) []OperationResult {
	results := make([]OperationResult, len(cfg.Roles))
	for n, role := range cfg.Roles {
		r, err := rc.Create(role)
		if err != nil {
			results[n] = OperationResult{
				Request: role,
				Result:  err,
			}
		} else {
			results[n] = OperationResult{
				Request: role,
				Result:  r,
			}
		}
	}
	return results
}

func (cfg *AuthorizationConfig) CreatePermissions(pc PermissionsClient) []OperationResult {
	results := make([]OperationResult, len(cfg.Permissions))
	for n, perm := range cfg.Permissions {
		p, err := pc.Create(perm)
		if err != nil {
			results[n] = OperationResult{
				Request: perm,
				Result:  err,
			}
		} else {
			results[n] = OperationResult{
				Request: perm,
				Result:  p,
			}
		}
	}
	return results
}
