package main

import (
	"google.golang.org/grpc"
	"log"
)

const (
	address = "localhost:9090"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	//c := iampb.NewIamClient(conn)
	//
	//// Contact the server and print out its response.
	//
	//meta := make(map[string]interface{})
	//body := make(map[string]interface{})
	//meta["name"] = "ohoh"
	//body["time"] = "wowo"
	//
	//requestMeta, err := structpb.NewStruct(meta) // Check to rules below to avoid errors
	//if err != nil {
	//	panic(err)
	//}
	//requestBody, err := structpb.NewStruct(body) // Check to rules below to avoid errors
	//if err != nil {
	//	panic(err)
	//}
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//defer cancel()
	//log.Printf("Greeting: %s", r.GetStatusCode())
}