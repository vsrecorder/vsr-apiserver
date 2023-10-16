package controllers

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vsrecorder/vsr-apiserver/pkg/controllers/helpers"
	"github.com/vsrecorder/vsr-apiserver/pkg/services"
)

var (
	ErrInvalidParameter      = errors.New("invalid parameter")
	ErrOfficialEventNotFound = errors.New("official event not found")
)

type OfficialEventController struct {
	router  *gin.Engine
	service services.OfficialEventServiceInterface
}

func NewOfficialEventController(
	router *gin.Engine,
	service services.OfficialEventServiceInterface,
) *OfficialEventController {
	return &OfficialEventController{router, service}
}

func (c *OfficialEventController) RegisterRoutes(relativePath string) {
	r := c.router.Group(relativePath + "/official_events")
	r.GET("", c.Get)
	r.GET("/:id", c.GetById)
}

func (c *OfficialEventController) Get(ctx *gin.Context) {
	if helpers.GetStartDate(ctx) != "" || helpers.GetEndDate(ctx) != "" {
		var layout = "2006-01-02"

		startDate, err := time.Parse(layout, helpers.GetStartDate(ctx))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": ErrInvalidParameter.Error(),
			})
			return
		}

		endDate, err := time.Parse(layout, helpers.GetEndDate(ctx))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": ErrInvalidParameter.Error(),
			})
			return
		}

		// startDate > endDate
		if !startDate.Before(endDate) && !startDate.Equal(endDate) {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": ErrInvalidParameter.Error(),
			})
			return
		}

		ret, err := c.service.FindByDate(ctx, startDate, endDate)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": ErrOfficialEventNotFound.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, ret)
	} else {
		if helpers.GetPage(ctx) == "" {
			page := 1
			limit := 20
			offset := limit * (page - 1)
			ret, err := c.service.Find(ctx, limit, offset)
			if err != nil {
				ctx.JSON(http.StatusNotFound, gin.H{
					"message": ErrOfficialEventNotFound.Error(),
				})
				return
			}

			ctx.JSON(http.StatusOK, ret)
		} else {
			// 取得したパラメータが数値か否か
			page, err := strconv.Atoi(helpers.GetPage(ctx))
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"message": ErrInvalidParameter.Error(),
				})
				return
			}

			// 取得したパラメータが負の数値ではないか
			if page <= 0 {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"message": ErrInvalidParameter.Error(),
				})
				return
			}

			limit := 20
			offset := limit * (page - 1)
			ret, err := c.service.Find(ctx, limit, offset)
			if err != nil {
				ctx.JSON(http.StatusNotFound, gin.H{
					"message": ErrOfficialEventNotFound.Error(),
				})
				return
			}

			ctx.JSON(http.StatusOK, ret)
		}
	}
}

func (c *OfficialEventController) GetById(ctx *gin.Context) {

	// 取得したパラメータが数値か否か
	tmpId, err := strconv.Atoi(helpers.GetId(ctx))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": ErrInvalidParameter.Error(),
		})
		return
	}

	// 取得したパラメータが負の数値ではないか
	if tmpId <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": ErrInvalidParameter.Error(),
		})
		return

	}

	id := uint(tmpId)
	ret, err := c.service.FindById(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": ErrOfficialEventNotFound.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, ret)
}
