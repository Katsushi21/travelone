FROM golang:1.17.5-alpine3.15 AS build

RUN mkdir -p /go/src
WORKDIR /go/src
COPY . .

RUN go build -o server .

FROM alpine:3.15
COPY --from=build /go/src/server /server

CMD ["/server"]
