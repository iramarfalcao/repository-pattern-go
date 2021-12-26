package interfaces

import model "github.com/iramarfalcao/repository-pattern-go/models/product"

type IProductRepository interface {
	Store(p model.Product) (model.Product, error)
	Delete(id uint) (model.Product, error)
	Update(p model.Product) error
	Get(id uint) (model.Product, error)
	GetAll() ([]model.Product, error)
}
