package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/joho/godotenv"
	"gitlab.com/Sab94/fabelio-product-price-monitor/crawlerClient"
	"gitlab.com/Sab94/fabelio-product-price-monitor/database"
	"gitlab.com/Sab94/fabelio-product-price-monitor/services/crawlerpb"
	"gitlab.com/Sab94/fabelio-product-price-monitor/services/priceMonitorpb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

type Product struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Url       string             `json:"url" bson:"url"`
	Name      string             `json:"name" bson:"name"`
	Images    string             `json:"images" bson:"images"`
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
	// Step 1: Fetch from website
	productRaw := crawlerClient.Crawl(url)

	// Step 2: Start a goroutine to ferch updates every hour

	// Step 3: Store into db
	product := Product{
		Url:    url,
		Images: productRaw.GetImage(),
		History: []PriceTime{
			{
				Price: productRaw.GetPrice(),
				Time:  time.Now().String(),
			},
		},
		Name:      productRaw.GetName(),
		CreatedAt: time.Now().String(),
	}
	collection := database.DB.Collection("products")
	_, err := collection.InsertOne(ctx, product)
	if err != nil {
		log.Fatal(err)
	}

	history := make([]*priceMonitorpb.PriceTime, len(product.History))

	for i := range product.History {
		history[i] = &priceMonitorpb.PriceTime{Price: product.History[i].Price, Time: product.History[i].Time}
	}

	productpb := priceMonitorpb.Product{
		Id:        "aaa",
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
		Images: "asdasd",
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
		Images: "asdasd",
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

func (*server) Crawl(ctx context.Context, req *crawlerpb.ProductUrl) (*crawlerpb.ProductInfo, error) {
	url := req.GetUrl()

	product := getProduct(url)

	res := &product

	return res, nil
}

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	} else {
		fmt.Println("env loaded")
	}
	database.Connect()

	priceMonitorServer := grpc.NewServer()
	crawlerServer := grpc.NewServer()
	priceMonitorpb.RegisterPriceMonitorServiceServer(priceMonitorServer, &server{})
	crawlerpb.RegisterCrawlerServiceServer(crawlerServer, &server{})

	// grpc
	listenPriceMonitor, err := net.Listen("tcp", ":50051")
	listenCrawl, err := net.Listen("tcp", ":50052")

	if err != nil {
		grpclog.Fatalf("failed starting grpc servers: %v", err)
	}

	go func() {
		fmt.Println("PriceMonitor server running on port 50051")
		if err := priceMonitorServer.Serve(listenPriceMonitor); err != nil {
			grpclog.Fatalf("failed starting PriceMonitor server: %v", err)
		}
	}()

	go func() {
		fmt.Println("Crawl server running on port 50052")
		if err := crawlerServer.Serve(listenCrawl); err != nil {
			grpclog.Fatalf("failed starting Crawl server: %v", err)
		}
	}()

	// grpc Web
	wrappedServer := grpcweb.WrapServer(priceMonitorServer)
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

func parseMap(aMap map[string]interface{}) {
	for key, val := range aMap {
		switch concreteVal := val.(type) {
		case map[string]interface{}:
			parseMap(val.(map[string]interface{}))
		case []interface{}:
			parseArray(val.([]interface{}))
		default:
			if key == "price" {
				fmt.Println(key, ":", concreteVal)
			}
		}
	}
}

func parseArray(anArray []interface{}) {
	for _, val := range anArray {
		switch concreteVal := val.(type) {
		case map[string]interface{}:
			parseMap(val.(map[string]interface{}))
		case []interface{}:
			parseArray(val.([]interface{}))
		default:
			_ = concreteVal
		}
	}
}

func getProduct(url string) crawlerpb.ProductInfo {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}
	product := crawlerpb.ProductInfo{}
	doc.Find("meta").Each(func(i int, s *goquery.Selection) {
		if name, _ := s.Attr("property"); name == "og:title" {
			product.Name, _ = s.Attr("content")
		} else if name, _ := s.Attr("property"); name == "og:image" {
			product.Image, _ = s.Attr("content")
		} else if name, _ := s.Attr("property"); name == "product:price:amount" {
			product.Price, _ = s.Attr("content")
		}
	})

	return product
}
