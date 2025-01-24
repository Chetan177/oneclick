package db

import (
	"context"
	"time"


	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/chetan177/oneclick/models"
)

const (
	dbName            = "oneclick"
	systemCollection  = "system"
	projectCollection = "project"
	apiCollection     = "api"

	tokenName = "api_token"
)

type DB struct {
	Client *mongo.Client
}

func NewDB(uri string) (*DB, error) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	return &DB{Client: client}, nil
}

func (db *DB) Close() {
	if db.Client != nil {
		db.Client.Disconnect(context.TODO())
	}
}

func (db *DB) GetSystemData() (models.SystemData, error) {
	collection := db.Client.Database("oneclick").Collection("system")
	var systemData models.SystemData
	err := collection.FindOne(context.TODO(), bson.M{"name": tokenName}).Decode(&systemData)
	return systemData, err
}

func (db *DB) CreateOrUpdateSystemData(token string) error {
	collection := db.Client.Database("oneclick").Collection("system")
	filter := bson.M{"name": tokenName}
	update := bson.M{"$set": bson.M{"token": token}}
	opts := options.Update().SetUpsert(true)
	_, err := collection.UpdateOne(context.TODO(), filter, update, opts)
	return err
}

func (db *DB) CreateOrGetToken(token string) (string, error) {
	collection := db.Client.Database("oneclick").Collection("system")
	filter := bson.M{"name": tokenName}
	update := bson.M{"$setOnInsert": bson.M{"token": token, "created_at": time.Now(), "updated_at": time.Now()}} 
	opts := options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After)

	var systemData models.SystemData
	err := collection.FindOneAndUpdate(context.TODO(), filter, update, opts).Decode(&systemData)
	if err != nil {
		return "", err
	}

	return systemData.Token, nil
}

