package application

import "nombre-del-proyecto/src/sells/domain"

type DeleteSellUseCase struct {
	db domain.ISell
}

func NewDeleteSellUseCase(db domain.ISell) *DeleteSellUseCase {
    return &DeleteSellUseCase{db: db}
}

func (ds *DeleteSellUseCase) Execute(id int) (bool, error) {
	state, err := ds.db.DeleteSell(id)
	if err!= nil {
        return false, err
    }
	return state, nil
}