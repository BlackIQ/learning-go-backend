# BUILD Stage
FROM golang:1.26-trixie AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -mod=readonly -o /app/gin-basics .


# PRODUCTION Stage
FROM debian:trixie-slim AS production

WORKDIR /app

RUN mkdir -p /app/data

COPY --from=build /app/gin-basics /app/gin-basics

EXPOSE 8080

ENV PORT=8080
ENV DB_PATH=/app/data/go.db

ENTRYPOINT ["/app/gin-basics"]
