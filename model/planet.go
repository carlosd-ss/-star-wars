package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Planet struct {
	ID      primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Name    string             `bson:"name" json:"name"`
	Climate string             `bson:"climate" json:"climate"`
	Terrain string             `bson:"terrain" json:"terrain"`
}
