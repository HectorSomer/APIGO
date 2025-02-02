package application

import (
	"api-hexagonal/src/products/domain"
	"api-hexagonal/src/products/domain/entities"
	"fmt"
)

type GetProductUseCase struct {
	db domain.IProduct
}

func NewGetProductUseCase(db domain.IProduct) *GetProductUseCase {
    return &GetProductUseCase{db: db}
}

func (gp *GetProductUseCase) Execute() (*[]entities.Product, error) {
	products, err := gp.db.GetProducts()
	fmt.Println("products", products)
    if err != nil {
        return nil, err
    }
    return products, nil
}