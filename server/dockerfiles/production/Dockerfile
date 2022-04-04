FROM golang:1.18.0-alpine3.15 as dev

ENV ROOT=/src/app
ENV CGO_ENABLED 0
WORKDIR ${ROOT}

RUN apk update
COPY go.mod ./
RUN go mod tidy
EXPOSE 8000

CMD ["go", "run", "main.go"]


FROM golang:1.18.0-alpine3.15 as builder

ENV ROOT=/src/app
WORKDIR ${ROOT}

RUN apk update
COPY go.mod ./
RUN go mod tidy

COPY . ${ROOT}
RUN CGO_ENABLED=0 GOOS=linux go build -o $ROOT/binary


FROM scratch as prod

ENV ROOT=/src/app
WORKDIR ${ROOT}
COPY --from=builder ${ROOT}/binary ${ROOT}

EXPOSE 8000
CMD ["/src/app/binary"]