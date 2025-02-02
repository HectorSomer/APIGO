package application

import (
	"nombre-del-proyecto/src/products/domain"
    "fmt"
)

type DeleteProductUseCase struct {
	db domain.IProduct
}

func NewDeleteProductUseCase(db domain.IProduct) *DeleteProductUseCase {
    return &DeleteProductUseCase{db: db}
}
func (gp *DeleteProductUseCase) Execute(id int) (bool, error) {
	state, err := gp.db.DeleteProduct(id)
	fmt.Println("state", state)
    if err != nil {
        return  false, err
    }
    return state, nil
}