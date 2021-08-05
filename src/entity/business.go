package entity

type BusinessData struct {
	Data   []CodefData
	Result CodefResult
}

type CodefData struct {
	Code                   string `json:"code" example:"CF-00000"`
	Message                string `json:"message" example:"성공"`
	ExtraMessage           string `json:"extraMessage" example:""`
	ResBusinessStatus      string `json:"resBusinessStatus" example:"부가가치세 일반과세자 입니다."`
	ResClosingDate         string `json:"resClosingDate" example:""`
	ResCompanyIdentityNo   string `json:"resCompanyIdentityNo" example:"1231212345"`
	ResTaxationTypeCode    string `json:"resTaxationTypeCode" example:"1"`
	ResTransferTaxTypeDate string `json:"resTransferTaxTypeDate" example:""`
}

type CodefResult struct {
	Code          string `json:"code"`
	ExtraMessage  string `json:"extraMessage"`
	Message       string `json:"message"`
	TransactionId string `json:"transactionId"`
}
