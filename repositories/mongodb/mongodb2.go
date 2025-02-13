package repository

import (
	// "fmt"
	"context"
	// Local Packages
	errors "products/errors"
	model "products/models"

	// External Packages
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type asamplerepo struct{
	client *mongo.Client
	collection string
}

func NewaSampleRepo(client *mongo.Client ) *asamplerepo{
	return &asamplerepo{client: client, collection:"test"}
}

func (s *asamplerepo) CreateASampleByAppId(ctx context.Context,appID string, model *model.Asamplemodel) error {
	collection := s.client.Database(appID).Collection(s.collection)
	_, err := collection.InsertOne(ctx, model)
	if err!=nil {
		return err
	}
	return nil
}

// correct
func(s *asamplerepo) GetASampleByAppId(ctx context.Context, appID string, id string) (*model.Asamplemodel,error){
	collection := s.client.Database(appID).Collection(s.collection)
	var asample model.Asamplemodel
	filter := bson.M{"_id": id}
	err := collection.FindOne(ctx, filter).Decode(&asample)
	
	if err!=nil{
		return nil,err
	}
	return &asample,nil	
}

func(s *asamplerepo) UpdateASampleByAppId(ctx context.Context, appID string, model *model.Asamplemodel) error{
	collection := s.client.Database(appID).Collection(s.collection)

	filter := bson.M{"_id": model.ID}
	update := bson.M{"$set": model}
	_,err := collection.UpdateOne(ctx, filter,update)
	if err!=nil{
		return err
	}
	return nil
}

func(s *asamplerepo) DeleteASampleByAppId(ctx context.Context, appID string, id string) error{
	collection := s.client.Database(appID).Collection(s.collection)
	filter := bson.M{"_id": id}
	res,err := collection.DeleteOne(ctx, filter)
	if err!=nil{
		return err
	}
	if res.DeletedCount == 0{
		return errors.E(errors.NotFound)
	}
	return nil
}