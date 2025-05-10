FROM golang:1.24.3-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o project-service ./cmd/server/main.go

FROM alpine:3.18

WORKDIR /app

RUN apk --no-cache add ca-certificates
COPY --from=builder /app/project-service .

EXPOSE 50053

CMD ["./project-service", "server"]