package models

import "gorm.io/datatypes"

type Transaction struct {
	ID            uint           `gorm:"primarykey" json:"id"`
	Items         string         `json:"items"`
	ShopName      string         `json:"shop_name"`
	Date          datatypes.Date `json:"date"`
	Category      string         `json:"category"`
	PaymentMethod string         `json:"payment_method"`
	Amount        int            `json:"amount"`
	TransactionBy string         `json:"transaction_by"`
}
