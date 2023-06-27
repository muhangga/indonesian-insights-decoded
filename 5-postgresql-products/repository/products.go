package repository

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db}
}

func (r *ProductRepository) FetchProduct() ([]Product, error) {

	var products []Product
	var sqlStatement = `SELECT id, name, price FROM products`

	rows, err := r.db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product Product
		err = rows.Scan(&product.ID, &product.Name, &product.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func (r *ProductRepository) SaveProduct(product Product) error {
	var sqlStatement = `INSERT INTO products (name, price) VALUES ($1, $2) RETURNING id`

	_, err := r.db.Exec(sqlStatement, product.Name, product.Price)
	if err != nil {
		return err
	}
	return nil
}
