package crawlerClient

import (
	"context"
	"log"

	"gitlab.com/Sab94/fabelio-product-price-monitor/services/crawlerpb"
	"google.golang.org/grpc"
)

func Crawl(url string) *crawlerpb.ProductInfo {
	cc, err := grpc.Dial("localhost:50052", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect %v", err)
	}
	defer cc.Close()

	c := crawlerpb.NewCrawlerServiceClient(cc)

	req := &crawlerpb.ProductUrl{
		Url: url,
	}

	res, err := c.Crawl(context.Background(), req)

	if err != nil {
		log.Fatalf("error while calling watcher %v", err)
	}

	return res
}
