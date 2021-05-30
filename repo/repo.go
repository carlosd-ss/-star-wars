package repo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/carlosd-ss/star-wars/model"
	"go.mongodb.org/mongo-driver/mongo"
)

func Create(client *mongo.Client, db, collection string, planet model.Planet) error {
	planet.ID = primitive.NewObjectID()
	res := client.Database(db).Collection(collection)
	_, err := res.InsertOne(context.TODO(), planet)
	if err != nil {
		return err
	}
	return nil
}

func SearchForId(client *mongo.Client, db, collection string, id primitive.ObjectID) (planet model.Planet, err error) {
	result := model.Planet{}
	filter := bson.M{"_id": id}
	res := client.Database(db).Collection(collection)
	err = res.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func SearchForName(client *mongo.Client, db, collection string, name string) (planet model.Planet, err error) {
	result := model.Planet{}
	filter := bson.M{"name": name}
	res := client.Database(db).Collection(collection)
	err = res.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func ListAll(client *mongo.Client, db, collection string) (pts []model.Planet, err error) {
	filter := bson.D{{}}
	planets := []model.Planet{}
	res := client.Database(db).Collection(collection)
	cur, findError := res.Find(context.TODO(), filter)
	if findError != nil {
		return planets, findError
	}
	for cur.Next(context.TODO()) {
		t := model.Planet{}
		err := cur.Decode(&t)
		if err != nil {
			return planets, err
		}
		planets = append(planets, t)
	}
	cur.Close(context.TODO())
	if len(planets) == 0 {
		return planets, mongo.ErrNoDocuments
	}
	return planets, nil

}
func Delete(client *mongo.Client, db, collection string, id primitive.ObjectID) error {
	filter := bson.M{"_id": id}
	res := client.Database(db).Collection(collection)
	_, err := res.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil
}
