package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vsrecorder/vsr-apiserver/pkg/controllers/dtos"
	"github.com/vsrecorder/vsr-apiserver/pkg/controllers/helpers"
	"github.com/vsrecorder/vsr-apiserver/pkg/controllers/middlewares"
	"github.com/vsrecorder/vsr-apiserver/pkg/services"
)

const (
	RECORDS_PATH = "/records"
)

type RecordController struct {
	router  *gin.Engine
	service services.RecordServiceInterface
}

func NewRecordController(
	router *gin.Engine,
	service services.RecordServiceInterface,
) *RecordController {
	return &RecordController{router, service}
}

func (c *RecordController) RegisterRoutes(relativePath string) {
	{
		r := c.router.Group(relativePath + RECORDS_PATH)
		r.Use(middlewares.RequiredAuthorization)
		r.POST("", c.Create)
		r.PUT("/:id", c.Update)
		r.DELETE("/:id", c.Delete)
	}

	{
		r := c.router.Group(relativePath + RECORDS_PATH)
		r.Use(middlewares.OptionalAuthorization)
		r.GET("", c.Get)
	}

	{
		r := c.router.Group(relativePath + RECORDS_PATH)
		r.GET("/:id", c.GetById)
		r.GET("/:id"+GAMES_PATH, c.GetGameById)
	}
}

func (c *RecordController) Get(ctx *gin.Context) {
	uid, existsUID := helpers.GetUID(ctx)

	page, err := ParsePage(ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	offset := PAGE_LIMIT * (page - 1)

	if existsUID {
		ret, err := c.service.FindByUID(ctx, uid, PAGE_LIMIT, offset)

		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": RecordNotFound.Error(),
			})
			return
		}

		//ctx.JSON(http.StatusOK, ret)
		ctx.JSON(http.StatusOK, gin.H{
			"page":    page,
			"limit":   PAGE_LIMIT,
			"offset":  offset,
			"records": ret,
		})

		return
	} else {
		ret, err := c.service.Find(ctx, PAGE_LIMIT, offset)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": RecordNotFound.Error(),
			})
			return
		}

		//ctx.JSON(http.StatusOK, ret)
		ctx.JSON(http.StatusOK, gin.H{
			"page":    page,
			"limit":   PAGE_LIMIT,
			"offset":  offset,
			"records": ret,
		})

		return
	}
}

func (c *RecordController) GetById(ctx *gin.Context) {
	id := helpers.GetId(ctx)

	ret, err := c.service.FindById(ctx, id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, ret)
}

func (c *RecordController) GetGameById(ctx *gin.Context) {
	id := helpers.GetId(ctx)

	ret, err := c.service.FindGameById(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, ret)
}

func (c *RecordController) Create(ctx *gin.Context) {
	uid, _ := helpers.GetUID(ctx)
	dto := dtos.Record{}
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ret, err := c.service.Create(ctx, uid, &dto)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, ret)
}

func (c *RecordController) Update(ctx *gin.Context) {
	id := helpers.GetId(ctx)
	uid, _ := helpers.GetUID(ctx)
	dto := dtos.Record{}
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ret, err := c.service.Update(ctx, id, uid, &dto)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, ret)
}

func (c *RecordController) Delete(ctx *gin.Context) {
	uid, _ := helpers.GetUID(ctx)
	id := helpers.GetId(ctx)

	err := c.service.Delete(ctx, id, uid)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"message": "accepted",
	})
}
