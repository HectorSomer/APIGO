package domain

import (
	"nombre-del-proyecto/src/products/domain/entities"
)

type IProduct interface {
	CreateProduct(product entities.Product) (*entities.Product, error)
	GetProducts() (*[]entities.Product, error)
	EditProduct(id int, productToUpdate entities.UpdateProduct) (*entities.UpdateProduct, error)
	DeleteProduct(id int) (bool, error)
}
