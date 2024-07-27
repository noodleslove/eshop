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

// Signup is a method to add a customer to the database
func (c *customerService) Signup(customer *model.Customer) (int64, error) {
	// Save the customer to the database
	return c.repository.Save(customer)
}

// Update is a method to update a customer in the database
func (c *customerService) Update(id uint, customer *model.Customer) error {
	// Update the customer in the database
	return c.repository.Update(id, customer)
}

// Delete is a method to delete a customer from the database
func (c *customerService) Delete(id uint) error {
	// Delete the customer from the database
	return c.repository.Delete(id)
}

// FindAll is a method to find all customers in the database
func (c *customerService) FindAll() ([]*model.Customer, error) {
	// Find all customers in the database
	return c.repository.FindAll()
}

// FindByID is a method to find a customer by its ID in the database
func (c *customerService) FindByID(id uint) (*model.Customer, error) {
	// Find the customer by its ID in the database
	return c.repository.FindByID(id)
}

// FindAllAddresses is a method to find all addresses associated with a customer
func (c *customerService) FindAllAddresses(id uint) ([]model.Address, error) {
	// Find the customer by its ID in the database
	if customer, err := c.repository.FindByID(id); err == nil {
		return customer.Addresses, nil
	} else {
		return nil, err
	}
}

// NewCustomerService creates a new customer service with the given customer repository
func NewCustomerService(repository repository.CustomerRepository) CustomerService {
	return &customerService{
		repository: repository,
	}
}
