package request

type TransactionRequestModel struct {
	MerchantNoRek string `json:"merchant_rek"`
	Amount        int    `json:"amount"`
}
