FROM golang:1.22

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

ENV GOOS=linux
ENV GOARCH=amd64

RUN go build ./cmd/web

RUN chmod +x ./web

RUN ls -la

COPY . .

EXPOSE 8080

CMD ["./web"]