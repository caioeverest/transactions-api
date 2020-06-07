package model

import "time"

type Account struct {
	AccountID      int    `json:"account_id" gorm:"Column:account_id;primary_key;AUTO_INCREMENT"`
	DocumentNumber string `json:"document_number" gorm:"Column:document_number;index:doc_num;not null;UNIQUE"`
}

type Transaction struct {
	TransactionID   int       `json:"transaction_id" gorm:"Column:transaction_id;primary_key;AUTO_INCREMENT"`
	AccountID       int       `json:"account_id" gorm:"Column:account_id;index:acc_id;not null"`
	OperationTypeId int       `gorm:"Column:operation_type_id;index:op_id;not null"`
	Amount          float64   `json:"amount" gorm:"Column:amount;not null"`
	EventDate       time.Time `json:"event_date" gorm:"Column:event_date;DEFAULT:now();not null"`
}

type Operation struct {
	OperationTypeID int    `json:"operation_type_id" gorm:"Column:operation_type_id;primary_key;AUTO_INCREMENT"`
	Description     string `json:"description" gorm:"Column:description;not null"`
}
