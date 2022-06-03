package main

import (
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

func CreateDatabase(client *mongo.Client) *mongo.Database {
	// if database not exist it will create else will use existing database 
	database := client.Database("testing")
	fmt.Println("🚀✨ Database created successfully")
	return database

}