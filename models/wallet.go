package models

type Wallet struct {
	Username string  `json:"username"`
	Balance  float64 `json:"balance"`
}

type DataResponse map[string]Wallet

type ServiceResponse map[string]Wallet

type Wlt struct {
	Balance float64 `json:"balance"`
}
