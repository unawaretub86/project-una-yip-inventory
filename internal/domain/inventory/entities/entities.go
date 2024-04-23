package entities

type (
	TechItem struct {
		ID          int64   `json:"id" uri:"id"`
		Name        string  `json:"name"`
		Category    string  `json:"category"`
		Description string  `json:"description"`
		Quantity    int     `json:"quantity"`
		Price       float64 `json:"price"`
	}

	Inventory []TechItem
)
