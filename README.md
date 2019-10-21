# fabelio-product-price-monitor

# Stack : go, angular, mongodb

# To convert `.proto` files check `protoscripts` file

## Run the server 
```
go mod download
go run main.go
```

## Build 
```
go mod download
go build
```

## gRPC services 
    - priceMonitorpb : saves and shows products
    - crawlerpb : crawls provided links

### `priceMonitorpb` runs on  port 50051 and `crawlerpb` runs on  port 50052 also `priceMonitorpb` is exposed to http on port 8080

### adding one product url stores it in mongodb and starts a cron to save price changes hourly


## Run the frontend
```
cd web-client
ng serve
```
## Build frontend
```
cd web-client
ng build --prod
```