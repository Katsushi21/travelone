FROM golang:1.18.0-alpine3.15 as dev

ENV ROOT=/src/app
ENV CGO_ENABLED 0
WORKDIR ${ROOT}
COPY go.mod ./

# hadolint ignore=DL3018
RUN apk update && \
  apk --no-cache add git && \
  go mod tidy && \
  go install github.com/cosmtrek/air@v1.29.0

CMD ["air", "-c", ".air.toml"]