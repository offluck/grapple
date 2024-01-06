FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN mkdir -p bin/app
RUN go build -o bin/app/grappler github.com/offluck/grapple/cmd/app

FROM alpine:latest AS runner
WORKDIR /app
COPY --from=builder /app/bin/app/grappler .
EXPOSE 8000
CMD ["./grappler"]
