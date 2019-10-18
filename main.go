package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"gitlab.com/Sab94/fabelio-product-price-monitor/services/priceMonitorpb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

type Product struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Url       string             `json:"url" bson:"url"`
	Images    []string           `json:"images" bson:"images"`
	History   []PriceTime        `json:"history" bson:"history"`
	CreatedAt string             `json:"created_at" bson:"created_at"`
}

type PriceTime struct {
	Price string `json:"price" bson:"price"`
	Time  string `json:"time" bson:"time"`
}

type server struct{}

func (*server) AddProduct(ctx context.Context, req *priceMonitorpb.AddProductRequest) (*priceMonitorpb.AddProductResponse, error) {
	url := req.GetUrl()
	fmt.Println(url)
	// Step 1: Fetch from website

	// Step 2: Store into db

	// Step 3: Start a goroutine to ferch updates every hour
	product := Product{
		Url:    "asd",
		Images: []string{"asdasd"},
		History: []PriceTime{
			{
				Price: "132",
				Time:  time.Now().String(),
			},
		},
		CreatedAt: time.Now().String(),
	}

	history := make([]*priceMonitorpb.PriceTime, len(product.History))

	for i := range product.History {
		history[i] = &priceMonitorpb.PriceTime{Price: product.History[i].Price, Time: product.History[i].Time}
	}

	productpb := priceMonitorpb.Product{
		Id:        "qwe",
		Url:       product.Url,
		History:   history,
		Images:    product.Images,
		CreatedAt: product.CreatedAt,
	}
	res := &priceMonitorpb.AddProductResponse{
		Product: &productpb,
	}
	return res, nil
}

func (*server) GetProduct(ctx context.Context, req *priceMonitorpb.GetProductRequest) (*priceMonitorpb.ProductResponse, error) {
	url := req.GetId()
	fmt.Println(url)
	// Step 1: Fetch from db

	// Step 2: Fetch into website (current price)
	product := Product{
		Url:    "asd",
		Images: []string{"asdasd"},
		History: []PriceTime{
			{
				Price: "132",
				Time:  time.Now().String(),
			},
		},
		CreatedAt: time.Now().String(),
	}

	history := make([]*priceMonitorpb.PriceTime, len(product.History))

	for i := range product.History {
		history[i] = &priceMonitorpb.PriceTime{Price: product.History[i].Price, Time: product.History[i].Time}
	}

	productpb := priceMonitorpb.Product{
		Id:        "qwe",
		Url:       product.Url,
		History:   history,
		Images:    product.Images,
		CreatedAt: product.CreatedAt,
	}
	res := &priceMonitorpb.ProductResponse{
		Product: &productpb,
	}
	return res, nil
}

func (*server) GetProducts(ctx context.Context, req *priceMonitorpb.GetProductsRequest) (*priceMonitorpb.ProductsResponse, error) {

	// Step 1: Fetch from db
	product := Product{
		Url:    "asd",
		Images: []string{"asdasd"},
		History: []PriceTime{
			{
				Price: "132",
				Time:  time.Now().String(),
			},
		},
		CreatedAt: time.Now().String(),
	}

	history := make([]*priceMonitorpb.PriceTime, len(product.History))

	for i := range product.History {
		history[i] = &priceMonitorpb.PriceTime{Price: product.History[i].Price, Time: product.History[i].Time}
	}

	productpb := priceMonitorpb.Product{
		Id:        "qwe",
		Url:       product.Url,
		History:   history,
		Images:    product.Images,
		CreatedAt: product.CreatedAt,
	}
	res := &priceMonitorpb.ProductsResponse{
		Products: append([]*priceMonitorpb.Product{}, &productpb),
	}
	return res, nil
}

func main() {
	grpcServer := grpc.NewServer()
	priceMonitorpb.RegisterPriceMonitorServiceServer(grpcServer, &server{})

	// grpc
	listen, err := net.Listen("tcp", ":50051")

	if err != nil {
		grpclog.Fatalf("failed starting grpc server: %v", err)
	}

	go func() {
		fmt.Println("grpc server running on port 50051")
		if err := grpcServer.Serve(listen); err != nil {
			grpclog.Fatalf("failed starting grpc server: %v", err)
		}
	}()

	// grpc Web
	wrappedServer := grpcweb.WrapServer(grpcServer)
	handler := func(resp http.ResponseWriter, req *http.Request) {
		allowCors(resp, req)
		wrappedServer.ServeHTTP(resp, req)
	}
	httpServer := http.Server{
		Addr:    ":8000",
		Handler: http.HandlerFunc(handler),
	}
	go func() {
		fmt.Println("http server running on port 8000")
		if err := httpServer.ListenAndServe(); err != nil {
			grpclog.Fatalf("failed starting http server: %v", err)
		}
	}()

	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig

}

func allowCors(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Expose-Headers", "grpc-status, grpc-message")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, XMLHttpRequest, x-user-agent, x-grpc-web, grpc-status, grpc-message")
}
