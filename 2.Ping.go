package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

func Ping(ctx context.Context,client *mongo.Client){
	err := client.Ping(ctx, nil)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("🚀✨ Mongodb connected successfully & ping successfully")
}