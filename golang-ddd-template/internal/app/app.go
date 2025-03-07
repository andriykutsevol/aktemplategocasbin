package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/andriykutsevol/WeatherServer/configs"

	"github.com/andriykutsevol/WeatherServer/internal/presentation/http/middleware"
	"github.com/gin-gonic/gin"

	"github.com/andriykutsevol/WeatherServer/internal/presentation/http/handler"
	"github.com/andriykutsevol/WeatherServer/internal/presentation/http/router"

	"github.com/andriykutsevol/WeatherServer/internal/app/application"

	//services
	"github.com/andriykutsevol/WeatherServer/internal/domain/menu"
	"github.com/andriykutsevol/WeatherServer/internal/domain/user"

	"github.com/andriykutsevol/WeatherServer/internal/infrastructure"
	"github.com/andriykutsevol/WeatherServer/internal/infrastructure/mongo"
	"github.com/andriykutsevol/WeatherServer/internal/infrastructure/pg"
	pgstorage "github.com/andriykutsevol/WeatherServer/internal/infrastructure/pg/storage"
)



func Run(configPath string) {

	var wg sync.WaitGroup

	//=========================================================================
	// Config
	//=========================================================================


	configs.MustLoad("../../configs/config.toml")
	//configs.PrintWithJSON()


	//=========================================================================
	// Repo Builder
	// Dependency injection for Infrastructure layer
	// Driven Adapters (Hexagonal Architecture )
	// Repository implementation - adapter.
	//=========================================================================

	var infrarepos *infrastructure.InfraRepos

    db_type, ok := os.LookupEnv("DBTYPE")
    if !ok {
		log.Fatal("DBTYPE environment variable is not set")
    }


	if(db_type == "mongo"){
		var err error
		infrarepos, err = mongo.BuildRespositories()
		if err != nil {
			log.Fatal("Cannot mongo.BuildRespositories():", err)
		}

	}else if db_type == "pg"{

		var err error
		// https://medium.com/geekculture/work-with-go-postgresql-using-pgx-caee4573672

		// if we use docker run with --nerwork, we have to use container name (@template_go_react_pg)
		// if we use docker-compose, we have to use service name (@postgres).

		casbinDatabaseUri, ok := os.LookupEnv("PGCASBINURI")
		if !ok {
			log.Fatal("PGCASBINURI environment variable is not set")
		}
		casbinDbService, err := pgstorage.NewDatabaseService(casbinDatabaseUri)
		if err != nil {
			log.Fatal("Unable to connect to casbin database:", err)
		}
		defer casbinDbService.Pool.Close()


		weatherDatabaseUri, ok := os.LookupEnv("PGWEATHERURI")
		if !ok {
			log.Fatal("PGWEATHERURI environment variable is not set")
		}
		weatherDbService, err := pgstorage.NewDatabaseService(weatherDatabaseUri)
		if err != nil {
			log.Fatal("Unable to connect to weather database:", err)
		}
		defer casbinDbService.Pool.Close()
		

		infrarepos, err = pg.BuildRespositories(casbinDbService, weatherDbService)
		if err != nil {
			log.Fatal("Cannot pg.BuildRespositories():", err)
		}

	} else {
		log.Fatal("Error: db_type is wrong. It have to be either 'mongo' or 'pg'. ")
	}


	//=========================================================================
	// Driving Adapters  (Hexagonal Architecture )
	//=========================================================================
	
	//-----------------------------------------------------------
	// Dependency injection for Domain Services
	//-----------------------------------------------------------

	userservice := user.NewService(
		infrarepos.AuthRepository, 
		infrarepos.UserRepository, 
		infrarepos.RoleRepository, 
		infrarepos.UserRoleRepository)
		
	menuService := menu.NewService(
		infrarepos.MenuRepository, 
		infrarepos.MenuActionRepository, 
		infrarepos.MenuActionResourceRepository)


	//-----------------------------------------------------------
	//Dependency injection for Application layer
	//-----------------------------------------------------------

	// In Domain-Driven Design (DDD), the application layer is typically not considered an adapter.
	// Instead, it serves as a distinct layer responsible for coordinating interactions between the domain model and external systems or user interfaces.

	// TODO: Whether to launch or not should be decided by the config file.
	seedApplication := application.NewSeed(
		menuService,
		infrarepos.MenuActionRepository,
		infrarepos.RoleRepository, 
		infrarepos.RoleMeuRepository, 
		infrarepos.UserRepository, 
		infrarepos.UserRoleRepository, 
		infrarepos.WeatherRepository,
	)
	seedApplication.Execute(context.TODO())

	rbacAdapter := application.NewRbacAdapter(infrarepos.RbacRepository)

	syncedEnforcer, cleanup3, err := InitCasbin(rbacAdapter)
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}	
	_ = cleanup3

	roleApplication := application.NewRole(
			rbacAdapter, 
			syncedEnforcer, 
			infrarepos.RoleRepository, 
			infrarepos.RoleMeuRepository, 
			infrarepos.UserRepository)

	userApplication := application.NewUser(
			infrarepos.AuthRepository, 
			rbacAdapter, syncedEnforcer, 
			infrarepos.UserRepository, 
			infrarepos.UserRoleRepository, 
			infrarepos.RoleRepository)

	loginApplication := application.NewLogin(
			infrarepos.AuthRepository, 
			infrarepos.UserRepository, 
			infrarepos.RoleRepository, 
			infrarepos.UserRoleRepository, 
			userservice, 
			infrarepos.MenuRepository, 
			infrarepos.MenuActionRepository, 
			infrarepos.RoleMeuRepository)	

	menuApplication := application.NewMenu(menuService)


	demosApplication := application.NewDemos(
		infrarepos.UserRepository, 
		infrarepos.UserRoleRepository)

	weatherApplication := application.NewWeather(infrarepos.WeatherRepository)

	//-----------------------------------------------------------
	// Dependency injection for Presentation layer
	//-----------------------------------------------------------


	roleHandler := handler.NewRole(roleApplication)
	userHandler := handler.NewUser(userApplication)
	menuHandler := handler.NewMenu(menuApplication)
	loginHandler := handler.NewLogin(loginApplication)


	demosHandler := handler.NewDemos(demosApplication)

	weatherHandler := handler.NewWeather(weatherApplication)

	routerRouter := router.NewRouter(
		infrarepos.AuthRepository, 
		syncedEnforcer, 
		loginHandler, 
		menuHandler, 
		roleHandler, 
		userHandler, 
		demosHandler, 
		weatherHandler,
	)



	//=========================================================================
	// Init Gin Engine
	//=========================================================================

	//gin.SetMode("debug")
	//gin.SetMode(gin.ReleaseMode)
	ginEngine := gin.New()

	// CORS
	if configs.C.CORS.Enable {
		ginEngine.Use(middleware.CORSMiddleware())
	}	

	//--------------------------------------------------------------
	ginEngine.GET("/swagger.yaml", func(c *gin.Context) {
        c.File("../../internal/presentation/swagger/swagger.yaml")
    })
	ginEngine.Static("/swaggerui/", "../../swaggerui")
	//--------------------------------------------------------------


	routerRouter.Register(ginEngine)

	srv := &http.Server{
		Addr:         "0.0.0.0:8080",
		Handler:      ginEngine,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Println("Starting HTTP server on port 8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Println("Error starting server:", err)
		}

	}()

	wg.Wait()
	log.Println("Server stopped")

}