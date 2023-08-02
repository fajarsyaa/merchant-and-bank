package model

type CustomerModel struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	NoRek    string `json:"NoRek"`
	Balance  int    `json:"balance"`
}
