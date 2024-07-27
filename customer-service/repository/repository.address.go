package repository

import (
	"example/customer-service/model"

	"gorm.io/gorm"
)

type AddressRepository interface {
	CRUDRepository[*model.Address]
}

type addressRepository struct {
	db *gorm.DB
}

// Save is a method to save an address to the database
func (a *addressRepository) Save(address *model.Address) (int64, error) {
	// Save the address to the database
	tx := a.db.Save(address)
	// Return the number of rows affected and the error
	return tx.RowsAffected, tx.Error
}

// Update is a method to update an address in the database
func (a *addressRepository) Update(id uint, address *model.Address) error {
	// Update the address in the database
	tx := a.db.Model(&model.Address{}).Where("id = ?", id).Updates(address)
	// Return the error
	return tx.Error
}

// Delete is a method to delete an address from the database
func (a *addressRepository) Delete(id uint) error {
	// Delete the address from the database
	tx := a.db.Where("id = ?", id).Delete(&model.Address{})
	// Return the error
	return tx.Error
}

// FindAll is a method to find all addresses in the database
func (a *addressRepository) FindAll() ([]*model.Address, error) {
	// Initialize a slice of addresses
	var addresses []*model.Address
	// Find all addresses in the database
	tx := a.db.Find(&addresses)
	// Return the addresses and the error
	return addresses, tx.Error
}

// FindByID is a method to find an address by its ID in the database
func (a *addressRepository) FindByID(id uint) (*model.Address, error) {
	// Initialize an address
	var address model.Address
	// Find the address by its ID in the database
	tx := a.db.Where("id = ?", id).First(&address)
	// Return the address and the error
	return &address, tx.Error
}

// NewAddressRepository is a factory method to create a new AddressRepository
func NewAddressRepository(db *gorm.DB) AddressRepository {
	return &addressRepository{db}
}
