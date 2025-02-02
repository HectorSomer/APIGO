package entities

type UpdatedSell struct {
	Concept string  `json:"concept"`
	TotalPrice       float32 `json:"total_price"`
	Date    string  `json:"date"`
}