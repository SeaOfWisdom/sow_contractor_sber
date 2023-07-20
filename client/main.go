package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/SeaOfWisdom/sow_proto/contractor-srv"

	"google.golang.org/grpc"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := grpc.Dial("localhost:5305", []grpc.DialOption{grpc.WithInsecure()}...)
	failOnError(err, "Geo-service. While connetion client")
	defer conn.Close()

	cli := pb.NewContractorServiceClient(conn)

	reqG := pb.PurchaseWorkRequest{
		ReaderAddress: "xxxx",
		WorkId:        "1",
		AuthorAddress: "Yyyyy",
		Price:         "5000",
	}
	responseG, err := cli.PurchaseWork(context.Background(), &reqG)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(responseG)
}
