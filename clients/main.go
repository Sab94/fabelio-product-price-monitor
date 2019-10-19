package main

import (
	"context"
	"fmt"
	"log"

	"gitlab.com/Sab94/fabelio-product-price-monitor/services/priceMonitorpb"
	"google.golang.org/grpc"
)

func main() {

	fmt.Printf("priceMonitorpb")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect %v", err)
	}
	defer cc.Close()

	c := priceMonitorpb.NewPriceMonitorServiceClient(cc)

	req := &priceMonitorpb.AddProductRequest{
		Url: "https://fabelio.com/ip/tirai-kuro.html",
	}

	res, err := c.AddProduct(context.Background(), req)

	if err != nil {
		log.Fatalf("error while calling watcher %v", err)
	}
	log.Println(res)
}
