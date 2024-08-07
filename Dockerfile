# syntax=docker/dockerfile:1

FROM golang:1.22

WORKDIR /app

COPY go.mod  ./

RUN go mod download

COPY . .

RUN go build -o /website

EXPOSE 8080

CMD [ "/website" ]
