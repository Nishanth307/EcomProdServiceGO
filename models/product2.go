package models

import "time"

type Asamplemodel struct {
	ID        string    `json:"_id" bson:"_id"`
	StartDate time.Time `json:"start_date" bson:"start_date"`
	EndDate   time.Time `json:"end_date" bson:"end_date"`
}
