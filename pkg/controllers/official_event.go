package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vsrecorder/vsr-apiserver/pkg/controllers/helpers"
	"github.com/vsrecorder/vsr-apiserver/pkg/services"
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
	r.GET("/:id"+RECORDS_PATH, c.GetRecordById)
}

func (c *OfficialEventController) Get(ctx *gin.Context) {
	if helpers.GetStartDate(ctx) != "" || helpers.GetEndDate(ctx) != "" {
		startDate, endDate, err := ParseDate(ctx)
		if err != nil {
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
		return
	} else {
		page, err := ParsePage(ctx)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
			})
			return
		}

		offset := PAGE_LIMIT * (page - 1)

		ret, err := c.service.Find(ctx, PAGE_LIMIT, offset)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": ErrOfficialEventNotFound.Error(),
			})
			return
		}

		//ctx.JSON(http.StatusOK, ret)
		ctx.JSON(http.StatusOK, gin.H{
			"page":            page,
			"limit":           PAGE_LIMIT,
			"offset":          offset,
			"official_events": ret,
		})

		return
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

func (c *OfficialEventController) GetRecordById(ctx *gin.Context) {
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
	ret, err := c.service.FindRecordById(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": ErrOfficialEventNotFound.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, ret)
}
