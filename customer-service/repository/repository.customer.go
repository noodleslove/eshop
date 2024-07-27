package repository

import (
	"example/customer-service/model"

	"gorm.io/gorm"
)

type CustomerRepository interface {
	CRUDRepository[*model.Customer]
}

type customerRepository struct {
	db *gorm.DB
}

// Save is a method to save a customer to the database
func (c *customerRepository) Save(customer *model.Customer) (int64, error) {
	// Save the customer to the database
	tx := c.db.Save(customer)
	// Return the number of rows affected and the error
	return tx.RowsAffected, tx.Error
}

// Update is a method to update a customer in the database
func (c *customerRepository) Update(id uint, customer *model.Customer) error {
	// Update the customer in the database
	tx := c.db.Model(&model.Customer{}).Where("id = ?", id).Updates(customer)
	// Return the error
	return tx.Error
}

// Delete is a method to delete a customer from the database
func (c *customerRepository) Delete(id uint) error {
	// Delete the customer from the database
	tx := c.db.Where("id = ?", id).Delete(&model.Customer{})
	// Return the error
	return tx.Error
}

// FindAll is a method to find all customers in the database
func (c *customerRepository) FindAll() ([]*model.Customer, error) {
	// Initialize a slice of customers
	var customers []*model.Customer
	// Find all customers in the database
	tx := c.db.Find(&customers)
	// Return the customers and the error
	return customers, tx.Error
}

// FindByID is a method to find a customer by its ID in the database
func (c *customerRepository) FindByID(id uint) (*model.Customer, error) {
	// Initialize a customer
	var customer model.Customer
	// Find the customer by its ID in the database
	tx := c.db.Where("id = ?", id).First(&customer)
	// Return the customer and the error
	return &customer, tx.Error
}

// NewCustomerRepository is a factory function to create a new CustomerRepository
func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepository{db}
}
