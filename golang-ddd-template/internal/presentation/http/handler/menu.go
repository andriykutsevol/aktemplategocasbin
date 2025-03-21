package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/andriykutsevol/WeatherServer/internal/app/application"
	"github.com/andriykutsevol/WeatherServer/internal/domain/menu"
	"github.com/andriykutsevol/WeatherServer/internal/domain/pagination"
	"github.com/andriykutsevol/WeatherServer/internal/presentation/http"
	"github.com/andriykutsevol/WeatherServer/internal/presentation/http/request"
	"github.com/andriykutsevol/WeatherServer/internal/presentation/http/response"
)

type Menu interface {
	Query(c *gin.Context)
	QueryTree(c *gin.Context)
	Get(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	Enable(c *gin.Context)
	Disable(c *gin.Context)
}

func NewMenu(menuApp application.Menu) Menu {
	return &menuHandler{
		menuApp: menuApp,
	}
}

type menuHandler struct {
	menuApp application.Menu
}

func (a *menuHandler) Query(c *gin.Context) {
	ctx := c.Request.Context()
	var params request.MenuQueryParam
	if err := http.ParseQuery(c, &params); err != nil {
		http.ResError(c, err)
		return
	}

	domainParams := menu.QueryParam{
		IDs:              params.IDs,
		Name:             params.Name,
		PrefixParentPath: params.PrefixParentPath,
		QueryValue:       params.QueryValue,
		ParentID:         params.ParentID,
		ShowStatus:       params.ShowStatus,
		Status:           params.Status,
		PaginationParam: pagination.Param{
			Pagination: true,
		},
		OrderFields: pagination.NewOrderFields(pagination.NewOrderField("sequence", pagination.OrderByDESC)),
	}
	result, p, err := a.menuApp.Query(ctx, domainParams)
	if err != nil {
		http.ResError(c, err)
		return
	}
	http.ResPage(
		c,
		response.MenusFromDomain(result),
		p,
	)
}

func (a *menuHandler) QueryTree(c *gin.Context) {
	ctx := c.Request.Context()
	var params request.MenuQueryParam
	if err := http.ParseQuery(c, &params); err != nil {
		http.ResError(c, err)
		return
	}
	domainParams := menu.QueryParam{
		IDs:              params.IDs,
		Name:             params.Name,
		PrefixParentPath: params.PrefixParentPath,
		QueryValue:       params.QueryValue,
		ParentID:         params.ParentID,
		ShowStatus:       params.ShowStatus,
		Status:           params.Status,
		PaginationParam: pagination.Param{
			Pagination: true,
		},
		OrderFields: pagination.NewOrderFields(pagination.NewOrderField("sequence", pagination.OrderByDESC)),
	}

	result, _, err := a.menuApp.Query(ctx, domainParams)
	if err != nil {
		http.ResError(c, err)
		return
	}
	http.ResList(c, response.MenusFromDomain(result).ToTree())
}

func (a *menuHandler) Get(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.menuApp.Get(ctx, c.Param("id"))
	if err != nil {
		http.ResError(c, err)
		return
	}
	http.ResSuccess(c, response.MenuFromDomain(item))
}




func (a *menuHandler) Create(c *gin.Context) {

	ctx := c.Request.Context()
	var item response.Menu
	if err := http.ParseJSON(c, &item); err != nil {
		http.ResError(c, err)
		return
	}

	item.Creator = http.GetUserID(c)
	
	result, err := a.menuApp.Create(ctx, item.ToDomain())

	if err != nil {
		http.ResError(c, err)
		return
	}

	http.ResSuccess(c, response.NewIDResult(result))
}



func (a *menuHandler) Update(c *gin.Context) {
	ctx := c.Request.Context()
	var item response.Menu
	if err := http.ParseJSON(c, &item); err != nil {
		http.ResError(c, err)
		return
	}

	err := a.menuApp.Update(ctx, c.Param("id"), item.ToDomain())
	if err != nil {
		http.ResError(c, err)
		return
	}
	http.ResOK(c)
}

func (a *menuHandler) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.menuApp.Delete(ctx, c.Param("id"))
	if err != nil {
		http.ResError(c, err)
		return
	}
	http.ResOK(c)
}

func (a *menuHandler) Enable(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.menuApp.UpdateStatus(ctx, c.Param("id"), 1)
	if err != nil {
		http.ResError(c, err)
		return
	}
	http.ResOK(c)
}

func (a *menuHandler) Disable(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.menuApp.UpdateStatus(ctx, c.Param("id"), 2)
	if err != nil {
		http.ResError(c, err)
		return
	}
	http.ResOK(c)
}
