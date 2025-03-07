package mongo


import (
	"log"
	"github.com/andriykutsevol/WeatherServer/internal/infrastructure"
	"github.com/andriykutsevol/WeatherServer/internal/infrastructure/mongo/storage"
	roleinfra "github.com/andriykutsevol/WeatherServer/internal/infrastructure/mongo/role"
	menuinfra "github.com/andriykutsevol/WeatherServer/internal/infrastructure/mongo/menu"
	rolemenuinfra "github.com/andriykutsevol/WeatherServer/internal/infrastructure/mongo/rolemenu"
	menuactioninfra "github.com/andriykutsevol/WeatherServer/internal/infrastructure/mongo/menuaction"
	menuactionresourceinfra "github.com/andriykutsevol/WeatherServer/internal/infrastructure/mongo/menuactionresource"
	userinfra "github.com/andriykutsevol/WeatherServer/internal/infrastructure/mongo/user"
	userroleinfra "github.com/andriykutsevol/WeatherServer/internal/infrastructure/mongo/userrole"
	authinfra "github.com/andriykutsevol/WeatherServer/internal/infrastructure/redis/auth"
	rbacinfra "github.com/andriykutsevol/WeatherServer/internal/infrastructure/rbac"
	weatherinfra "github.com/andriykutsevol/WeatherServer/internal/infrastructure/mongo/weather"
)





func BuildRespositories () (*infrastructure.InfraRepos, error){
	infrarepos := &infrastructure.InfraRepos{}


	//-------------------------------------------------------
	// Connect to MongoDB

	// Establish MongoDB connection
	mongoclient, err := storage.ConnectToMongoDB()
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}

	// Create a new instance of DatabaseService with the MongoDB client injected
	mongodbService := storage.NewDatabaseService(mongoclient)
	mdb := mongodbService.GetDatabase("weather")


	//-------------------------------------------------------
	// RoleRepository
	roleStorage := &storage.MongoStorage{}
	roleCollection := mdb.GetCollection("role")
	roleStorage.SetCollection(roleCollection)
	infrarepos.RoleRepository = roleinfra.NewRepository(roleStorage)

	//-------------------------------------------------------
	// MenuRepository
	menuStorage := &storage.MongoStorage{}
	menuCollection := mdb.GetCollection("menu")
	menuStorage.SetCollection(menuCollection)
	infrarepos.MenuRepository = menuinfra.NewRepository(menuStorage)

	//-------------------------------------------------------
	// RolemenuRepository
	rolemenuStorage := &storage.MongoStorage{}
	rolemenuCollection := mdb.GetCollection("rolemenu")
	rolemenuStorage.SetCollection(rolemenuCollection)
	infrarepos.RoleMeuRepository = rolemenuinfra.NewRepository(rolemenuStorage)

	//-------------------------------------------------------
	// MenuActionRepository
	menuactionStorage :=  &storage.MongoStorage{}
	menuactionCollection := mdb.GetCollection("menuaction")
	menuactionStorage.SetCollection(menuactionCollection)
	infrarepos.MenuActionRepository = menuactioninfra.NewRepository(menuactionStorage)

	//-------------------------------------------------------
	// MenuActionResourceRepository
	menuactionresourceStorage :=  &storage.MongoStorage{}
	menuactionresourceCollection := mdb.GetCollection("menuactionresource")
	menuactionresourceStorage.SetCollection(menuactionresourceCollection)
	infrarepos.MenuActionResourceRepository = menuactionresourceinfra.NewRepository(menuactionresourceStorage, infrarepos.MenuActionRepository)


	//-------------------------------------------------------
	// UserRepository
	userStorage := &storage.MongoStorage{}
	userCollection := mdb.GetCollection("user")
	userStorage.SetCollection(userCollection)
	infrarepos.UserRepository = userinfra.NewRepository(userStorage)


	//-------------------------------------------------------
	// UserRoleRepository
	userroleStorage := &storage.MongoStorage{}
	userrroleCollection := mdb.GetCollection("userrole")
	userroleStorage.SetCollection(userrroleCollection)
	infrarepos.UserRoleRepository = userroleinfra.NewRepository(userroleStorage)


	//-------------------------------------------------------
	// AuthRepository
	authRedisStorage := &authinfra.AuthStorage{}
	// authRedisStorage.SetDatabaseService(redisdbService)
	// authRedisStorage.SetKeyPrefix("auth_")
	infrarepos.AuthRepository = authinfra.NewRepository(authRedisStorage)

	//-------------------------------------------------------
	// RbacRepository
	infrarepos.RbacRepository = rbacinfra.NewRepository(
			infrarepos.RoleRepository, 
			infrarepos.RoleMeuRepository, 
			infrarepos.MenuActionResourceRepository, 
			infrarepos.UserRepository, 
			infrarepos.UserRoleRepository)



	//-------------------------------------------------------
	// weatherRepository
	// TODO: This should be simplified just to:
	// weatherStorage := &mongoinfra.MongoStorage{}
	// weatherStorage.SetCollection("weather")

	weatherStorage := &storage.MongoStorage{}
	weatherCollection := mdb.GetCollection("weather")
	weatherStorage.SetCollection(weatherCollection)
	infrarepos.WeatherRepository = weatherinfra.NewRepository(weatherStorage)


	return infrarepos, nil
}























	






