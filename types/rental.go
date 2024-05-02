package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type CarToRental struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name   string             `bson:"name" json:"name"`
	Rating int                `bson:"rating" json:"rating"`
}

type Car struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	CarBodyType string             `bson:"carBodyType" json:"carBodyType"`
	Prise       float64            `bson:"prise" json:"prise"`
}
