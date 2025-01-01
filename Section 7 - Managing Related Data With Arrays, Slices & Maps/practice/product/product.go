package product

type Product struct {
    id int
    title string
    price float64
}

func New(id int, title string, price float64) Product {
    return Product{
        id: id,
        title: title,
        price: price,
    }
}