FROM golang:1.22-alpine AS build
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o app .

FROM alpine
WORKDIR /
COPY --from=build /app/app .

CMD ["/app"]
