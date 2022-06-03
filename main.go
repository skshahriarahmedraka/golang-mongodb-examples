package main

import (
	"context"
	// "fmt"
	// "log"
	"time"

	// "time"

	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo/readpref"
)

// Connection URI
// const uri = "mongodb://admin:password@localhost:27017/?maxPoolSize=20&w=majority"

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	// CONNECT TO MONGODB 
	client :=ConnectionToMongodb(ctx)

	// PING TO MONGODB
	Ping(ctx,client)

	// DROP COLLECTION
	DropCollection(ctx,client.Database("testing").Collection("mycollection"))
	// DROP DATABASE
	DropDatabase(ctx,client.Database("testing"))
	// CREATE DATABASE
	database:=CreateDatabase(client)

	// CREATE COLLECTION
	collection:=CreateCollection(database)

	// INSERT INSIDE COLLECTION 
	user := bson.D{{"_id","raka"},{"fullName", "skshahriar ahmed raka"}, {"age", 30},{"os",99}}
	u2:= bson.D{
		{"_id","skraka"},
		{"UserID","skraka"},
		{"UserName","Sk Shahriar Ahmed Raka"},
		{"UserTitle","Working at Google Inc."},
		{"UserImage","https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060640/samples/jpeg_1_qlbtcn.jpg"},
		{"Badges",bson.D{{"Gold",999},{"Silver",888},{"Bronze",777}}},
		{"Location","Dhaka, Bangladesh"},
		{"MembershipTime","3 year 5 Month"},
		{"LastSeen","This Week"},
		{"Aboutme","A Curious Learner, Full-Stack Web Developer, Security Researcher\nHere are my skills and strengths:\n✓ Expert in Golang\n ✓ Expert in Fiber framework (using Golang) \n ✓ Expert in WebAssembly (using Golang)  Expert in Golang     ✓ Expert in Fiber framework (using Golang)    ✓ Expert in WebAssembly (using Golang) ✓ Expert in database design, development, optimization, and migration    (PostgreSQL, MySQL, MongoDB , Redis) ✓ Expert in ( Grpc, protocol buffer ) ✓ Experienced in ( WebSocket, Socket.io, WebRTC ) for real-time client and server applications ✓ Experienced in ( Svelte.js ) and some knowledge in ( TypeScript ) ✓ Good understanding of ( Docker, Bash, PowerShell, Git,    Nginx, Kubernetes )        Github : github.com/skshahriarahmedraka    Upwork : upwork.com/o/profiles/users/~0107ef3405bffbe57e    Linkedin : linkedin.com/in/sk-shahriar-ahmed-raka-862a31193/  Telegram : t.me/shahriarraka \", \"MySite\":\"https://stackoverflow.com/users/12216779/sk-shahriar-ahmed-raka"},
		{"Mysite",""},
		{"Github","https://github.com/skshahriarahmedraka"},
		{"Twitter","https://twitter.com/shahriarraka7"},
		{"TopTags",bson.A{"go","rust","python","svelte","postgresql"}},
		{"TopTagsPercent",bson.A{50,30,20,5,2}},
		{"SelectedPanel","Profile"}}

	InsertOne(ctx,collection,&user)
	InsertOne(ctx,collection,&u2)


	// INSERT MANY 
	docs:= []interface{}{
		bson.D{{"_id","insertMany1"},{"name","skshahriar"},{"roll",234}},
		bson.D{{"_id","insertMany2"},{"name","RealMadrid"},{"roll",453}},
		bson.D{{"_id","insertMany3"},{"name","liverpool"},{"roll",99}},
		bson.D{{"_id","insertMany4"},{"name","mancity"},{"roll",4}},
		bson.D{{"_id","insertMany5"},{"name","newcacle"},{"roll",33}},
		bson.D{{"_id","insertMany6"},{"name","Wolves"},{"roll",3}},
		bson.D{{"_id","insertMany7"},{"name","weee"},{"roll",7343},{"birthday",34}},
		bson.D{{"_id","insertMany8"},{"name","myyy"},{"roll",8363},{"birthday",4}},
		bson.D{{"_id","insertMany9"},{"name","youuu"},{"roll",7643},{"birthday",544}},
		bson.D{{"_id","insertMany10"},{"name","theee"},{"roll",9043},{"birthday",4}},
	}
	InsertMany(ctx,collection,docs)


	// FIND 
	filter:= bson.D{{"roll",bson.D{{"$gt",5}}}}
	Find(ctx,collection,filter)

	// FIND ONE 
	filter= bson.D{{"roll",bson.D{{"$lt",100}}}}
	FindOne(ctx,collection,filter)


	// FIND ONE AND DELETE
	filter= bson.D{{"roll",bson.D{{"$eq",33}}}}
	FindOneAndDelete(ctx,collection,filter)
	
	// FIND ONE AND REPLACE
	filter= bson.D{{"roll",bson.D{{"$eq",3}}}}
	replacement:= 	bson.D{{"_id","insertMany6"},{"name","Wolves FC"},{"roll",38}}
	FindOneAndReplace(ctx,collection,filter,replacement)

	// FIND AND UPDATE 
	filter= bson.D{{"name",bson.D{{"$eq","skshahriar"}}}}
	update := bson.D{{"$set",bson.D{{"name","raka"}}}}
	FindOneAndUpdate(ctx,collection,filter,update)

	// UPDATE ONE

	filter =bson.D{{"roll",453}}
	update =bson.D{{"$set",bson.D{{"name","MyMadrid"}}}} 
	UpdateOne(ctx,collection,filter,update)

	// UPDATE MANY
	filter= bson.D{{"roll",bson.D{{"$gt",999}}}}
	today:= time.Now().Format("2006-01-02 3:4:5 pm")
	update = bson.D{{"$set",bson.D{{"birthday",today}}}}
	UpdateMany(ctx,collection,filter,update)




	// DELETE ONE 
	filter= bson.D{{"fullName", "skshahriar ahmed raka"}}
	DeleteOne(ctx,*collection,filter)

	



	// client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	// if err != nil { 
	// 	log.Fatalln("error : ",err)
	// }
	// Create a new client and connect to the server
	// client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	// if err != nil {
	// 	panic(err)
	// }

	// Check the connection
	// err = client.Ping(ctx, nil)

	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Mongodb connected successfull")


	// usersCollection := client.Database("testing").Collection("users")
	// user := bson.D{{"fullName", "skshahriar ahmed raka"}, {"age", 30},{"os",99}}
	// // insert the bson object using InsertOne()
	// result, err := usersCollection.InsertOne(ctx, user)
	// // check for errors in the insertion
	// if err != nil {
	// 	panic(err)
	// }
	// // display the id of the newly inserted object
	// fmt.Println(result)
}



