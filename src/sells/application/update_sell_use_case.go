package application

import (
	"nombre-del-proyecto/src/sells/domain"
	"nombre-del-proyecto/src/sells/domain/entities"
	"fmt"
)

type UpdateSellUseCase struct {
	db domain.ISell
}

func NewUpdateSellUseCase(db domain.ISell) *UpdateSellUseCase {
    return &UpdateSellUseCase{db: db}
}

func (gp *UpdateSellUseCase) Execute(id int, sellToUpdate entities.UpdatedSell) (*entities.UpdatedSell, error) {
	sell, err := gp.db.EditSell(id, sellToUpdate)
	fmt.Println("products", sell)
    if err != nil {
        return nil, err
    }
    return sell, nil
}
