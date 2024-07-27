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

func (c *customerRepository) Save(customer *model.Customer) (int64, error) {
	tx := c.db.Save(customer)
	return tx.RowsAffected, tx.Error
}

func (c *customerRepository) Update(id uint, customer *model.Customer) error {
	tx := c.db.Model(&model.Customer{}).Where("id = ?", id).Updates(customer)
	return tx.Error
}

func (c *customerRepository) Delete(id uint) error {
	tx := c.db.Where("id = ?", id).Delete(&model.Customer{})
	return tx.Error
}

func (c *customerRepository) FindAll() ([]*model.Customer, error) {
	var customers []*model.Customer
	tx := c.db.Find(&customers)
	return customers, tx.Error
}

func (c *customerRepository) FindByID(id uint) (*model.Customer, error) {
	var customer model.Customer
	tx := c.db.Where("id = ?", id).First(&customer)
	return &customer, tx.Error
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepository{db}
}
