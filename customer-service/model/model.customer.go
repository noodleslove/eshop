package model

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	FirstName string `gorm:"not null" json:"FirstName"`
	LastName  string `gorm:"not null" json:"LastName"`
	Gender    string `json:"Gender"`
	Phone     string `json:"Phone"`

	Addresses []Address `gorm:"many2many:customer_addresses;" json:"addresses"`
}
