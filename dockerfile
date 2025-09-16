# 使用官方 Go 镜像作为构建环境
FROM golang:1.22 AS builder

WORKDIR /app

# 复制代码并构建
COPY . .
RUN go mod tidy && go build -o go-ci-demo .

# 使用轻量级运行时镜像
FROM debian:bullseye-slim

WORKDIR /app

COPY --from=builder /app/go-ci-demo /app/

ENTRYPOINT ["/app/go-ci-demo"]

