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

func (a *addressService) Add(address *model.Address) (int64, error) {
	return a.repository.Save(address)
}

func (a *addressService) Update(id uint, address *model.Address) error {
	return a.repository.Update(id, address)
}

func (a *addressService) Delete(id uint) error {
	return a.repository.Delete(id)
}

func (a *addressService) FindByID(id uint) (*model.Address, error) {
	return a.repository.FindByID(id)
}

func (a *addressService) FindAllCustomers(id uint) ([]model.Customer, error) {
	if address, err := a.repository.FindByID(id); err == nil {
		return address.Customers, nil
	} else {
		return nil, err
	}
}

func NewAddressService(repository repository.AddressRepository) AddressService {
	return &addressService{
		repository: repository,
	}
}
