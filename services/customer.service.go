package services

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type service struct {
	ctx        context.Context
	collection *mongo.Collection
}


func InsertCustomer()