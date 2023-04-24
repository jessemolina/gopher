package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Custom server data type model.
type Person struct {
	Name string `json:"name" bson:"name"`
	Age  int    `json:"age" bson:"age"`
}

func main() {
	// Set the client context, options and connect.
	ctx := context.Background()
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	// Check the connection.
	err = verifyConnection(ctx, client)
	if err != nil {
		log.Fatal(err)
	}

	// Create a collection.
	collection := client.Database("contacts").Collection("people")

	// Insert a Person entry
	person := Person{"Jesse", 18}
	err = createPerson(ctx, collection, person)
	if err != nil {
		log.Fatal(err)
	}

	// Retrieve a Person entry
	filter := bson.M{"name": "Jesse"}
	err = retrievePerson(ctx, collection, filter)
	if err != nil {
		log.Fatal(err)
	}

	// Update a Person entry
	update := bson.M{"name": "Jesse", "age": 21}
	err = updatePerson(ctx, collection, filter, bson.M{"$set": update})
	if err != nil {
		log.Fatal(err)
	}

	// Delete a Person entry
	err = deletePerson(ctx, collection, filter)
	if err != nil {
		log.Fatal(err)
	}
}

// Verify connection to database.
func verifyConnection(ctx context.Context, client *mongo.Client) error {
	err := client.Ping(context.Background(), nil)
	if err != nil {
		return err
	}

	log.Println("Connection verified.")
	return nil
}

// Create Person entry from Person.
func createPerson(ctx context.Context, collection *mongo.Collection, person Person) error {
	result, err := collection.InsertOne(context.Background(), person)
	if err != nil {
		return err
	}

	log.Println(result.InsertedID, "Person entry inserted.")
	return nil
}

// Retrieve a Person entry from filter.
func retrievePerson(ctx context.Context, collection *mongo.Collection, filter bson.M) error {
	var person Person
	err := collection.FindOne(ctx, filter).Decode(&person)
	if err != nil {
		return err
	}
	log.Println("Person entry retrieved")
	fmt.Println(person)
	return nil
}

// Update a Person entry from filter with update.
func updatePerson(ctx context.Context, collection *mongo.Collection, filter bson.M, update bson.M) error {
	results, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	log.Println(results.UpsertedID, "Person entry updated")
	return nil
}

// Delete a Person entry from filter.
func deletePerson(ctx context.Context, collection *mongo.Collection, filter bson.M) error {
	results, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	log.Println(results.DeletedCount, "Person entry deleted")
	return nil
}
