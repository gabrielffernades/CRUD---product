package product

var products []Product

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
