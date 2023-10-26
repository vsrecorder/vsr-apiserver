package controllers

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vsrecorder/vsr-apiserver/pkg/controllers/helpers"
)

const (
	PAGE_LIMIT  = 20
	DATE_LAYOUT = "2006-01-02"
)

func ParsePage(ctx *gin.Context) (int, error) {
	pageNumber := helpers.GetPage(ctx)

	if pageNumber == "" {
		page := 1

		return page, nil
	} else {
		page, err := strconv.Atoi(pageNumber)

		if err != nil { // 取得したパラメータが数値か否か
			return 0, err
		} else if page <= 0 {
			return 1, nil
		}

		return page, nil
	}
}

func ParseDate(ctx *gin.Context) (time.Time, time.Time, error) {
	startDate, err := time.Parse(DATE_LAYOUT, helpers.GetStartDate(ctx))
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	endDate, err := time.Parse(DATE_LAYOUT, helpers.GetEndDate(ctx))
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	startDate = time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 0, 0, 0, 0, jst)
	endDate = time.Date(endDate.Year(), endDate.Month(), endDate.Day(), 23, 59, 59, 59, jst)

	// startDate > endDate
	if !startDate.Before(endDate) && !startDate.Equal(endDate) {
		return time.Time{}, time.Time{}, err
	}

	return startDate, endDate, nil
}
