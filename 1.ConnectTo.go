package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func ConnectionToMongodb(ctx context.Context) *mongo.Client{
	//LOAD ENVIRONMENT VARIABLE 
	err:= godotenv.Load()
	if err !=nil {log.Fatalln("❌ ",err)}



	client, err := mongo.Connect(ctx, options.Client().ApplyURI(fmt.Sprint("mongodb://",os.Getenv("MONGO_USER"),":",os.Getenv("MONGO_PASSWORD"),"@",os.Getenv("MONGO_HOSTIP"),":",os.Getenv("MONGO_PORT"),"/?maxPoolSize=20&w=majority")))
	if err != nil {log.Fatalln("❌ error : ",err)}

	fmt.Println("🚀✨ Mongodb Connected successfully")
	return client
}