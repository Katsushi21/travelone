FROM golang:1.18.0-alpine3.15 AS builder
WORKDIR /src/app
ENV CGO_ENABLED 0
ENV GOOS linux

COPY go.mod .
COPY go.sum .
RUN  go mod download

COPY . .

RUN go build -o ./binary .

FROM alpine:3.15 AS runner
WORKDIR /src/app

COPY --from=builder /src/app/binary .

RUN addgroup --system --gid 1001 golang && \
  adduser --system --uid 1001 golang
USER golang:golang

CMD ["/src/app/binary"]