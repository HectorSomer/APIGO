package entities 

type Sell struct {
	ID          int32     `json:"id"`
	Concept string  `json:"concept"`
	Date    string  `json:"date"`
	Total_Price       float32 `json:"total_price"`
}