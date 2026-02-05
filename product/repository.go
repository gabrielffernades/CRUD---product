package product

var products = []Product{}

var nextID = 1

func SaveProduct(product Product) Product {
	product.ID = nextID
	nextID++

	products = append(products, product)

	return product
}

func FindAllProducts() []Product {
	return products
}

func FindProductByID(id int) (Product, bool) {
	for _, product := range products {
		if product.ID == id {
			return product, true
		}
	}
	return Product{}, false
}

func UpdateProductByID(id int, input UpdateProductInput) (Product, bool) {
	for i, product := range products {
		if product.ID == id {
			product.Name = input.Name
			product.Description = input.Description
			product.Price = input.Price

			products[i] = product
			return product, true
		}
	}
	return Product{}, false
}

func DeleteProductByID(id int) bool {
	for i, product := range products {
		if product.ID == id {
			products = append(products[:i], products[i+1:]...)
			return true
		}
	}
	return false
}
