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
	GAMES_PATH = "/games"
)

type GameController struct {
	router  *gin.Engine
	service services.GameServiceInterface
}

func NewGameController(
	router *gin.Engine,
	service services.GameServiceInterface,
) *GameController {
	return &GameController{router, service}
}

func (c *GameController) RegisterRoutes(relativePath string) {
	{
		r := c.router.Group(relativePath + GAMES_PATH)
		r.Use(middlewares.RequiredAuthorization)
		r.POST("", c.Create)
		r.PUT("/:id", c.Update)
		r.DELETE("/:id", c.Delete)
	}

	{
		r := c.router.Group(relativePath + GAMES_PATH)
		r.GET("/:id", c.GetById)
	}
}

func (c *GameController) GetById(ctx *gin.Context) {
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

func (c *GameController) Create(ctx *gin.Context) {
	uid, _ := helpers.GetUID(ctx)
	dto := dtos.Game{}
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

func (c *GameController) Update(ctx *gin.Context) {
	id := helpers.GetId(ctx)
	uid, _ := helpers.GetUID(ctx)
	dto := dtos.Game{}
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ret, err := c.service.Update(ctx, id, uid, &dto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, ret)
}

func (c *GameController) Delete(ctx *gin.Context) {
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
