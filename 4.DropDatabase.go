package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

func DropDatabase(ctx context.Context,database *mongo.Database){
	err:= database.Drop(ctx)
	if err!=nil {log.Fatalln("❌ : ",err)}
	fmt.Println("🚀✨ Database droped successfully")
}