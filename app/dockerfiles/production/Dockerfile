# 依存パッケージのインストール
FROM golang:1.18.0-alpine3.15 as deps

ENV ROOT=/src/app
ENV CGO_ENABLED 0
WORKDIR ${ROOT}

COPY go.mod ${ROOT}
RUN apk update && \
  apk --no-cache add git=2.34.2-r0 && \
  go mod tidy

COPY . ${ROOT}

CMD ["go", "run", "main.go"]

FROM golang:1.18.0-alpine3.15 as builder

ENV ROOT=/src/app
WORKDIR ${ROOT}

COPY go.mod ${ROOT}
RUN apk update && \
  go mod tidy
COPY . ${ROOT}

RUN CGO_ENABLED=0 GOOS=linux go build -o ${ROOT}/binary

FROM builder as runner

ENV ROOT=/src/app
WORKDIR ${ROOT}
COPY --from=builder ${ROOT}/binary ${ROOT}

CMD ["/src/app/binary"]