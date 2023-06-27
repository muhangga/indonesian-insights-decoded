package repository

import (
	"github.com/muhangga/internal/domain"
	"gorm.io/gorm"
)

type ProductsRepository interface {
	GetAll() ([]domain.Product, error)
	Save(product domain.Product) (domain.Product, error)
	Update(product domain.Product) (domain.Product, error)
	Delete(id int) (domain.Product, error)
	GetByID(id int) (domain.Product, error)
}

type productsRepository struct {
	*gorm.DB
}

func NewProductsRepository(db *gorm.DB) *productsRepository {
	return &productsRepository{db}
}

func (r *productsRepository) GetAll() ([]domain.Product, error) {
	var products []domain.Product
	err := r.Find(&products).Error
	if err != nil {
		return products, err
	}
	return products, nil
}

func (r *productsRepository) Save(product domain.Product) (domain.Product, error) {
	tx := r.Begin()
	err := tx.Create(&product).Error
	if err != nil {
		tx.Rollback()
		return product, err
	}
	tx.Commit()

	return product, nil
}

func (r *productsRepository) Update(product domain.Product) (domain.Product, error) {

	tx := r.Begin()
	save := tx.Table("products").Save(&product).Error
	if save != nil {
		tx.Rollback()
		return product, save
	}
	tx.Commit()

	return product, nil
}

func (r *productsRepository) Delete(id int) (domain.Product, error) {
	var product domain.Product
	err := r.Where("id = ?", id).Delete(&product).Error
	if err != nil {
		return product, err
	}
	return product, nil
}

func (r *productsRepository) GetByID(id int) (domain.Product, error) {
	var product domain.Product
	err := r.Where("id = ?", id).First(&product).Error
	if err != nil {
		return product, err
	}
	return product, nil
}
