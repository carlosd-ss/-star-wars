package repo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/carlosd-ss/-star-wars/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	Client     *mongo.Client
	Db         string
	Collection string
}

type Service struct {
	R *Repository
}

func (r *Repository) Create(planet model.Planet) error {
	planet.ID = primitive.NewObjectID()
	client := r.Client.Database(r.Db).Collection(r.Collection)
	_, err := client.InsertOne(context.TODO(), planet)
	fmt.Println(r)
	if err != nil {
		return err
	}
	return nil
}

//func (r *Repository) SearchForId(Id primitive.ObjectID) (planet model.Planet, err error) {
//	result := model.Planet{}
//	filter := bson.D{primitive.E{Key: "Id", Value: Id}}
//	collection := mcon.GetCollection(r.Db, r.Collection)
//	err = collection.FindOne(context.TODO(), filter).Decode(&result)
//	if err != nil {
//		return result, err
//	}
//	return result, nil
//}
//func (r *Repository) ListAll() (pts []model.Planet, err error) {
//	filter := bson.D{{}}
//	planets := []model.Planet{}
//	collection := mcon.GetCollection(r.Db, r.Collection)
//	cur, findError := collection.Find(context.TODO(), filter)
//	if findError != nil {
//		return planets, findError
//	}
//	for cur.Next(context.TODO()) {
//		t := model.Planet{}
//		err := cur.Decode(&t)
//		if err != nil {
//			return planets, err
//		}
//		planets = append(planets, t)
//	}
//	cur.Close(context.TODO())
//	if len(planets) == 0 {
//		return planets, mongo.ErrNoDocuments
//	}
//	return planets, nil
//
//}
//func (r *Repository) Delete(Id string) (err error) {
//	filter := bson.D{primitive.E{Key: "Id", Value: Id}}
//	collection := mcon.GetCollection(r.Db, r.Collection)
//	_, err = collection.DeleteOne(context.TODO(), filter)
//	if err != nil {
//		return err
//	}
//	return nil
//}
