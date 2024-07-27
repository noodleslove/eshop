package service

import (
	"example/customer-service/model"
	"example/customer-service/repository"
)

type AddressService interface {
	Add(address *model.Address) (int64, error)
	Update(id uint, address *model.Address) error
	Delete(id uint) error
	FindByID(id uint) (*model.Address, error)
	FindAllCustomers(id uint) ([]model.Customer, error)
}

type addressService struct {
	repository repository.AddressRepository
}

// Add is a method to add an address to the database
func (a *addressService) Add(address *model.Address) (int64, error) {
	// Save the address to the database
	return a.repository.Save(address)
}

// Update is a method to update an address in the database
func (a *addressService) Update(id uint, address *model.Address) error {
	// Update the address in the database
	return a.repository.Update(id, address)
}

// Delete is a method to delete an address from the database
func (a *addressService) Delete(id uint) error {
	// Delete the address from the database
	return a.repository.Delete(id)
}

// FindByID is a method to find an address by its ID in the database
func (a *addressService) FindByID(id uint) (*model.Address, error) {
	// Find the address by its ID in the database
	return a.repository.FindByID(id)
}

// FindAllCustomers is a method to find all customers associated with an address
func (a *addressService) FindAllCustomers(id uint) ([]model.Customer, error) {
	// Find the address by its ID in the database
	if address, err := a.repository.FindByID(id); err == nil {
		return address.Customers, nil
	} else {
		return nil, err
	}
}

// NewAddressService creates a new address service with the given address repository
func NewAddressService(repository repository.AddressRepository) AddressService {
	return &addressService{
		repository: repository,
	}
}
