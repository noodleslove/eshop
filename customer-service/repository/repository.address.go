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

func (a *addressRepository) Save(address *model.Address) (int64, error) {
	tx := a.db.Save(address)
	return tx.RowsAffected, tx.Error
}

func (a *addressRepository) Update(id uint, address *model.Address) error {
	tx := a.db.Model(&model.Address{}).Where("id = ?", id).Updates(address)
	return tx.Error
}

func (a *addressRepository) Delete(id uint) error {
	tx := a.db.Where("id = ?", id).Delete(&model.Address{})
	return tx.Error
}

func (a *addressRepository) FindAll() ([]*model.Address, error) {
	var addresses []*model.Address
	tx := a.db.Find(&addresses)
	return addresses, tx.Error
}

func (a *addressRepository) FindByID(id uint) (*model.Address, error) {
	var address model.Address
	tx := a.db.Where("id = ?", id).First(&address)
	return &address, tx.Error
}

func NewAddressRepository(db *gorm.DB) AddressRepository {
	return &addressRepository{db}
}
