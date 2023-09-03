FROM golang:1.21 as base

ENV GO111MODULE=on

FROM base as dev

WORKDIR /app

COPY . .
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-teacher-api

EXPOSE 8080

CMD ["/docker-teacher-api"]