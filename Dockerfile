
FROM golang:1.24-alpine AS base
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod tidy
COPY . .

FROM base AS backend-builder
RUN go build -o main ./cmd/backend/main.go

FROM base AS publisher-builder
RUN go build -o mqtt_publisher ./cmd/script_mqtt_publisher/mqtt_publisher.go

FROM base AS worker-builder
RUN go build -o geofence_worker ./cmd/geofence_worker/rabbitmq_listener.go

FROM alpine:latest AS backend
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=backend-builder /app/main .
COPY --from=backend-builder /app/.env .
EXPOSE 3000
CMD ["./main"]

FROM alpine:latest AS publisher
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=publisher-builder /app/mqtt_publisher .
COPY --from=publisher-builder /app/.env .
CMD ["./mqtt_publisher"]

FROM alpine:latest AS worker
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=worker-builder /app/geofence_worker .
COPY --from=worker-builder /app/.env .
CMD ["./geofence_worker"]
