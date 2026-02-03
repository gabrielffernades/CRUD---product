package product

func GetProducts() []Product {
	return FindAllProducts()
}

func CreateProduct(input CreateProductInput) Product {
	product := Product{
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
	}

	SaveProduct(product)

	return product
}
