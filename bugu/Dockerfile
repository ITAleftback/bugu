FROM golang:1.15

ENV GO115MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY="https://goproxy.cn,direct"

WORKDIR /home/go/src/bugu

COPY . .
RUN go build -o app .
EXPOSE 8000
CMD ["./app"]