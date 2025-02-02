package application

import (
	"nombre-del-proyecto/src/sells/domain"
	"nombre-del-proyecto/src/sells/domain/entities"
	"fmt"
)

type CreateSellUseCase struct {
	db domain.ISell
}

func NewCreateSellUseCase(db domain.ISell) *CreateSellUseCase {
    return &CreateSellUseCase{db: db}
}
func (gp *CreateSellUseCase) Execute(sell entities.Sell) (*entities.Sell, error) {
	product, err := gp.db.CreateSell(sell)
	fmt.Println("sells", product)
    if err != nil {
        return nil, err
    }
    return product, nil
}
