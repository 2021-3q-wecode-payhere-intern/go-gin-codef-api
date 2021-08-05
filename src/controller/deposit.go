package controller

import (
	"go-gin-codef-api/src/http/request"
	"go-gin-codef-api/src/http/response"
	"go-gin-codef-api/src/properties"
	"go-gin-codef-api/src/service"
	"go-gin-codef-api/src/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Tags        Deposit API
// @Summary     일자별 입금금액 조회
// @Description 월단위 기준으로 일자별 입금금액을 조회한다.
// @Accept      json
// @Produce     json
// @Param       year query string true "연도 - ex)2021"
// @Param       month query string true "월 - ex)04"
// @Router      /deposit [get]
// @Success     200 {string} string "OK"
func GetDepositByDaily(c *gin.Context) {
	year := c.Query("year")
	month := c.Query("month")

	if request.IsEmptyByParam(year) || request.IsEmptyByParam(month) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": response.INVALID_REQUEST,
		})
		return
	}

	lastDay := utils.GetLastDay(year, month)
	startDate := utils.GetConcatDate(year, month, properties.MONTH_BY_FIRST_DAY)
	endDate := utils.GetConcatDate(year, month, lastDay)

	resultMap := service.GetDepositByDailySum(startDate, endDate)

	c.JSON(http.StatusOK, gin.H{
		"message": response.SUCCESS,
		"data":    resultMap,
	})
}

// @Tags        Deposit API
// @Summary     카드사별 입금금액 조회
// @Description 일자 기준으로 카드사별 입금금액을 조회한다.
// @Accept      json
// @Produce     json
// @Param       date path string true "일자"
// @Router      /deposit/:date [get]
// @Success     200 {string} string "OK"
func GetDepositByCard(c *gin.Context) {
	date := c.Param("date")

	if request.IsEmptyByParam(date) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": response.INVALID_REQUEST,
		})
		return
	}

	resultMap := service.GetDepositByCardSum(date)

	c.JSON(http.StatusOK, gin.H{
		"message": response.SUCCESS,
		"data":    resultMap,
	})
}
