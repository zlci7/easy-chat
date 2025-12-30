# 多阶段构建：第一阶段编译
FROM golang:alpine AS builder

WORKDIR /app

# 配置 Go 模块代理（加速）
ENV GOPROXY=https://goproxy.cn,direct
ENV GO111MODULE=on
ENV GOTOOLCHAIN=auto

# 复制 go.mod 和 go.sum
COPY go.mod go.sum ./
RUN go mod download

# 复制整个项目
COPY . .

# 编译 User RPC
WORKDIR /app/apps/user/rpc
RUN CGO_ENABLED=0 GOOS=linux go build -o user-rpc .

# 第二阶段：运行
FROM alpine:latest

WORKDIR /app

# 配置阿里云软件源
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
    apk --no-cache add ca-certificates tzdata

# 从构建阶段复制可执行文件
COPY --from=builder /app/apps/user/rpc/user-rpc .
COPY --from=builder /app/apps/user/rpc/etc/user-docker.yaml ./etc/user.yaml

EXPOSE 8080

CMD ["./user-rpc", "-f", "./etc/user.yaml"]