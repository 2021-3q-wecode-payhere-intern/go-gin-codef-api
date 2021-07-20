package service

import (
	"encoding/json"
	"go-gin-codef-api/src/codef"
)

type BusinessData struct {
	Data   []CodefData
	Result CodefResult
}

type CodefData struct {
	Code                   string `json:"code"`
	Message                string `json:"message"`
	ExtraMessage           string `json:"extraMessage"`
	ResBusinessStatus      string `json:"resBusinessStatus"`
	ResClosingDate         string `json:"resClosingDate"`
	ResCompanyIdentityNo   string `json:"resCompanyIdentityNo"`
	ResTaxationTypeCode    string `json:"resTaxationTypeCode"`
	ResTransferTaxTypeDate string `json:"resTransferTaxTypeDate"`
}

type CodefResult struct {
	Code          string `json:"code"`
	ExtraMessage  string `json:"extraMessage"`
	Message       string `json:"message"`
	TransactionId string `json:"transactionId"`
}

func GetBusinessStatus(endPoint string, serviceType int, param map[string]interface{}) (BusinessData, error) {
	codefInstance := codef.GetCodefInstance()
	sType := codef.GetServiceType(serviceType)

	codefResult, err := codefInstance.RequestProduct(endPoint, sType, param)
	if err != nil {
		return BusinessData{}, err
	}

	var businessData BusinessData
	json.Unmarshal([]byte(codefResult), &businessData)

	return businessData, nil
}
