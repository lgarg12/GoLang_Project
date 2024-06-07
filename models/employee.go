package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Employee struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"name" bson:"name"`
	Position string             `json:"position" bson:"position"`
	Salary   int                `json:"salary" bson:"salary"`
}
