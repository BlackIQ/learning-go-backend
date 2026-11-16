FROM golang:1.26-alpine AS build

WORKDIR /app

RUN apk add --no-cache gcc musl-dev

COPY go.mod go.sum ./
COPY vendor ./vendor

RUN CGO_ENABLED=1 go build -mod=vendor -o /app/gin-basics .

FROM alpine:latest

RUN apk add --no-cache ca-certificates tzdata

WORKDIR /app

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

RUN mkdir -p /app/data && chown -R appuser:appgroup /app

COPY --from=build /app/gin-basics /app/gin-basics

USER appuser

EXPOSE 8080

ENV PORT=8080
ENV DB_PATH=/app/data/go.db

ENTRYPOINT ["/app/gin-basics"]
