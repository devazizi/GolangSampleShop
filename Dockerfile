# build app
FROM golang:1.22-alpine AS builder
# RUN apk add --no-cache git gcc libc-dev
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go mod download
RUN go build -o app

# final container app
From alpine:latest
WORKDIR /app
COPY --from=builder /app /app
ENTRYPOINT ./app
EXPOSE 3000
