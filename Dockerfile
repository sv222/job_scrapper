FROM golang:1.17 as builder
WORKDIR /app
COPY . .
RUN go mod download && CGO_ENABLED=0 GOOS=linux go build -a -o app ./cmd/app/main.go

FROM alpine:3.15.0
WORKDIR /app
COPY --from=builder /app/app .
CMD ["/app/app"]