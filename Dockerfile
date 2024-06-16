FROM golang:1.22

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

ENV GOOS=linux
ENV GOARCH=amd64

RUN go build -o main ./cmd/web

RUN chmod +x /app/main

COPY . .

EXPOSE 8080

CMD ["./main"]