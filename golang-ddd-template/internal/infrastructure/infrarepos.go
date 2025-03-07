package infrastructure

import (
	rolerepo "github.com/andriykutsevol/WeatherServer/internal/domain/user/role"
	menurepo "github.com/andriykutsevol/WeatherServer/internal/domain/menu"
	rolemenurepo "github.com/andriykutsevol/WeatherServer/internal/domain/user/rolemenu"
	menuactionrepo "github.com/andriykutsevol/WeatherServer/internal/domain/menu/menuaction"
	menuactionresourcerepo "github.com/andriykutsevol/WeatherServer/internal/domain/menu/menuactionresource"
	userrepo "github.com/andriykutsevol/WeatherServer/internal/domain/user"
	userrolerepo "github.com/andriykutsevol/WeatherServer/internal/domain/user/userrole"
	authrepo "github.com/andriykutsevol/WeatherServer/internal/domain/auth"
	rbacrepo "github.com/andriykutsevol/WeatherServer/internal/domain/rbac"
	weatherrepo "github.com/andriykutsevol/WeatherServer/internal/domain/weather"	
)


type InfraRepos struct {
	RoleRepository rolerepo.Repository
	MenuRepository menurepo.Repository
	RoleMeuRepository rolemenurepo.Repository
	MenuActionRepository menuactionrepo.Repository
	MenuActionResourceRepository menuactionresourcerepo.Repository
	UserRepository userrepo.Repository
	UserRoleRepository userrolerepo.Repository
	AuthRepository authrepo.Repository
	RbacRepository rbacrepo.Repository
	WeatherRepository weatherrepo.Repository
}
