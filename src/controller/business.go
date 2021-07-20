package controller

import (
	"go-gin-codef-api/src/message"
	"go-gin-codef-api/src/properties"
	"go-gin-codef-api/src/service"
	"go-gin-codef-api/src/utils"
	"net/http"
	"runtime"
	"strings"

	"github.com/gin-gonic/gin"
)

// 휴폐업 상태 조회
func GetBusinessStatus(c *gin.Context) {
	bizNumber := c.Query("biz_number")
	bizNumber = strings.ReplaceAll(bizNumber, "-", "")

	if utils.IsEmptyByParam(bizNumber) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": message.RESPONSE_INVALID_REQUEST,
			"data":    map[string]string{},
		})
		return
	}

	reqData := []map[string]string{
		{"reqIdentity": bizNumber},
	}

	parameter := map[string]interface{}{
		"organization":    properties.ORGANIZATION_CODE,
		"reqIdentityList": reqData,
	}

	codefResult, err := service.GetBusinessStatus(properties.BUSINESS_STATUS_END_POINT, properties.TYPE_DEMO, parameter)
	if err != nil {
		_, fileName, line, _ := runtime.Caller(0)
		msg := utils.GenerateErrorMsg(fileName, line)

		c.JSON(http.StatusServiceUnavailable, gin.H{
			"message": msg,
			"data":    map[string]string{},
		})
		return
	}

	if codefResult.Result.Code != properties.CODEF_SUCCESS_CODE {
		errorMsg := utils.GenerateCodefErrorMsg(codefResult.Result.Code + codefResult.Result.Message)

		c.JSON(http.StatusServiceUnavailable, gin.H{
			"message": errorMsg,
			"data":    map[string]string{},
		})
		return
	}

	if len(codefResult.Data) <= 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": message.RESPONSE_RESULT_EMPTY,
			"data":    map[string]string{},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": message.SUCCESS,
		"data":    codefResult.Data[0],
	})
}
