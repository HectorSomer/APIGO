package application

import (
	"nombre-del-proyecto/src/products/domain"
	"nombre-del-proyecto/src/products/domain/entities"
    "fmt"
)

type UpdateProductUseCase struct {
	db domain.IProduct
}

func NewUpdateProductUseCase(db domain.IProduct) *UpdateProductUseCase {
    return &UpdateProductUseCase{db: db}
}

func (gp *UpdateProductUseCase) Execute(id int, productToUpdate entities.UpdateProduct) (*entities.UpdateProduct, error) {
	product, err := gp.db.EditProduct(id, productToUpdate)
	fmt.Println("products", product)
    if err != nil {
        return nil, err
    }
    return product, nil
}
