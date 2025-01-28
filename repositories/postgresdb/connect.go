package postgresdb

import (
	// Go Internal Packages
	"context"
	"fmt"
	"log"
	"time"

	// External Packages
	"database/sql"
	_ "github.com/lib/pq"
)

func Connect(uri string) (*sql.DB, error) {
	fmt.Println("Starting PostgreSQL Database Connection")

	// Open a connection to the database
	db, err := sql.Open("postgres", uri)
	if err != nil {
		return nil, fmt.Errorf("failed to open PostgreSQL connection: %w", err)
	}

	// Create a context with a timeout to ping the database
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Ping the database to check if it's reachable
	if err := db.PingContext(ctx); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping PostgreSQL: %w", err)
	}

	// Ensure the 'products' table exists
	_, err = db.Exec(`
		DROP TABLE IF EXISTS products;
		CREATE TABLE IF NOT EXISTS products (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			price NUMERIC NOT NULL,
			description VARCHAR(255) NOT NULL
		)
	`)
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ensure products table exists: %w", err)
	}

	// Insert data into the products table (including description)
	_, err = db.Exec("INSERT INTO products (name, price, description) VALUES ($1, $2, $3)", "Sample Product", 99.99, "This is a sample product description.")
	if err != nil {
		db.Close()
		log.Fatal("Error inserting product:", err)
	} else {
		fmt.Println("Data inserted successfully!")
	}

	fmt.Println("Connected to PostgreSQL")
	return db, nil
}