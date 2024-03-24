package service

import (
	"context"
	"errors"
	"pinky-go/internal/model"
	"sync"
)

var (
	products   = []model.Product{}
	productsMu sync.Mutex // ensures thread-safe access to the products slice
)

type productServiceImpl struct{}

func NewProductService() ProductService {
	return &productServiceImpl{}
}

func (s *productServiceImpl) CreateProduct(ctx context.Context, product model.Product) (model.Product, error) {
	productsMu.Lock()
	defer productsMu.Unlock()

	product.ID = generateNewID() // Assume generateNewID generates a unique ID for the product
	products = append(products, product)
	return product, nil
}
func (s *productServiceImpl) GetAllProducts(ctx context.Context) ([]model.Product, error) {
	productsMu.Lock()
	defer productsMu.Unlock()

	return products, nil
}

func (s *productServiceImpl) GetProductByID(ctx context.Context, id string) (model.Product, error) {
	productsMu.Lock()
	defer productsMu.Unlock()

	for _, product := range products {
		if product.ID == id {
			return product, nil
		}
	}
	return model.Product{}, errors.New("product not found")
}

func (s *productServiceImpl) UpdateProduct(ctx context.Context, id string, update model.Product) (model.Product, error) {
	productsMu.Lock()
	defer productsMu.Unlock()

	for i, product := range products {
		if product.ID == id {
			// Update product fields here
			products[i] = update
			return products[i], nil
		}
	}
	return model.Product{}, errors.New("product not found")
}

func (s *productServiceImpl) DeleteProduct(ctx context.Context, id string) error {
	productsMu.Lock()
	defer productsMu.Unlock()

	for i, product := range products {
		if product.ID == id {
			products = append(products[:i], products[i+1:]...)
			return nil
		}
	}
	return errors.New("product not found")
}

func generateNewID() string {
	return ""
}
