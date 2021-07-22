package entity

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
