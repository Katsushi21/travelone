FROM golang:1.18.0-alpine3.15 as dev

ENV ROOT=/src/app
ENV CGO_ENABLED 0
WORKDIR ${ROOT}

COPY go.mod ${ROOT}
RUN apk update && \
  apk --no-cache add git=2.34.2-r0 && \
  go mod tidy && \
  go install github.com/cosmtrek/air@v1.29.0

COPY . ${ROOT}

CMD ["air", "-c", ".air.toml"]