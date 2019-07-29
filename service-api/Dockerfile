FROM golang:1.12 AS builder
RUN mkdir /service-api
ADD go.* /service-api/
ADD *.go /service-api/
WORKDIR /service-api
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o service-api .

FROM scratch
COPY --from=builder /service-api/service-api .
ENTRYPOINT ["./service-api"]