FROM golang:1.24.3-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o ./bin/project-service ./cmd/server/main.go
RUN go build -o ./bin/project-client ./cmd/client/main.go

FROM alpine:3.18 AS bin

RUN apk --no-cache add ca-certificates tzdata

ENV TZ=Asia/Tokyo

COPY --from=builder /app/bin/project-service /usr/local/bin/project-service
COPY --from=builder /app/bin/project-client /usr/local/bin/project-client

EXPOSE 50053

ENTRYPOINT ["project-service"]
CMD ["server"]