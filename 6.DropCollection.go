package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)


func DropCollection(ctx context.Context,coll *mongo.Collection ){
	

	err:= coll.Drop(ctx)

	if err!=nil {log.Fatalln("❌ : ",err)}

	fmt.Println("🚀✨ Collection dropped successfully")

}