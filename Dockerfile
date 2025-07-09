# syntax=docker/dockerfile:1

FROM golang:1.24

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o . ./cmd/eis

EXPOSE 8080

CMD ["./eis"]
