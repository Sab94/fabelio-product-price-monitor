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
	"github.com/robfig/cron/v3"
	"gitlab.com/Sab94/fabelio-product-price-monitor/crawlerClient"
	"gitlab.com/Sab94/fabelio-product-price-monitor/database"
	"gitlab.com/Sab94/fabelio-product-price-monitor/services/crawlerpb"
	"gitlab.com/Sab94/fabelio-product-price-monitor/services/priceMonitorpb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

type Product struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Url         string             `json:"url" bson:"url"`
	Name        string             `json:"name" bson:"name"`
	Images      string             `json:"images" bson:"images"`
	History     []PriceTime        `json:"history" bson:"history"`
	CreatedAt   string             `json:"created_at" bson:"created_at"`
	CronID      cron.EntryID       `json:"cron_id" bson:"cron_id"`
	Description string             `json:"description" bson:"description"`
	Price       string             `json:"price" bson:"price"`
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
	collection := database.DB.Collection("products")
	cronEntryID, err := c.AddFunc("@every 2m", func() {
		product := Product{}
		err := collection.FindOne(context.Background(), bson.M{"url": url}).Decode(&product)
		if err != nil {
			log.Fatal(err)
		}

		p := getProduct(url)
		pt := PriceTime{
			Price: p.GetPrice(),
			Time:  time.Now().String(),
		}
		product.History = append(product.History, pt)
		_, err = collection.UpdateOne(context.Background(), bson.M{"_id": product.ID}, bson.D{
			{"$set", bson.D{
				{"history", product.History},
				{"price", pt.Price},
			}},
		})

		if err != nil {
			log.Fatal(err)
		}

	})
	if err != nil {
		log.Fatal(err)
	}
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
		Name:        productRaw.GetName(),
		CreatedAt:   time.Now().String(),
		CronID:      cronEntryID,
		Description: productRaw.GetDescription(),
		Price:       productRaw.GetPrice(),
	}

	_, err = collection.InsertOne(ctx, product)
	if err != nil {
		log.Fatal(err)
	}

	history := make([]*priceMonitorpb.PriceTime, len(product.History))

	for i := range product.History {
		history[i] = &priceMonitorpb.PriceTime{Price: product.History[i].Price, Time: product.History[i].Time}
	}

	productpb := priceMonitorpb.Product{
		Id:          "aaa",
		Url:         product.Url,
		History:     history,
		Images:      product.Images,
		CreatedAt:   product.CreatedAt,
		Description: product.Description,
		Price:       product.Price,
	}
	res := &priceMonitorpb.AddProductResponse{
		Product: &productpb,
	}
	return res, nil
}

func (*server) GetProduct(ctx context.Context, req *priceMonitorpb.GetProductRequest) (*priceMonitorpb.ProductResponse, error) {
	productID, _ := primitive.ObjectIDFromHex(req.GetId())
	fmt.Println(productID)
	// Step 1: Fetch from db
	collection := database.DB.Collection("products")
	product := Product{}
	err := collection.FindOne(context.Background(), bson.M{"_id": productID}).Decode(&product)
	if err != nil {
		log.Fatal(err)
	}

	// Step 2: Fetch into website (current price)
	history := make([]*priceMonitorpb.PriceTime, len(product.History))

	for i := range product.History {
		history[i] = &priceMonitorpb.PriceTime{Price: product.History[i].Price, Time: product.History[i].Time}
	}

	productpb := priceMonitorpb.Product{
		Id:          product.ID.Hex(),
		Url:         product.Url,
		History:     history,
		Images:      product.Images,
		CreatedAt:   product.CreatedAt,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
	}
	res := &priceMonitorpb.ProductResponse{
		Product: &productpb,
	}
	return res, nil
}

func (*server) GetProducts(ctx context.Context, req *priceMonitorpb.GetProductsRequest) (*priceMonitorpb.ProductsResponse, error) {

	// Step 1: Fetch from db
	_context := context.Background()
	collection := database.DB.Collection("products")

	cur, err := collection.Find(_context, bson.D{}, nil)

	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(_context)

	productpbs := []*priceMonitorpb.Product{}
	for cur.Next(_context) {
		result := Product{}
		err := cur.Decode(&result)
		productpb := &priceMonitorpb.Product{
			Id:          result.ID.Hex(),
			Url:         result.Url,
			Images:      result.Images,
			CreatedAt:   result.CreatedAt,
			Name:        result.Name,
			Description: result.Description,
			Price:       result.Price,
		}
		productpbs = append(productpbs, productpb)
		if err != nil {
			log.Fatal(err)
		}
	}

	res := &priceMonitorpb.ProductsResponse{
		Products: productpbs,
	}
	return res, nil
}

func (*server) Crawl(ctx context.Context, req *crawlerpb.ProductUrl) (*crawlerpb.ProductInfo, error) {
	url := req.GetUrl()

	product := getProduct(url)

	res := &product

	return res, nil
}

var c = cron.New()

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	} else {
		fmt.Println("env loaded")
	}
	database.Connect()
	c.Start()
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
		} else if name, _ := s.Attr("property"); name == "og:description" {
			product.Description, _ = s.Attr("content")
		}
	})

	return product
}
