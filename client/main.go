package main

import (
	"fmt"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"time"

	//"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	user "Galaraz/gRPC-mongo/proto/gen"
)

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":8200", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	usr := user.NewUserServiceClient(conn)

	fmt.Print("Enter the UserName: ")
	var inputUser string
	fmt.Scanln(&inputUser)

	response, err := usr.GetUser(context.Background(), &user.UserRequest{Username: inputUser})
	if err != nil {
		log.Fatalf("Error when calling GetUser: %s", err)
	}

	log.Printf("Response from server: %v", response)

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		fmt.Println(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		fmt.Println(err)
	}

	db := client.Database("gitUsers")
	collection := db.Collection("users")

	res, err := collection.InsertOne(context.Background(), response)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.InsertedID.(primitive.ObjectID).Timestamp())

}
