package repository

type CRUDRepository[T any] interface {
	Save(entity T) (int64, error)
	Update(id uint, entity T) error
	Delete(id uint) error
	FindAll() ([]T, error)
	FindByID(id uint) (T, error)
}
