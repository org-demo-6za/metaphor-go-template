# todo: add build stages, make image smaller, run it with less privileges
FROM golang:1.18-alpine3.16

WORKDIR /opt/app

COPY . .

RUN cd cmd/http && go build -o /opt/app/metaphor-api

CMD ["./metaphor-api"]