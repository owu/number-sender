FROM golang:1.25 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
# 优化构建：添加-lsflags="-w -s"去除调试信息，减小二进制文件大小
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o main .

FROM alpine:3.19
# 只安装必要的包
RUN apk --no-cache add ca-certificates tzdata
RUN mkdir -p /home/www/apps/number-sender-go/config /home/www/logs/number-sender-go
WORKDIR /home/www/apps/number-sender-go
# 从builder阶段复制编译好的二进制文件
COPY --from=builder /app/main /home/www/apps/number-sender-go/main
# 复制配置文件
COPY --from=builder /app/config/config.toml /home/www/apps/number-sender-go/config/config.toml
RUN chmod +x /home/www/apps/number-sender-go/main
CMD ["/home/www/apps/number-sender-go/main", "--config=/home/www/apps/number-sender-go/config/config.toml"]