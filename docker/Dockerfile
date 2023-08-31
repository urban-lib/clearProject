FROM golang:1.21-alpine3.18 AS builder

ARG SERVICE

RUN mkdir /internal
COPY . /internal
WORKDIR /internal/$SERVICE

RUN apk add --no-cache git

RUN go build -o /app .

FROM alpine:latest

RUN apk add -U tzdata

ENV TZ="UTC"

WORKDIR /root
COPY --from=builder /internal/$SERVICE .

CMD ["/app"]