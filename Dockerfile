FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o main .

FROM alpine
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 3000
CMD ["./main"]