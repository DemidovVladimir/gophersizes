FROM golang AS builder

ENV GO111MODULE=on

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./...

FROM alpine:latest AS production

COPY --from=builder /app .

EXPOSE 8080

ENTRYPOINT ["./main"]
