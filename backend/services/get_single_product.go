package services

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary api untuk mendapatkan data product satuan
// @Tags 	product
// @Param 	id	path	int	true	"Product ID"
// @Success	200	{object}	model.Product
// @Router	/product/{id}	[get]
func (ser *Service) GetSingleProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ser.sendError(ctx, http.StatusBadRequest, "invalid product id")
		return
	}

	result, err := ser.db.FetchSingleProduct(idInt)
	if err != nil {
		ser.sendError(ctx, http.StatusNotFound, "data not found")
		return
	}

	ser.sendJson(ctx, http.StatusOK, result)
}
