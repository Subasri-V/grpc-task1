package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Customer struct{
	Account_id primitive.ObjectID `json:"accountid" bson:"accountid"`
	Name string `json:"name" bson:"name"`
	Bank_id string `json:"bankid" bson:"accountid"`
}