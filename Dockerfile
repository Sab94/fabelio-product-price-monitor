FROM golang as builder

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GODS=linux GOARCH=amd64 go build
FROM ubuntu

COPY --from=builder /app/fabelio-product-price-monitor /app/

EXPOSE 50051
EXPOSE 50052

ENTRYPOINT ["/app/fabelio-product-price-monitor"]
