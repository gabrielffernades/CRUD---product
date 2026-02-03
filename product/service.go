package product

func GetProducts() []Product {
	return []Product{
		{
			ID:          1,
			Name:        "Teclado gamer",
			Description: "Teclado gamer com iluminação RGB",
			Price:       800.00,
		},
	}
}
