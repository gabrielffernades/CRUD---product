package product

func GetProducts() []Product {
	return FindAllProducts()
}

func GetProductByID(id int) (Product, bool) {
	return FindProductByID(id)
}

func CreateProduct(input CreateProductInput) Product {
	product := Product{
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
	}

	product = SaveProduct(product)

	return product
}

func DeleteProduct(id int) bool {
	return DeleteProductByID(id)
}
