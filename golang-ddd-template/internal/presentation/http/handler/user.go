package handler 

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/andriykutsevol/WeatherServer/internal/app/application"

	"github.com/andriykutsevol/WeatherServer/internal/domain/errors"
	"github.com/andriykutsevol/WeatherServer/internal/domain/user"
	"github.com/andriykutsevol/WeatherServer/internal/domain/pagination"

	"github.com/andriykutsevol/WeatherServer/internal/presentation/http"

	"github.com/andriykutsevol/WeatherServer/internal/presentation/http/response"
	"github.com/andriykutsevol/WeatherServer/internal/presentation/http/request"

)



type User interface {
	Query(c *gin.Context)
	Get(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	Enable(c *gin.Context)
	Disable(c *gin.Context)
}

type userHandler struct {
	userApp application.User
}


func NewUser(userApp application.User) User {
	return &userHandler{userApp: userApp}
}


func (a *userHandler) Create(c *gin.Context) {
	fmt.Println("User Handler Create()")

	ctx := c.Request.Context()
	var item request.User

	if err := http.ParseJSON(c, &item); err != nil {
		http.ResError(c, err)
		return
	} else if item.Password == "" {
		http.ResError(c, errors.New400Response("password is empty"))
		return
	}

	item.Creator = http.GetUserID(c)


	result, err := a.userApp.Create(ctx, item.ToDomain(), item.RoleIDs)

	if err != nil {
		http.ResError(c, err)
		return
	}
	http.ResSuccess(c, response.NewIDResult(result))	

}



func (a *userHandler) Query(c *gin.Context) {



	ctx := c.Request.Context()
	var params request.UserQueryParam
	if err := http.ParseQuery(c, &params); err != nil {
		http.ResError(c, err)
		return
	}
	if v := c.Query("roleIDs"); v != "" {
		params.RoleIDs = strings.Split(v, ",")
	}

	domainParams := user.QueryParams{
		PaginationParam: pagination.Param{Pagination: true},
		OrderFields:     nil,
		UserName:        params.UserName,
		QueryValue:      params.QueryValue,
		Status:          params.Status,
		RoleIDs:         params.RoleIDs,
	}

	result, p, err := a.userApp.QueryShow(ctx, domainParams)
	if err != nil {
		http.ResError(c, err)
		return
	}
	http.ResPage(c, response.UsersFromDomain(result), p)	


}


func (a *userHandler) Get(c *gin.Context) {}




func (a *userHandler) Update(c *gin.Context) {


}


func (a *userHandler) Delete(c *gin.Context) {

}


func (a *userHandler) Enable(c *gin.Context) {}



func (a *userHandler) Disable(c *gin.Context) {}
















