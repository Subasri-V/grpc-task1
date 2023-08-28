package main

import (
	"context"
	"sync"

	ts "grpc-task1/Task1"
	"grpc-task1/config"

	"go.mongodb.org/mongo-driver/mongo"
)

type server struct {
	mu        sync.Mutex
	customers map[string]*ts.CustDetails
	ts.UnimplementedBankserviceServer
}

type database struct{
	collection *mongo.Collection
	client *mongo.Client
}

func (s *server) InsertCustomer(ctx context.Context, req *ts.CustDetails) (*ts.Response, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	client, err := config.ConnectDataBase()
	if err != nil {
		return nil, err
	}
	defer client.Disconnect(ctx)

	collection := client.Database("grpcCust").Collection("grpcCustCollection")

	_, err = collection.InsertOne(ctx, req)
  if err != nil {
    return nil, err

	return &ts.Response{req.Name,int32(req.Balance),[]byte(req.BankId),req.Accountid},nil

}
}
