FROM golang:1.16-alpine as builder

RUN mkdir /build
WORKDIR /build

COPY . .

RUN export GO111MODULE=on

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o url-parser ./cmd/web

FROM alpine:latest

WORKDIR /app
COPY --from=builder /build/ui ./ui
COPY --from=builder /build/url-parser ./url-parser

CMD ./url-parser