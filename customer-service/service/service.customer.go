package service

import (
	"example/customer-service/model"
	"example/customer-service/repository"
)

type CustomerService interface {
	Signup(customer *model.Customer) (int64, error)
	Update(id uint, customer *model.Customer) error
	Delete(id uint) error
	FindAll() ([]*model.Customer, error)
	FindByID(id uint) (*model.Customer, error)

	FindAllAddresses(id uint) ([]model.Address, error)
}

type customerService struct {
	repository repository.CustomerRepository
}

func (c *customerService) Signup(customer *model.Customer) (int64, error) {
	return c.repository.Save(customer)
}

func (c *customerService) Update(id uint, customer *model.Customer) error {
	return c.repository.Update(id, customer)
}

func (c *customerService) Delete(id uint) error {
	return c.repository.Delete(id)
}

func (c *customerService) FindAll() ([]*model.Customer, error) {
	return c.repository.FindAll()
}

func (c *customerService) FindByID(id uint) (*model.Customer, error) {
	return c.repository.FindByID(id)
}

func (c *customerService) FindAllAddresses(id uint) ([]model.Address, error) {
	if customer, err := c.repository.FindByID(id); err == nil {
		return customer.Addresses, nil
	} else {
		return nil, err
	}
}

func NewCustomerService(repository repository.CustomerRepository) CustomerService {
	return &customerService{
		repository: repository,
	}
}
