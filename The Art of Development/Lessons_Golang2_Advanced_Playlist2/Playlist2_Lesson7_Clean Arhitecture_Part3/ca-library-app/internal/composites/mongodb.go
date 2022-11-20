package composites

import (
	"ca-library-app/pkg/client/mongodb"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBComposite struct {
	db *mongo.Database
}

func NewMongoDBComposite(ctx context.Context, Host, Port, UserName, Password, Database, AuthDB string) (*MongoDBComposite, error) {
	client, err := mongodb.NewClient(ctx, Host, Port, UserName, Password, Database, AuthDB)
	if err != nil {
		return nil, err
	}
	return &MongoDBComposite{db: client}, nil
}
