package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/andriykutsevol/WeatherServer/internal/app/application"

	"github.com/andriykutsevol/WeatherServer/internal/domain/pagination"
	"github.com/andriykutsevol/WeatherServer/internal/domain/user/role"

	"github.com/andriykutsevol/WeatherServer/internal/presentation/http/request"
	"github.com/andriykutsevol/WeatherServer/internal/presentation/http/response"
	"github.com/andriykutsevol/WeatherServer/internal/presentation/http"

)



type Role interface {
	Query(c *gin.Context)
	QuerySelect(c *gin.Context)
	Get(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	Enable(c *gin.Context)
	Disable(c *gin.Context)
}


type roleHandler struct {
	roleApp application.Role
}


func NewRole(roleApp application.Role) Role {
	return &roleHandler{roleApp: roleApp}
}



func (a *roleHandler) Create(c *gin.Context) {

	ctx := c.Request.Context()
	var item request.Role

	if err := http.ParseJSON(c, &item); err != nil {
		http.ResError(c, err)
		return
	}

	// TODO. User Login.
	item.Creator = http.GetUserID(c)

	result, err := a.roleApp.Create(ctx, item.ToDomain())
	if err != nil {
		http.ResError(c, err)
		return
	}
	http.ResSuccess(c, response.NewIDResult(result))
}


func (a *roleHandler) Query(c *gin.Context) {

	ctx := c.Request.Context()
	var params request.RoleQueryParam
	if err := http.ParseQuery(c, &params); err != nil {
		http.ResError(c, err)
		return
	}

	domainParams := role.QueryParam{
		PaginationParam: pagination.Param{Pagination: true},
		OrderFields:     pagination.NewOrderFields(pagination.NewOrderField("sequence", pagination.OrderByDESC)),
		IDs:             params.IDs,
		Name:            params.Name,
		QueryValue:      params.QueryValue,
		UserID:          params.UserID,
		Status:          params.Status,
	}

	result, p, err := a.roleApp.Query(ctx, domainParams)

	if err != nil {
		http.ResError(c, err)
		return
	}

	http.ResPage(c, response.RolesFromDomain(result), p)	

}



func (a *roleHandler) Get(c *gin.Context) {


}


func (a *roleHandler) Delete(c *gin.Context) {

}



func (a *roleHandler) QuerySelect(c *gin.Context) {

}


func (a *roleHandler) Update(c *gin.Context) {

}


func (a *roleHandler) Enable(c *gin.Context) {

}


func (a *roleHandler) Disable(c *gin.Context) {


}