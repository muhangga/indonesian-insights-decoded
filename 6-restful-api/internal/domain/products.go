package domain

type Product struct {
	ID    int64   `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type ProductRequest struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
