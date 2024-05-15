package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joidiego/PASTI_KEl04/dto"
	"github.com/joidiego/PASTI_KEl04/entity"
	"github.com/joidiego/PASTI_KEl04/helper"
	"github.com/joidiego/PASTI_KEl04/service"
)

// SliderController is a contract about something that this controller can do
type SliderController interface {
	All(ctx *gin.Context)
	FindByID(ctx *gin.Context)
	Insert(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type sliderController struct {
	sliderService service.SliderService
}

// NewSliderController creates a new instance of SliderController
func NewSliderController(sliderService service.SliderService) SliderController {
	return &sliderController{
		sliderService: sliderService,
	}
}

func (c *sliderController) All(ctx *gin.Context) {
	var sliders []entity.Slider = c.sliderService.All()
	if len(sliders) == 0 {
		res := helper.BuildErrorResponse("No data found", "No sliders available", helper.EmptyObj{})
		ctx.JSON(http.StatusNotFound, res)
		return
	}
	res := helper.BuildResponse(true, "OK!", sliders)
	ctx.JSON(http.StatusOK, res)
}

func (c *sliderController) FindByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to get ID", "No param ID were found", helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	slider := c.sliderService.FindByID(id)
	if (slider == entity.Slider{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given ID", helper.EmptyObj{})
		ctx.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK!", slider)
		ctx.JSON(http.StatusOK, res)
	}
}

func (c *sliderController) Insert(ctx *gin.Context) {
	// Check if the count is less than 3 before inserting
	if c.sliderService.Count() >= 3 {
		res := helper.BuildErrorResponse("Failed to process request", "Maximum number of sliders reached", helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	var sliderCreateDTO dto.SliderCreateDTO
	errDTO := ctx.ShouldBind(&sliderCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	result := c.sliderService.Insert(sliderCreateDTO)
	response := helper.BuildResponse(true, "OK!", result)
	ctx.JSON(http.StatusCreated, response)
}

func (c *sliderController) Update(ctx *gin.Context) {
	var sliderUpdateDTO dto.SliderUpdateDTO
	errDTO := ctx.ShouldBind(&sliderUpdateDTO)
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
	sliderUpdateDTO.ID = id
	result := c.sliderService.Update(sliderUpdateDTO)
	response := helper.BuildResponse(true, "OK!", result)
	ctx.JSON(http.StatusOK, response)
}

func (c *sliderController) Delete(ctx *gin.Context) {
	var slider entity.Slider
	id, errID := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if errID != nil {
		res := helper.BuildErrorResponse("Failed to get ID", "No param ID were found", helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	slider.ID = id
	c.sliderService.Delete(slider)
	res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}
