package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type Database struct {
	Db       *mongo.Database
	Host     string
	Username string
	Password string
	AppName  string
	DbName   string
}

func (d *Database) Connect() {
	// Create the connection URI

	url := fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority&appName=%s", d.Username, d.Password, d.Host, d.AppName)

	// Use the SetServerAPIOptions() method to set the version of the Stable API on the client
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(url).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Ping the database to ensure a successful connection
	if err := client.Ping(context.TODO(), nil); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	// Assign the connected database to the struct
	d.Db = client.Database(d.DbName)
	fmt.Println("Successfully connected to MongoDB!")
}

func (d *Database) Disconnect() {
	if err := d.Db.Client().Disconnect(context.TODO()); err != nil {
		log.Fatalf("Failed to disconnect from MongoDB: %v", err)
	}
	fmt.Println("Successfully disconnected from MongoDB!")
}
func (d *Database) PrintUsers() {
	// Get the user collection
	collection := d.Db.Collection("users")

	// Find all documents
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatalf("Failed to find documents: %v", err)
	}
	defer cursor.Close(context.TODO())

	// Iterate through the cursor and print documents
	for cursor.Next(context.TODO()) {
		var user bson.M
		if err = cursor.Decode(&user); err != nil {
			log.Fatalf("Failed to decode document: %v", err)
		}
		fmt.Println(user)
	}

	if err := cursor.Err(); err != nil {
		log.Fatalf("Cursor error: %v", err)
	}
}
