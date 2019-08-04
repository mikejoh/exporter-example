FROM golang:1.12 AS builder
RUN mkdir /service-api-exporter
ADD go.* /service-api-exporter/
ADD *.go /service-api-exporter/
WORKDIR /service-api-exporter
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o service-api-exporter .

FROM scratch
COPY --from=builder /service-api-exporter/service-api-exporter .
ENTRYPOINT ["./service-api-exporter"]