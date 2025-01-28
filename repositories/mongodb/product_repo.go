package repository

import (
	// Go Internal Packages
	"context"

	// Local Packages
	"products/models"

	// External Packages
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepository struct {
	db *mongo.Collection
	// client *mongo.Client
}

func NewProductRepository(client *mongo.Client) *ProductRepository { //,db *mongo.Collection
	collection := client.Database("product-db").Collection("product")
	return &ProductRepository{db: collection}
	// return &ProductRepo{client: client,db: db}
}

func (r *ProductRepository) GetProductById(ctx context.Context, id int) (*models.Product, error) {
	filter := bson.D{bson.E{Key: "_id", Value: id}}
	var product models.Product
	err := r.db.FindOne(ctx, filter).Decode(&product) // find one document and decode it into product
	if err != nil {
		return nil, err
	}
	return &product, err
}

func (r *ProductRepository) GetAllProducts(ctx context.Context) ([]models.Product, error) {
	cursor, err := r.db.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx) // close the cursor once the function is done
	var products []models.Product
	for cursor.Next(ctx) {
		var product models.Product
		err := cursor.Decode(&product)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (r *ProductRepository) CreateProduct(ctx context.Context, product models.Product) error {
	_, err := r.db.InsertOne(ctx, product)
	return err
}

func (r *ProductRepository) UpdateProduct(ctx context.Context, id int, product models.Product) error {
	// collection := r.client.Database("product-db").Collection("product")
	filter := bson.D{bson.E{Key: "_id", Value: id}}
	update := bson.D{bson.E{Key: "$set", Value: product}}
	_, err := r.db.UpdateOne(ctx, filter, update)
	return err
}

func (r *ProductRepository) DeleteProductById(ctx context.Context, id int) error {
	// collection := r.client.Database("product-db").Collection("product")
	filter := bson.D{bson.E{Key: "_id", Value: id}} // _id is primary key
	_, err := r.db.DeleteOne(ctx, filter)
	return err
}