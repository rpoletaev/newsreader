FROM golang:1.13-alpine as builder

RUN apk update && apk add --no-cache git alpine-sdk make
WORKDIR /newsreader
COPY . .

RUN make build

FROM alpine

WORKDIR /app
COPY --from=builder /orders/bin .

ENTRYPOINT ["./cmd"]