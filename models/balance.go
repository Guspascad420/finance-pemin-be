package models

type Balance struct {
	ID            uint   `gorm:"primarykey" json:"id"`
	AccountNumber string `json:"account_number"`
	BankName      string `json:"bank_name"`
}
