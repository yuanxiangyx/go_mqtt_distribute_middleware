FROM golang:alpine AS builder

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build
COPY . .

RUN go build -o run .

FROM scratch
WORKDIR /app
COPY config.json .

COPY --from=builder /build/run .

CMD ["/app/run"]
