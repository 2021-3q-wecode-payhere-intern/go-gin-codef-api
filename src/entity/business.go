package entity

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
