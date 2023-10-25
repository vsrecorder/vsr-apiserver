package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vsrecorder/vsr-apiserver/pkg/controllers/helpers"
	"github.com/vsrecorder/vsr-apiserver/pkg/controllers/middlewares"
	"github.com/vsrecorder/vsr-apiserver/pkg/services"
)

const (
	USERS_PATH = "/users"
)

type UserController struct {
	router  *gin.Engine
	service services.UserServiceInterface
}

func NewUserController(
	router *gin.Engine,
	service services.UserServiceInterface,
) *UserController {
	return &UserController{router, service}
}

func (c *UserController) RegisterRoutes(relativePath string) {
	{
		r := c.router.Group(relativePath + USERS_PATH)
		r.GET("/:id", c.GetById)
		r.GET("/:id"+RECORDS_PATH, c.GetRecordsById)
		r.GET("/:id"+GAMES_PATH, c.GetGamesById)
	}

	{
		r := c.router.Group(relativePath + USERS_PATH)
		r.Use(middlewares.OptionalAuthorization)
		r.GET("/:id"+DECKS_PATH, c.GetDecksById)
	}
}

func (c *UserController) GetById(ctx *gin.Context) {
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

func (c *UserController) GetRecordsById(ctx *gin.Context) {
	id := helpers.GetId(ctx)

	ret, err := c.service.FindRecordsById(ctx, id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, ret)
}

func (c *UserController) GetGamesById(ctx *gin.Context) {
	id := helpers.GetId(ctx)

	ret, err := c.service.FindGamesById(ctx, id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, ret)
}

func (c *UserController) GetDecksById(ctx *gin.Context) {
	id := helpers.GetId(ctx)
	uid, _ := helpers.GetUID(ctx)

	ret, err := c.service.FindDecksByIdWithUID(ctx, id, uid)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, ret)
}
