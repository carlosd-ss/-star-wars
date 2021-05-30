package handler

import "go.mongodb.org/mongo-driver/mongo"

type Repository struct {
	Client     *mongo.Client
	Db         string
	Collection string
}

type Service struct {
	R *Repository
}
