package repo

import "ecommerce/domain"

type ProductRepo interface {
	Create(p domain.Product) (*domain.Product, error)
	Get(productID int) (*domain.Product, error)
	List() ([]*domain.Product, error)
	Delete(productID int) error
	Update(p domain.Product) (*domain.Product, error)
}

type productRepo struct {
	productList []*domain.Product
}

func NewProductRepo() ProductRepo {
	repo := &productRepo{}

	generateInitialProducts(repo)
	return repo
}

func (r *productRepo) Create(p domain.Product) (*domain.Product, error) {
	p.ID = len(r.productList) + 1
	r.productList = append(r.productList, &p)

	return &p, nil
}

func (r *productRepo) Get(productID int) (*domain.Product, error) {
	for _, product := range r.productList {
		if product.ID == productID {
			return product, nil
		}
	}
	return nil, nil
}

func (r *productRepo) List() ([]*domain.Product, error) {
	return r.productList, nil
}

func (r *productRepo) Delete(productID int) error {
	var tempList []*domain.Product

	for _, product := range r.productList {
		if product.ID != productID {
			tempList = append(tempList, product)
		}
	}
	r.productList = tempList
	return nil
}

func (r *productRepo) Update(product domain.Product) (*domain.Product, error) {
	for idx, p := range r.productList {
		if p.ID == product.ID {
			r.productList[idx] = &product
		}
	}

	return &product, nil
}

func generateInitialProducts(r *productRepo) {
	product1 := &domain.Product{
		ID:          1,
		Title:       "Wireless Headphones",
		Description: "High-quality noise-cancelling headphones.",
		Price:       129.99,
		ImgURL:      "https://www.lovefoodhatewaste.com/sites/default/files/styles/twitter_card_image/public/2022-07/Citrus%20fruits.jpg.webp?itok=H1j9CCCS",
	}
	product2 := &domain.Product{
		ID:          2,
		Title:       "Smart Watch",
		Description: "Stylish smart watch with health tracking.",
		Price:       199.99,
		ImgURL:      "https://i0.wp.com/post.healthline.com/wp-content/uploads/2021/05/apples-1296x728-header.jpg?w=1155&h=1528",
	}
	product3 := &domain.Product{
		ID:          3,
		Title:       "Running Shoes",
		Description: "Lightweight shoes for everyday running.",
		Price:       89.50,
		ImgURL:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRZbMOVB8a8wRQ6e-UKZggiu7-edRAN1GolPQ&s",
	}
	product4 := &domain.Product{
		ID:          4,
		Title:       "Wireless Headphones",
		Description: "High-quality noise-cancelling headphones.",
		Price:       129.99,
		ImgURL:      "https://www.lovefoodhatewaste.com/sites/default/files/styles/twitter_card_image/public/2022-07/Citrus%20fruits.jpg.webp?itok=H1j9CCCS",
	}
	product5 := &domain.Product{
		ID: 5, Title: "Smart Watch",
		Description: "Stylish smart watch with health tracking.",
		Price:       199.99,
		ImgURL:      "https://i0.wp.com/post.healthline.com/wp-content/uploads/2021/05/apples-1296x728-header.jpg?w=1155&h=1528",
	}
	product6 := &domain.Product{
		ID:          6,
		Title:       "Running Shoes",
		Description: "Lightweight shoes for everyday running.",
		Price:       89.50,
		ImgURL:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRZbMOVB8a8wRQ6e-UKZggiu7-edRAN1GolPQ&s",
	}

	r.productList = append(r.productList, product1)
	r.productList = append(r.productList, product2)
	r.productList = append(r.productList, product3)
	r.productList = append(r.productList, product4)
	r.productList = append(r.productList, product5)
	r.productList = append(r.productList, product6)
}
