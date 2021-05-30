package repo

import (
	"reflect"
	"testing"

	cfg "github.com/carlosd-ss/star-wars/pkg/config"

	mcon "github.com/carlosd-ss/star-wars/pkg/mongo"

	"github.com/carlosd-ss/star-wars/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestCreate(t *testing.T) {
	type args struct {
		client     *mongo.Client
		db         string
		collection string
		planet     model.Planet
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "TestCreate", args: args{client: cfg.Connect, db: "starwars", collection: "planet", planet: model.Planet{}}, wantErr: false},
		{name: "TestCreateColection and db nil", args: args{client: cfg.Connect, db: "", collection: "", planet: model.Planet{}}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Create(tt.args.client, tt.args.db, tt.args.collection, tt.args.planet); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	type args struct {
		client     *mongo.Client
		db         string
		collection string
		id         primitive.ObjectID
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "TestDelete", args: args{client: cfg.Connect, db: "starwars", collection: "planet", id: primitive.NewObjectID()}, wantErr: false},
		{name: "TestDelete and db nil", args: args{client: cfg.Connect, db: "", collection: "", id: primitive.NewObjectID()}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Delete(tt.args.client, tt.args.db, tt.args.collection, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestListAll(t *testing.T) {
	type args struct {
		client     *mongo.Client
		db         string
		collection string
	}
	tests := []struct {
		name    string
		args    args
		wantPts []model.Planet
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "TestListAll and db nil", args: args{client: mcon.Connect(), db: "", collection: ""}, wantPts: []model.Planet{}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPts, err := ListAll(tt.args.client, tt.args.db, tt.args.collection)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotPts, tt.wantPts) {
				t.Errorf("ListAll() gotPts = %v, want %v", gotPts, tt.wantPts)
			}
		})
	}
}

func TestSearchForId(t *testing.T) {
	type args struct {
		client     *mongo.Client
		db         string
		collection string
		id         primitive.ObjectID
	}
	tests := []struct {
		name       string
		args       args
		wantPlanet model.Planet
		wantErr    bool
	}{
		// TODO: Add test cases.
		{name: "TestSearchForId and db nil", args: args{client: mcon.Connect(), db: "", collection: ""}, wantPlanet: model.Planet{}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPlanet, err := SearchForId(tt.args.client, tt.args.db, tt.args.collection, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("SearchForId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotPlanet, tt.wantPlanet) {
				t.Errorf("SearchForId() gotPlanet = %v, want %v", gotPlanet, tt.wantPlanet)
			}
		})
	}
}

func TestSearchForName(t *testing.T) {
	type args struct {
		client     *mongo.Client
		db         string
		collection string
		name       string
	}
	tests := []struct {
		name       string
		args       args
		wantPlanet model.Planet
		wantErr    bool
	}{
		// TODO: Add test cases.

		{name: "TestSearchForName and db nil", args: args{client: mcon.Connect(), db: "", collection: ""}, wantPlanet: model.Planet{}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPlanet, err := SearchForName(tt.args.client, tt.args.db, tt.args.collection, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("SearchForName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotPlanet, tt.wantPlanet) {
				t.Errorf("SearchForName() gotPlanet = %v, want %v", gotPlanet, tt.wantPlanet)
			}
		})
	}
}
