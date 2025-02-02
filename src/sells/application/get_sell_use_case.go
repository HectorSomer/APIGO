package application

import (
	"nombre-del-proyecto/src/sells/domain"
	"nombre-del-proyecto/src/sells/domain/entities"
	"fmt"
)

type GetAllSellsUseCase struct{
	db domain.ISell
}

func NewGetAllSellsUseCase(db domain.ISell) *GetAllSellsUseCase {
    return &GetAllSellsUseCase{db: db}
}

func (gp *GetAllSellsUseCase) Execute() (*[]entities.Sell, error) {
	sell, err := gp.db.GetAllSells()
	fmt.Println("sells", sell)
    if err != nil {
        return nil, err
    }
    return sell, nil
}