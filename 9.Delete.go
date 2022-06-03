package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DeleteOne(ctx context.Context,collection mongo.Collection,condition bson.D) {
	opts :=options.Delete().SetCollation(&options.Collation{
		Locale: "en_US",
		Strength: 1,
		CaseLevel: false,
	})

	res,err := collection.DeleteOne(ctx,condition,opts)
	if err !=nil { log.Fatalln("❌ : ",err) }

	fmt.Println("🇵🇸 Delete One successful & res.DeleteCount : ",res.DeletedCount)

}

func DeleteMany(){

}