package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func UpdateOne(ctx context.Context, coll *mongo.Collection, filter primitive.D, update primitive.D) {

	opts:= options.Update().SetUpsert(true)
	res,err:= coll.UpdateOne(ctx,filter,update,opts)

	if err!=nil {
		log.Fatalln("❌ UpdateOne : ",err)
	}

	if res.MatchedCount!= 0{
		fmt.Println("🚀✨matched and replaced exited document & result : ",res)	
		return
 	}
	if res.UpsertedCount!= 0{
		fmt.Println("🚀✨ docummt is not found so created a new document & documentID : ",res.UpsertedID)	
	}

}

func UpdateMany(ctx context.Context, coll *mongo.Collection, filter primitive.D, update primitive.D) {

	opts:= options.Update().SetUpsert(true)
	res,err := coll.UpdateMany(ctx,filter,update,opts)
	if err !=nil {
		log.Fatalln("❌ UpdateMany : ",err)
	}
	if res.MatchedCount!= 0{
		fmt.Println("🚀✨ matched and replaced existing document")
		return 
	}else {
		fmt.Println("🚀✨ UpdateMany Successfull & result : ",res)
	}
	
}

func UpdateById() {

}

func ReplaceByOne() {

}