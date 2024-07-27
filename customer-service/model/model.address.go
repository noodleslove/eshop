package model

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	Street1    string `gorm:"not null" json:"Street1"`
	Street2    string `json:"Street2"`
	City       string `gorm:"not null" json:"City"`
	State      string `gorm:"not null" json:"State"`
	PostalCode string `gorm:"not null" json:"PostalCode"`
	Country    string `gorm:"not null" json:"Country"`

	Customers []Customer `gorm:"many2many:customer_addresses;" json:"customers"`
}
