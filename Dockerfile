FROM golang:1.26-alpine AS build

WORKDIR /app

COPY go.mod go.sum vendor ./

RUN go mod download && go mod verify

COPY . .

RUN apk --no-cache gcc

RUN CGO_ENABLED=1 go build -o /app/gin-basics .

FROM alpine:latest AS runtime

RUN apk --no-cache add ca-certificates tzdata

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

COPY --from=build /app/gin-basics /gin-basics

USER appuser

EXPOSE 8080

ENTRYPOINT ["/gin-basics"]
