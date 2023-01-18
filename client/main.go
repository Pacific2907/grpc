package main

import (
	"context"
	"log"
	"time"

	pb "example.com/grpc/proto"
	"google.golang.org/grpc"
)

const (
	address = "localhost:9000"
)

func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("connection Failed: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserManagementClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	var new_users = make(map[string]int32)
	new_users["Prasant"] = 23
	new_users["Rohit"] = 20
	new_users["Mohan"] = 24
	for name, age := range new_users {
		r, err := c.CreateNewUser(ctx, &pb.NewUser{Name: name, Age: age})
		if err != nil {
			log.Fatalf("could not create user: %v", err)
		}
		log.Printf(`User Data:
NAME: %s
AGE: %d
ID: %d`, r.GetName(), r.GetAge(), r.GetId())

	}
}
