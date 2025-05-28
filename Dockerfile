FROM golang:1.22-alpine AS build

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o desafio .

FROM alpine:latest

WORKDIR /

COPY --from=build /app/desafio /desafio

RUN apk --no-cache add ca-certificates

EXPOSE 8081

ENTRYPOINT ["/desafio"]
