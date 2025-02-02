package entities

type UpdateProduct struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Stock       int     `json:"stock"`
	Price       float32 `json:"price"`
}

