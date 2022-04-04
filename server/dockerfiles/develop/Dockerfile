FROM golang:1.18.0-alpine3.15 as dev

ENV ROOT=/src/app
ENV CGO_ENABLED 0
WORKDIR ${ROOT}

RUN apk update
COPY go.mod ./
RUN go mod tidy
EXPOSE 8000

CMD ["go", "run", "main.go"]