FROM golang:latest AS build

WORKDIR /app
COPY . .
RUN go build -o desafio .

FROM alpine:latest
WORKDIR /
COPY --from=build /app/desafio /desafio

EXPOSE 8081
ENTRYPOINT ["/desafio"]
