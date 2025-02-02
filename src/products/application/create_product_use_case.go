package application

import (
	"api-hexagonal/src/products/domain/entities"
	"api-hexagonal/src/products/domain"
	
)
type CreateProductUseCase struct {
	db domain.IProduct
}

func NewCreateProductUseCase(db domain.IProduct) *CreateProductUseCase {
	return &CreateProductUseCase{db: db}
}

func (cp *CreateProductUseCase) Execute(product *entities.Product) (*entities.Product, error) {
	
	productCreated, err := cp.db.CreateProduct(*product)

	if(err != nil){
		return nil, err
	}
	return productCreated, nil
}