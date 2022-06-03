package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Find(ctx context.Context, coll *mongo.Collection ,filter primitive.D) {
	opts := options.Find().SetSort(bson.D{{"roll",1}})

	cur,err :=coll.Find(ctx,filter,opts)
	if err!=nil {log.Fatalln("❌ Find(): ",err)}

	var res []bson.M

	if err := cur.All(ctx,&res); err !=nil {
		log.Fatalln("❌ Find() Cur.All() : ",err)
	}

	fmt.Println("🚀✨ Find successfull & values ...")

	for _,r := range res {
		fmt.Println("type : ",r)
	}


}


func FindOne(ctx context.Context, coll *mongo.Collection ,filter primitive.D) {

	opts := options.FindOne().SetSort(bson.D{{"roll",-1}})

	var res bson.M 
	err := coll.FindOne(ctx,filter,opts).Decode(&res)
	if err!= nil {
		log.Fatalln("❌ findOne : ",err)
	}
	fmt.Println("🚀✨ FindOne successful & result: ",res)
}

func FindOneAndDelete(ctx context.Context, coll *mongo.Collection ,filter primitive.D) {
	opts:= options.FindOneAndDelete().SetProjection(bson.D{{"_id",1},{"name",1}})

	var deletedDocument bson.D
	err:= coll.FindOneAndDelete(ctx,filter,opts).Decode(&deletedDocument)
	if err!=nil {
		log.Fatalln("❌ FindOneAndDelete() ")
	}
	fmt.Println("🚀✨ FindOneAndDelete successful and deleteDocument : ",deletedDocument)

}
func FindOneAndReplace(ctx context.Context, coll *mongo.Collection ,filter primitive.D,replacement primitive.D) {
    
	opts:= options.FindOneAndReplace().SetUpsert(true)
	var replaceDocument bson.D
	err:=coll.FindOneAndReplace(ctx,filter,replacement,opts).Decode(&replaceDocument)
	if err !=nil{
		log.Fatalln("❌ FindOneAndReplace() : ",err)
	}
	fmt.Println("🚀✨ ~ file: 8.Find.go ~ line 64 ~ funcFindOneAndReplace ~ replaceDocument : ", replaceDocument)
	

}
func FindOneAndUpdate(ctx context.Context, coll *mongo.Collection ,filter primitive.D,update primitive.D) {
	opts:= options.FindOneAndUpdate().SetUpsert(true)
	var updateDocument bson.D
	err:=coll.FindOneAndUpdate(ctx,filter,update,opts).Decode(&updateDocument)
	if err !=nil{
		log.Fatalln("❌ FindOneAndUpdate() : ",err)
	}
	fmt.Println("🚀✨ FindOneAndUpdate & update doc: ", updateDocument)
	
}

