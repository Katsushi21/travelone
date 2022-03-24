FROM golang:1.18.0-alpine3.15 as dev

ENV ROOT=/go/src/app
ENV CGO_ENABLED 0
WORKDIR ${ROOT}

RUN apk update
COPY go.mod go.sum ./
RUN go mod download
EXPOSE 8000

CMD ["go", "run", "main.go"]


FROM golang:1.18.0-alpine3.15 as builder

ENV ROOT=/go/src/app
WORKDIR ${ROOT}

RUN apk update
COPY go.mod go.sum ./
RUN go mod download

COPY . ${ROOT}
RUN CGO_ENABLED=0 GOOS=linux go build -o $ROOT/binary


FROM scratch as prod

ENV ROOT=/go/src/app
WORKDIR ${ROOT}
COPY --from=builder ${ROOT}/binary ${ROOT}

EXPOSE 8000
CMD ["/go/src/app/binary"]