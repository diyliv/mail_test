FROM golang:1.18.5-alpine 

WORKDIR /build 

COPY go.mod .
COPY go.sum .

RUN go mod download 

COPY . .

RUN go build -o main cmd/api/main.go

EXPOSE 8080

ENTRYPOINT ["/build/main"]