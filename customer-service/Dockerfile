FROM golang:1.22.5 as builder

WORKDIR /go/src/customer-service
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o customer-service .

FROM alpine:latest as final

ENV CONNECTION_STRING="host=host.docker.internal user=gorm password=gorm dbname=gorm port=5432 sslmode=disable TimeZone=America/Los_Angeles"
ENV PORT=8080

WORKDIR /go/src/customer-service
COPY --from=builder /go/src/customer-service/customer-service .

EXPOSE 8080

CMD ["./customer-service"]
