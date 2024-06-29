FROM golang:1.22

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

ENV GOOS=linux
ENV GOARCH=amd64

COPY . .


RUN go build -o runner ./cmd/web
RUN chmod +x /app/runner

EXPOSE 8080

CMD ["./runner"]