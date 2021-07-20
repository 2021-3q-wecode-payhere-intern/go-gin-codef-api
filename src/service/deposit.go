package service

import (
	"context"
	"go-gin-codef-api/src/db"
	"go-gin-codef-api/src/properties"
	"log"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
)

type DepositData struct {
	ResCardCompany       string `json:"resCardCompany"`
	ResDepositDate       string `json:"resDepositDate"`
	ResMemberStoreNo     string `json:"resMemberStoreNo"`
	ResOtherDeposit      string `json:"resOtherDeposit"`
	ResPaymentAccount    string `json:"resPaymentAccount"`
	ResSalesAmount       string `json:"resSalesAmount"`
	CommEndDate          string `json:"commEndDate"`
	ResAccountIn         string `json:"resAccountIn"`
	ResSuspenseAmount    string `json:"resSuspenseAmount"`
	ResBankName          string `json:"resBankName"`
	ResSalesCount        string `json:"resSalesCount"`
	CommMemberStoreGroup string `json:"commMemberStoreGroup"`
	CommStartDate        string `json:"commStartDate"`
}

// 일자별 입금금액 합계
func GetDepositByDailySum(startDate, endDate string) map[string]interface{} {
	client := db.GetMongoDBConnect(properties.TEST_STORE_COLLECTION)

	filter := bson.M{
		"resDepositDate": bson.M{"$gte": startDate, "$lte": endDate},
	}

	cursor, err := client.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	dataMap := map[string]map[string]int{}
	var totalAmount int

	for cursor.Next(context.Background()) {
		var elem *DepositData

		err := cursor.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		// 총입금 금액
		tempTotalAmount, _ := strconv.Atoi(elem.ResAccountIn)
		totalAmount += tempTotalAmount

		// 입금금액
		resAccountIn, _ := strconv.Atoi(elem.ResAccountIn)

		if dataMap[elem.ResDepositDate] == nil {
			tempMap := map[string]int{}
			tempMap["sum_of_capture_amount"] = resAccountIn

			dataMap[elem.ResDepositDate] = tempMap
		} else {
			dataMap[elem.ResDepositDate]["sum_of_capture_amount"] += resAccountIn
		}
	}

	cursor.Close(context.Background())
	db.GetMongoDBDisConnect(client)

	resultMap := map[string]interface{}{}
	resultMap["calendar_data"] = dataMap    // 일자별 입금 금액
	resultMap["total_amount"] = totalAmount // 총 입금 금액

	return resultMap
}

// 카드사별 입금금액 합계
func GetDepositByCardSum(date string) map[string]interface{} {
	client := db.GetMongoDBConnect(properties.TEST_STORE_COLLECTION)

	filter := bson.M{
		"resDepositDate": date,
	}

	cursor, err := client.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	dataMap := map[string]map[string]int{}
	var totalAmount int

	for cursor.Next(context.Background()) {
		var elem *DepositData

		err := cursor.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		// 총입금 금액
		tempTotalAmount, _ := strconv.Atoi(elem.ResAccountIn)
		totalAmount += tempTotalAmount

		// 입금금액
		resAccountIn, _ := strconv.Atoi(elem.ResAccountIn)

		if dataMap[elem.ResCardCompany] == nil {
			tempMap := map[string]int{}
			tempMap["sum_of_capture_amount"] = resAccountIn

			dataMap[elem.ResCardCompany] = tempMap
		} else {
			dataMap[elem.ResCardCompany]["sum_of_capture_amount"] += resAccountIn
		}
	}

	cursor.Close(context.Background())
	db.GetMongoDBDisConnect(client)

	resultMap := map[string]interface{}{}
	resultMap["card_data"] = dataMap        // 카드사별 입금 금액
	resultMap["total_amount"] = totalAmount // 총 입금 금액

	return resultMap
}
