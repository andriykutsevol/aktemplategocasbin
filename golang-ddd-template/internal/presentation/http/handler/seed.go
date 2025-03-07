package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/andriykutsevol/WeatherServer/internal/presentation/http"
	"github.com/andriykutsevol/WeatherServer/internal/app/application"
)


type Seed interface {
	Execute(c *gin.Context)
}


type seedHandler struct {
	seedApp application.Seed
}


type Repsonse200 struct{
	Message string		`json:"message"` 
}


func NewSeed(seedApp application.Seed) Seed {
	return &seedHandler{
		seedApp: seedApp,
	}
}


func (a *seedHandler) Execute(c *gin.Context){

	//TODO: config
	err := a.seedApp.Execute(c)
	if err != nil {
		fmt.Println("a.seedApp.Execute error", err)
	}

	http.ResSuccess(c, Repsonse200{Message: "Ok"})
}