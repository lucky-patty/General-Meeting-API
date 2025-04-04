package db

import (
  "context"
  "time"

  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
)

type DBClient struct {
  Client *mongo.Client
  DB     *mongo.Database
}


func NewDatabase(uri, dbName string) (*DBClient, error) {
  ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
  defer cancel()

  clientOpts := options.Client().ApplyURI(uri)
  client, err := mongo.Connect(ctx, clientOpts)
  if er != nil {
    return nil, err 
  }

  return &DBClient{
    Client: client,
    DB:     client.Database(dbName),
  }, nil
}

func (db *DBClient) Close(ctx context.Context) error {
  return db.Client.Disconnect(ctx)
}
