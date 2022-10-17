package db

import (
	"context"
	"log"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/sirupsen/logrus"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dbName = "roava"
)

type MongoDB struct {
	Client *mongo.Client
}

func connect() *mongo.Client {
	godotenv.Load()
	url := os.Getenv("MONGO_URL")

	if url == "" {
		url = "mongodb://localhost:27017"
	}

	client, _ := mongo.NewClient(options.Client().ApplyURI(url))

	ctx, _ := context.WithTimeout(context.Background(), 50*time.Second)
	err := client.Connect(ctx)

	if err != nil {
		log.Println("db connection error", err)
		return nil
	}

	err = client.Ping(nil, nil)
	if err != nil {
		log.Println("db connection error", err)
		return nil
	}
	logrus.Println("db connected")

	return client
}

// create a connection to mongodb
func NewMongoConnection() *MongoDB {
	conn := &MongoDB{Client: connect()}

	return conn
}

func (db MongoDB) CreateTTLIndex(coll *mongo.Collection, expir *int32) error {
	index := mongo.IndexModel{
		Keys:    bson.D{{"created_at", 1}},
		Options: &options.IndexOptions{ExpireAfterSeconds: expir},
	}
	_, err := coll.Indexes().CreateOne(nil, index)

	if err != nil {
		return err
	}

	return nil
}

// returns flag for duplicate error
func (db MongoDB) IsMongoDuplicateError(err error) bool {
	if strings.Contains(err.Error(), "E11000") {
		return true
	}
	return false
}

func (db MongoDB) DropCollection(collection string) {
	db.Client.Database(dbName).Collection(collection).Drop(nil)
}

func (db MongoDB) AccountCollection() *mongo.Collection {
	return db.Client.Database(dbName).Collection("accounts")
}
