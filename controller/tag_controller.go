package controller

import (
	"backendgo/data/request"
	"backendgo/data/response"
	"backendgo/helper"
	"backendgo/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TagsController struct {
	TagsService service.TagsService
}

func NewTagsController(service service.TagsService) *TagsController {
	return &TagsController{
		TagsService: service,
	}
}

// CreateTags godoc
// @Summary      Create tag
// @Description  Save tag data in DB
// @Tags		 tags
// @Param        tags body request.CreateTagsRequest true "Create tags"
// @Produce      application/json
// @Success      200  {object}  response.Response{}
// @Router       /tags [post]
func (controller *TagsController) Create(ctx *gin.Context) {
	CreateTagsRequest := request.CreateTagsRequest{}
	err := ctx.ShouldBindJSON(&CreateTagsRequest)
	helper.ErrorPanic(err)

	controller.TagsService.Create(CreateTagsRequest)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	}
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// UpdateTags godoc
// @Summary      Update tag
// @Description  Update tags data
// @Tags		 tags
// @Param        tagId path string true "update tags by id"
// @Param        tags body request.UpdateTagsRequest true "Update tags"
// @Produce      application/json
// @Success      200  {object}  response.Response{}
// @Router       /tags/{tagId} [patch]
func (controller *TagsController) Update(ctx *gin.Context) {
	updateTagRequest := request.UpdateTagsRequest{}
	err := ctx.ShouldBindJSON(&updateTagRequest)
	helper.ErrorPanic(err)
	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)
	updateTagRequest.Id = id
	controller.TagsService.Update(updateTagRequest)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	}
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// DeleteTags godoc
// @Summary      Delete tag
// @Description  Delete tag data by id from db
// @Tags		 tags
// @Param        tagId path string true "delete tags by id"
// @Produce      application/json
// @Success      200  {object}  response.Response{}
// @Router       /tags/{tagId} [delete]
func (controller *TagsController) Delete(ctx *gin.Context) {
	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)
	controller.TagsService.Delete(id)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	}
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}

// FindByIdTags godoc
// @Summary      Get Single Tag by id
// @Description  Get tag single data by id if it mathes from db
// @Tags		 tags
// @Param        tagId path string true "find tag by id"
// @Produce      application/json
// @Success      200  {object}  response.Response{}
// @Router       /tags/{tagId} [get]
func (controller *TagsController) FindById(ctx *gin.Context) {
	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)
	tagResponse := controller.TagsService.FindById(id)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   tagResponse,
	}
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// FindByIdTags godoc
// @Summary      Get all Tags
// @Description  Return list of tags data
// @Tags		 tags
// @Success      200  {object}  response.Response{}
// @Router       /tags [get]
func (controller *TagsController) FindAll(ctx *gin.Context) {
	tagResponse := controller.TagsService.FindAll()

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   tagResponse,
	}

	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
