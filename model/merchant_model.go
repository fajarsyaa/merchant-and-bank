package model

type MerchantModel struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	NoRek    string `json:"NoRek"`
	Balance  int    `json:"balance"`
}
