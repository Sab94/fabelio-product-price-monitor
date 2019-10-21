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
		Url: "https://fabelio.com/ip/dabi-sofa-bed.html",
	}

	res, err := c.AddProduct(context.Background(), req)

	if err != nil {
		// log.Fatalf("error while calling %v", err)
	}
	if res.GetProduct().GetPrice() == "" {
		log.Println("blank string")
	}
	log.Println(res.GetProduct().GetPrice())
}
