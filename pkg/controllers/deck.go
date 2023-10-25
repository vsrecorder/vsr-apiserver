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
	DECKS_PATH = "/decks"
)

type DeckController struct {
	router  *gin.Engine
	service services.DeckServiceInterface
}

func NewDeckController(
	router *gin.Engine,
	service services.DeckServiceInterface,
) *DeckController {
	return &DeckController{router, service}
}

func (c *DeckController) RegisterRoutes(relativePath string) {
	{
		r := c.router.Group(relativePath + DECKS_PATH)
		r.Use(middlewares.RequiredAuthorization)
		r.GET("", c.Get)
		r.POST("", c.Create)
		r.PUT("/:id", c.Update)
		r.DELETE("/:id", c.Delete)
	}

	{
		r := c.router.Group(relativePath + DECKS_PATH)
		r.Use(middlewares.OptionalAuthorization)
		r.GET("/:id", c.GetById)
	}

	{
		r := c.router.Group(relativePath + DECKS_PATH)
		r.GET("/:id"+RECORDS_PATH, c.GetRecordById)
	}
}

func (c *DeckController) Get(ctx *gin.Context) {
	uid, _ := helpers.GetUID(ctx)

	page, err := ParsePage(ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	offset := PAGE_LIMIT * (page - 1)

	ret, err := c.service.FindByUID(ctx, uid, PAGE_LIMIT, offset)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": RecordNotFound.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"page":   page,
		"limit":  PAGE_LIMIT,
		"offset": offset,
		"decks":  ret,
	})

}

func (c *DeckController) GetById(ctx *gin.Context) {
	id := helpers.GetId(ctx)
	uid, _ := helpers.GetUID(ctx)

	ret, err := c.service.FindByIdWithUID(ctx, id, uid)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, ret)
}

func (c *DeckController) GetRecordById(ctx *gin.Context) {
	id := helpers.GetId(ctx)

	ret, err := c.service.FindRecordById(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, ret)
}

func (c *DeckController) Create(ctx *gin.Context) {
	uid, _ := helpers.GetUID(ctx)
	dto := dtos.Deck{}
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

func (c *DeckController) Update(ctx *gin.Context) {
	id := helpers.GetId(ctx)
	uid, _ := helpers.GetUID(ctx)
	dto := dtos.Deck{}
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

func (c *DeckController) Delete(ctx *gin.Context) {
	id := helpers.GetId(ctx)
	uid, _ := helpers.GetUID(ctx)

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
