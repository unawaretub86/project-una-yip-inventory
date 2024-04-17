package entities

type (
	TechItem struct {
		ID          int
		Name        string
		Category    string
		Description string
		Quantity    int
		Price       float64
	}

	Inventory []TechItem
)
