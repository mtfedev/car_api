package stores

import (
	"context"

	"github.com/mtfedev/car_api/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RentalStore interface {
	InsertCar(context.Context, *types.CarToRental) (*types.CarToRental, error)
	Update(context.Context, bson.M, bson.M) error
	GetCars(context.Context, bson.M) ([]*types.CarToRental, error)
	GetCarByID(context.Context, primitive.ObjectID) (*types.CarToRental, error)
}

type MongoCarStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func (s *MongoCarStore) Hotel(context.Context, *types.CarToRental) (*types.CarToRental, error) {
	panic("unimplemented")
}

func NewMongoCarStore(client *mongo.Client) *MongoCarStore {
	return &MongoCarStore{
		client: client,
		coll:   client.Database(DBNAME).Collection("cars"),
	}
}

func (s *MongoCarStore) GetCarByID(ctx context.Context, id primitive.ObjectID) (*types.CarToRental, error) {
	var car types.CarToRental
	if err := s.coll.FindOne(ctx, bson.M{"_id": id}).Decode(&car); err != nil {
		return nil, err
	}
	return &car, nil
}

func (s *MongoCarStore) GetCars(ctx context.Context, filter bson.M) ([]*types.CarToRental, error) {
	resp, err := s.coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var cars []*types.CarToRental
	if err := resp.All(ctx, &cars); err != nil {
		return nil, err
	}
	return cars, nil
}

func (s *MongoCarStore) Update(ctx context.Context, filter bson.M, update bson.M) error {
	_, err := s.coll.UpdateOne(ctx, filter, update)
	return err
}

func (s *MongoCarStore) InsertHotel(ctx context.Context, hotel *types.CarToRental) (*types.CarToRental, error) {
	resp, err := s.coll.InsertOne(ctx, hotel)
	if err != nil {
		return nil, err
	}
	hotel.ID = resp.InsertedID.(primitive.ObjectID)
	return hotel, nil
}
