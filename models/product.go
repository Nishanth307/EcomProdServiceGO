package models

type Product struct {
	// bson is the key used in MongoDB json is the key used in JSON
	Id          int     `json:"id" bson:"_id"`
	Price       float32 `json:"price" bson:"price"`
	Name        string  `json:"name" bson:"name"`
	Description string  `json:"description" bson:"description"`
}
