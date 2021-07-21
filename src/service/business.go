package service

import (
	"encoding/json"
	"go-gin-codef-api/src/codef"
	"go-gin-codef-api/src/entity"
)

// 휴폐업 상태 조회
func GetBusinessStatus(endPoint string, serviceType int, param map[string]interface{}) (entity.BusinessData, error) {
	codefInstance := codef.GetCodefInstance()
	sType := codef.GetServiceType(serviceType)

	codefResult, err := codefInstance.RequestProduct(endPoint, sType, param)
	if err != nil {
		return entity.BusinessData{}, err
	}

	var businessData entity.BusinessData
	json.Unmarshal([]byte(codefResult), &businessData)

	return businessData, nil
}
