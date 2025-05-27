FROM golang:latest AS build

WORKDIR /app
COPY . .
RUN go build -o desafio .

FROM scratch
WORKDIR /
COPY --from=build /app/desafio /desafio

EXPOSE 8081  # A porta que seu app Go realmente usa
ENTRYPOINT [ "/desafio" ]
