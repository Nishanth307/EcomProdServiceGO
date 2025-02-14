package main

import (
	// Go Internal Packages
	"fmt"
	"log"
	"net/http"
	
	// Local Packages
	config "products/config"
	handlers "products/handlers"
	mongorepo "products/repositories/mongodb"
	clickhouse "products/repositories/clickhouse"
	postgresrepo "products/repositories/postgresdb"
	services "products/services"

	// External Packages
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Starting the server...")
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	fmt.Println("Successfully loaded the config file!")
	ServerInitialization(cfg)
}

func ServerInitialization(k *config.Config) {

	//ClickHouse Connect
	clickhouseRepo, err := clickhouse.Connect(k.ClickHouse.URI)
	if err != nil {
		log.Fatalf("Failed to connect to ClickHouse: %v", err)
	}
	defer clickhouseRepo.Conn.Close()
	


	// Connecting to DB
	var productRepo services.ProductRepository
	if k.Bool.Enabled {
		mongo_client, err := mongorepo.Connect(k.Mongo.URI) 
		if err != nil {
			log.Fatalf("Failed to connect to db:%v", err)
		}
		productRepo = mongorepo.NewProductRepository(mongo_client) 
	} else{
		
		postgresCli, err := postgresrepo.Connect(k.Postgres.URI)
		if err != nil {
			log.Fatalf("PostgreSQL connection error: %v", err)
		}
		productRepo = postgresrepo.NewPostgresDB(postgresCli)
	}
	// productRepo = clickhouse.Connect()

	productService := services.NewService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	router := mux.NewRouter()
	router.HandleFunc("/products/{id}", productHandler.GetProductById).Methods("GET")
	router.HandleFunc("/products", productHandler.GetAllProducts).Methods("GET")
	router.HandleFunc("/products", productHandler.CreateProduct).Methods("POST")
	router.HandleFunc("/products/{id}", productHandler.UpdateProduct).Methods("PUT")
	router.HandleFunc("/products/{id}", productHandler.DeleteProductById).Methods("DELETE")
	fmt.Println("Server is running on port", k.Port.Port)
	// http.ListenAndServe(":8080", router)
	log.Fatal(http.ListenAndServe(":8080", router))
}
