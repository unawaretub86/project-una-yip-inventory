package entities

type (
	TechItem struct {
		ID          int64
		Name        string
		Category    string
		Description string
		Quantity    int
		Price       float64
	}

	Inventory []TechItem
)
