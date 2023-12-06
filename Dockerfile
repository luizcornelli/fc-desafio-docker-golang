FROM golang:latest AS build

WORKDIR /app

COPY . .

RUN go build -o desafio .

FROM scratch

WORKDIR /

COPY --from=build /app/desafio /desafio

ENTRYPOINT [ "/desafio" ]
