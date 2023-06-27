package service

import (
	"errors"

	"github.com/muhangga/internal/domain"
	"github.com/muhangga/internal/repository"
)

type ProductsService interface {
	FetchAllProducts() ([]domain.Product, error)
	CreateProduct(product domain.ProductRequest) (domain.Product, error)
	UpdateProduct(id int, product domain.ProductRequest) (domain.Product, error)
	DeleteProduct(id int) (domain.Product, error)
}

type productsService struct {
	productsRepo repository.ProductsRepository
}

func NewProductsService(productsRepo repository.ProductsRepository) *productsService {
	return &productsService{productsRepo}
}

func (s *productsService) FetchAllProducts() ([]domain.Product, error) {
	products, err := s.productsRepo.GetAll()
	if err != nil {
		return products, err
	}
	return products, nil
}

func (s *productsService) CreateProduct(productRequest domain.ProductRequest) (domain.Product, error) {
	if productRequest.Name == "" {
		return domain.Product{}, errors.New("name is required")
	}
	if productRequest.Price == 0 {
		return domain.Product{}, errors.New("price is required")
	}
	productToSave := domain.Product{
		Name:  productRequest.Name,
		Price: productRequest.Price,
	}

	product, err := s.productsRepo.Save(productToSave)
	if err != nil {
		return product, err
	}
	return product, nil
}

func (s *productsService) UpdateProduct(id int, productRequest domain.ProductRequest) (domain.Product, error) {
	if productRequest.Name == "" {
		return domain.Product{}, errors.New("name is required")
	}
	if productRequest.Price == 0 {
		return domain.Product{}, errors.New("price is required")
	}

	productId, err := s.productsRepo.GetByID(id)
	if err != nil {
		return productId, errors.New("product not found")
	}

	productId.Name = productRequest.Name
	productId.Price = productRequest.Price

	product, err := s.productsRepo.Update(productId)
	if err != nil {
		return product, err
	}

	return product, nil
}

func (s *productsService) DeleteProduct(id int) (domain.Product, error) {
	product, err := s.productsRepo.Delete(id)
	if err != nil {
		return product, err
	}
	return product, nil
}
