FROM golang:1.23.3-alpine as builder

WORKDIR /app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go test ./... -v
RUN go build -o server

FROM alpine

WORKDIR /app
COPY --from=builder /app/server /bin/server
COPY --from=builder /app/config.toml .

EXPOSE 8080

CMD ["server"]
