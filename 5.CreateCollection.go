package main

import (
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

func CreateCollection(database *mongo.Database) *mongo.Collection {

	// if collection created if not exits else use existing collection
	collection := database.Collection("mycollection")

	fmt.Println("🚀✨ Collection created successfully")
	return collection

}