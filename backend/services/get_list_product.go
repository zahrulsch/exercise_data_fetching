package services

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetListProductQuery struct {
	Page int `form:"page" default:"1"`
	Size int `form:"size" default:"10"`
}

func (query *GetListProductQuery) queryGuard() *GetListProductQuery {
	if query.Page < 1 {
		query.Page = 1
	}

	if query.Size < 1 {
		query.Size = 10
	}

	return query
}

// @Summary api untuk mendapatkan data list product
// @Tags	product
// @Success	200	{object}	database.FetchProductResult
// @Param	default	query	GetListProductQuery	false	"GetListProductQuery"
// @Router	/product	[get]
func (service *Service) GetListProduct(ctx *gin.Context) {
	query := GetListProductQuery{Page: 1, Size: 10}
	if err := ctx.BindQuery(&query); err != nil {
		service.sendError(ctx, http.StatusBadRequest, "invalid query params")
		return
	}

	query = *query.queryGuard()
	fmt.Printf("%+v\n", query)

	data, err := service.db.FetchProduct(query.Size, (query.Page-1)*query.Size)
	if err != nil {
		service.sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	service.sendJson(ctx, http.StatusOK, data)
}
