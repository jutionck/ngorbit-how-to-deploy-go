FROM golang:alpine

WORKDIR /app

COPY . .
COPY .env /app

RUN go mod tidy
RUN go build -o blog-app

ENTRYPOINT ["/app/blog-app"]
