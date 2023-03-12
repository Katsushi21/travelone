FROM golang:1.18.0-alpine3.15 AS dev
WORKDIR /src/auth
ENV CGO_ENABLED 0

COPY go.mod .
RUN apk update && \
  apk --no-cache add git=2.34.7-r0 && \
  go mod tidy && \
  go install github.com/cosmtrek/air@v1.41.0

COPY . .

CMD ["air", "-c", ".air.toml"]