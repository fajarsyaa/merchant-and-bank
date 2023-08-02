package model

type MerchantModel struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	NoRek   string `json:"NoRek"`
	Balance int    `json:"balance"`
}
