package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func InsertOne(ctx context.Context,collection *mongo.Collection,data *primitive.D){
	result, err := collection.InsertOne(ctx, *data)
	if err !=nil {log.Fatalln("❌ ",err)}

	fmt.Println("🚀✨ insert One Data successfully : ",result)

}




func InsertMany(ctx context.Context,coll *mongo.Collection,docs []interface{}){
	opts:=options.InsertMany().SetOrdered(false)

	res,err :=coll.InsertMany(ctx,docs,opts)
	if err!= nil {log.Fatalln("❌ : ",err)}

	fmt.Println("🚀✨ InsertMany successful & res.InseredIDs : ",res.InsertedIDs)

}


