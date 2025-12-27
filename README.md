# Easy Chat

基于 Go-Zero 的微服务即时通讯系统

## 📖 项目简介

Easy Chat 是一个采用微服务架构的即时通讯系统，使用 Go-Zero 框架开发，支持分布式部署。

## ✨ 特性

- 🚀 微服务架构，易于扩展
- 🔐 用户认证与授权
- 💬 实时消息推送
- 📦 Docker 容器化部署
- 🎯 服务注册与发现（Etcd）
- 💾 数据持久化（MySQL + Redis）

## 🏗️ 技术栈

- **框架**: Go-Zero
- **语言**: Go 1.21+
- **服务发现**: Etcd v3.5.10
- **数据库**: MySQL 8.0
- **缓存**: Redis 7.2
- **容器**: Docker & Docker Compose
- **镜像仓库**: 阿里云容器镜像服务

## 📁 项目结构

```
easy-chat/
├── apps/                   # 应用服务
│   └── user/              # 用户服务
│       └── rpc/           # RPC 服务
│           ├── etc/       # 配置文件
│           ├── internal/  # 内部逻辑
│           └── user.go    # 服务入口
├── deploy/                # 部署相关
│   ├── dockerfile/        # Dockerfile 文件
│   └── mk/               # Makefile 脚本
├── components/            # 组件数据目录
│   ├── etcd/
│   ├── redis/
│   └── mysql/
├── docker-compose.yaml    # Docker Compose 配置
└── Makefile              # 主构建文件
```

## 🚀 快速开始

### 前置要求

- Go 1.21+
- Docker & Docker Compose
- Make

### 本地开发

1. **克隆项目**
```bash
git clone <your-repo-url>
cd easy-chat
```

2. **安装依赖**
```bash
go mod tidy
```

3. **启动基础设施服务**
```bash
docker-compose up -d etcd redis mysql
```

4. **运行服务（本地）**
```bash
go run apps/user/rpc/user.go -f apps/user/rpc/etc/user.yaml
```

### Docker 部署

1. **构建并推送镜像**
```bash
make user-rpc-dev
```

2. **启动所有服务**
```bash
docker-compose up -d
```

3. **查看服务状态**
```bash
docker-compose ps
```

## 🔧 配置说明

### 环境变量

- `MYSQL_ROOT_PASSWORD`: MySQL root 密码（默认: 1234）
- `REDIS_PASSWORD`: Redis 密码（默认: 1234）

### 端口映射

| 服务 | 容器端口 | 宿主机端口 |
|------|---------|-----------|
| user-rpc | 8080 | 18080 |
| etcd | 2379 | 3379 |
| redis | 6379 | 16379 |
| mysql | 3306 | 13306 |

## 📊 服务管理

### 查看日志
```bash
# 所有服务
docker-compose logs -f

# 特定服务
docker-compose logs -f user-rpc
```

### 重启服务
```bash
# 重启特定服务
docker-compose restart user-rpc

# 重启所有服务
docker-compose restart
```

### 停止服务
```bash
# 停止所有服务
docker-compose down

# 停止并删除数据卷
docker-compose down -v
```

## 🔨 开发指南

### 添加新的 RPC 服务

1. 在 `apps/` 下创建新服务目录
2. 编写服务代码
3. 创建对应的 Dockerfile
4. 在 `deploy/mk/` 下创建构建脚本
5. 更新 `docker-compose.yaml`

### 构建流程

```bash
# 编译
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bin/user-rpc ./apps/user/rpc/user.go

# 构建镜像
docker build -f deploy/dockerfile/Dockerfile_user_rpc_dev -t easy-im-user-rpc-test .

# 推送镜像
make user-rpc-dev
```

## 🐛 故障排查

### iptables 错误
如果遇到 iptables 相关错误：
```bash
sudo iptables -t filter -N DOCKER-ISOLATION-STAGE-1
sudo iptables -t filter -N DOCKER-ISOLATION-STAGE-2
sudo systemctl restart docker
```

### 镜像拉取失败
配置使用华为云镜像加速，已在 docker-compose.yaml 中配置。

## 📄 License

MIT License

## 🤝 贡献

欢迎提交 Issue 和 Pull Request！

## 📮 联系方式

- Email: your-email@example.com
- GitHub: [@your-username](https://github.com/your-username)

