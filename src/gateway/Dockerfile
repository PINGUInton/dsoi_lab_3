FROM golang:1.25.1

WORKDIR /gateway

COPY ./ ./

RUN go mod tidy
RUN go mod vendor
RUN go build -o gateway ./cmd/main.go

EXPOSE 8080
ENTRYPOINT ["./gateway"]
