package entities

type Product struct {
	ID          int32     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Stock       int     `json:"stock"`
	Price       float32 `json:"price"`
}

