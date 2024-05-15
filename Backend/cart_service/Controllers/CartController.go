package controller

import (
	service "cart_service/Service"
	"cart_service/dto"
	"cart_service/entity"
	"cart_service/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CartController is a contract about something that this controller can do
type CartController interface {
	All(ctx *gin.Context)
	Insert(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type cartController struct {
	CartService service.CartService
}

// NewCartController creates a new instance of AuthController
func NewCartController(CartService service.CartService) CartController {
	return &cartController{
		CartService: CartService,
	}
}

func (c *cartController) All(ctx *gin.Context) {
	// get carts where body request ID == carts.user_id
	userID, err := strconv.ParseUint(ctx.Query("user_id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to get ID", "No param ID were found", helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	carts := c.CartService.All(userID)
	res := helper.BuildResponse(true, "OK!", carts)
	ctx.JSON(http.StatusOK, res)
}

func (c *cartController) Insert(ctx *gin.Context) {
	var cartCreateDTO dto.CartCreateDTO
	errDTO := ctx.ShouldBind(&cartCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	result := c.CartService.Insert(cartCreateDTO)
	response := helper.BuildResponse(true, "OK!", result)
	ctx.JSON(http.StatusCreated, response)
}

func (c *cartController) Update(ctx *gin.Context) {
	var cartUpdateDTO dto.CartUpdateDTO
	errDTO := ctx.ShouldBind(&cartUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	id, errID := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if errID != nil {
		res := helper.BuildErrorResponse("Failed to get ID", "No param ID were found", helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	cartUpdateDTO.ID = id
	result := c.CartService.Update(cartUpdateDTO)
	response := helper.BuildResponse(true, "OK!", result)
	ctx.JSON(http.StatusOK, response)
}

func (c *cartController) Delete(ctx *gin.Context) {
	var cart entity.Cart
	id, errID := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if errID != nil {
		res := helper.BuildErrorResponse("Failed to get ID", "No param ID were found", helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	cart.ID = id
	c.CartService.Delete(cart)
	res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}
