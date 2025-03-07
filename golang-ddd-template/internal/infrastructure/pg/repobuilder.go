package pg

import(
	"github.com/andriykutsevol/WeatherServer/internal/infrastructure"
	"github.com/andriykutsevol/WeatherServer/internal/infrastructure/pg/storage"
	menuinfra "github.com/andriykutsevol/WeatherServer/internal/infrastructure/pg/menu"
	menuactioninfra "github.com/andriykutsevol/WeatherServer/internal/infrastructure/pg/menuaction"
	menuactionresourceinfra "github.com/andriykutsevol/WeatherServer/internal/infrastructure/pg/menuactionresource"
	roleinfra "github.com/andriykutsevol/WeatherServer/internal/infrastructure/pg/role"
	rolemenuinfra "github.com/andriykutsevol/WeatherServer/internal/infrastructure/pg/rolemenu"
	userinfra "github.com/andriykutsevol/WeatherServer/internal/infrastructure/pg/user"
	userroleinfra "github.com/andriykutsevol/WeatherServer/internal/infrastructure/pg/userrole"

	authinfra "github.com/andriykutsevol/WeatherServer/internal/infrastructure/redis/auth"
	rbacinfra "github.com/andriykutsevol/WeatherServer/internal/infrastructure/rbac"

	weatherinfra "github.com/andriykutsevol/WeatherServer/internal/infrastructure/pg/weather"

)




func BuildRespositories (
	casbinservice storage.DatabaseService, 
	weatherservice storage.DatabaseService) (*infrastructure.InfraRepos, error){

	infrarepos := &infrastructure.InfraRepos{}

	infrarepos.MenuRepository = menuinfra.NewRepository(casbinservice)
	infrarepos.MenuActionRepository = menuactioninfra.NewRepository(casbinservice)
	infrarepos.MenuActionResourceRepository = menuactionresourceinfra.NewRepository(casbinservice, infrarepos.MenuActionRepository)
	infrarepos.RoleRepository = roleinfra.NewRepository(casbinservice)
	infrarepos.RoleMeuRepository = rolemenuinfra.NewRepository(casbinservice)
	infrarepos.UserRepository = userinfra.NewRepository(casbinservice)
	infrarepos.UserRoleRepository = userroleinfra.NewRepository(casbinservice)

	authRedisStorage := &authinfra.AuthStorage{}
	infrarepos.AuthRepository = authinfra.NewRepository(authRedisStorage)

	infrarepos.RbacRepository = rbacinfra.NewRepository(
		infrarepos.RoleRepository, 
		infrarepos.RoleMeuRepository, 
		infrarepos.MenuActionResourceRepository, 
		infrarepos.UserRepository, 
		infrarepos.UserRoleRepository)

	infrarepos.WeatherRepository = weatherinfra.NewRepository(weatherservice)	

	return infrarepos, nil
}