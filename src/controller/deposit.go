package controller

import (
	"go-gin-codef-api/src/message"
	"go-gin-codef-api/src/properties"
	"go-gin-codef-api/src/service"
	"go-gin-codef-api/src/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 일자별 입금내역
func GetDepositByDaily(c *gin.Context) {
	year := c.Query("year")
	month := c.Query("month")

	if utils.IsEmptyByParam(year) || utils.IsEmptyByParam(month) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": message.RESPONSE_INVALID_REQUEST,
		})
		return
	}

	lastDay := utils.GetLastDay(year, month)
	startDate := utils.GetConcatDate(year, month, properties.MONTH_BY_FIRST_DAY)
	endDate := utils.GetConcatDate(year, month, lastDay)

	resultMap := service.GetDepositByDailySum(startDate, endDate)

	c.JSON(http.StatusOK, gin.H{
		"message": message.SUCCESS,
		"data":    resultMap,
	})
}

// 카드사별 입금내역 - 일자기준
func GetDepositByCard(c *gin.Context) {
	date := c.Param("date")

	if utils.IsEmptyByParam(date) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": message.RESPONSE_INVALID_REQUEST,
		})
		return
	}

	resultMap := service.GetDepositByCardSum(date)

	c.JSON(http.StatusOK, gin.H{
		"message": message.SUCCESS,
		"data":    resultMap,
	})
}
