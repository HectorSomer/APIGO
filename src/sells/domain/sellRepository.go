package domain

import (
	"api-hexagonal/src/sells/domain/entities"
)

type ISell interface {
	CreateSell(sell entities.Sell) (*entities.Sell, error)
	GetAllSells() (*[]entities.Sell, error)
	EditSell(id int, sellToUpdate entities.UpdatedSell) (*entities.UpdatedSell, error)
	DeleteSell(id int) (bool, error)
}

