# Build stage
FROM golang:1.23rc2-alpine3.20 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go
RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz

# Run stage
FROM alpine:3.20
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/migrate.linux-amd64 ./migrate
COPY app.env .
COPY start.sh .
COPY wait-for.sh .
COPY db/migration ./migration

EXPOSE 8080
CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]
#docker run --name simplebank --network bank-network -p 8080:8080 -e GIN_MODE=release -e DB_SOURCE="postgresql://root:secret@postgres12:5432/simple_bank?sslmode=disable" simplebank:latest