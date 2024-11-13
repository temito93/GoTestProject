package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"goproject/internal/config"
	"goproject/internal/database/index"
	"log"
	"sync"
	"time"
)

type MongoDBConnection struct {
	Client *mongo.Client
	DB     *mongo.Database
}

var (
	globalDBConnection *MongoDBConnection
	once               sync.Once
)

// ConnectDB initializes the global database connection
func ConnectDB() error {
	var err error
	once.Do(func() {
		cfg := config.MongoConfig()

		if cfg.URI == "" {
			err = fmt.Errorf("MONGODB_URI environment variable not set")
			return
		}
		if cfg.Name == "" {
			err = fmt.Errorf("MONGO_DB environment variable not set")
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.URI))
		if err != nil {
			log.Printf("Failed to connect to MongoDB: %v", err)
			return
		}

		if err := client.Ping(ctx, nil); err != nil {
			log.Printf("Failed to ping MongoDB: %v", err)
			return
		}

		db := client.Database(cfg.Name)

		if err := index.CreateUserIndex(db); err != nil {
			log.Printf("Failed to create user index: %v", err)
			return
		}

		log.Println("Connected to MongoDB")
		globalDBConnection = &MongoDBConnection{Client: client, DB: db}
	})
	return err
}

// GetDBConnection returns the global database connection
func GetDBConnection() *MongoDBConnection {
	return globalDBConnection
}

func (conn *MongoDBConnection) Disconnect() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := conn.Client.Disconnect(ctx); err != nil {
		log.Printf("Error disconnecting from MongoDB: %v", err)
	} else {
		log.Println("Disconnected from MongoDB")
	}
}
