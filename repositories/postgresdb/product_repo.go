package postgresdb

import (
	// Go Internal Packages
	"context"
	"errors"
	
	// Local Packages
	models "products/models"

	// External Packages
	"database/sql"
	"github.com/lib/pq"
)

type postgresdb struct {
	db *sql.DB
}

func NewPostgresDB(db *sql.DB) *postgresdb {
	return &postgresdb{db: db}
}

func (p *postgresdb) GetProductById(ctx context.Context, id int) (*models.Product, error) {
	product := &models.Product{}
	query := `SELECT id, name, description, price FROM products WHERE id = $1`
	err := p.db.QueryRowContext(ctx, query, id).Scan(&product.Id, &product.Name, &product.Description, &product.Price)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("Product not found")
		}
		return nil, err
	}
	return product, nil
}

func (p *postgresdb) GetAllProducts(ctx context.Context) ([]models.Product, error) {
	var products []models.Product
	query := `SELECT id, name, description, price FROM products`
	rows, err := p.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		product := models.Product{}
		err = rows.Scan(&product.Id, &product.Name, &product.Description, &product.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (p *postgresdb) CreateProduct(ctx context.Context, product models.Product) error {
	query := `INSERT INTO products (name, description, price) VALUES ($1, $2, $3) RETURNING id`
	err := p.db.QueryRowContext(ctx, query, product.Name, product.Description, product.Price).Scan(&product.Id)
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			if pgErr.Code == "23505" {
				return errors.New("Product already exists")
			}
		}
		return err
	}
	return nil
}

func (p *postgresdb) UpdateProduct(ctx context.Context, id int, product models.Product) error {
	query := `UPDATE products SET name = $1, description = $2, price = $3 WHERE id = $4`
	_, err := p.db.ExecContext(ctx, query, product.Name, product.Description, product.Price, id)
	if err != nil {
		return err
	}
	return nil
}

func (p *postgresdb) DeleteProductById(ctx context.Context, id int) error {
	query := `DELETE FROM products WHERE id = $1`
	_, err := p.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}