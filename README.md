# Easy Chat

> 🚧 项目开发中 - Demo 阶段

基于 Go-Zero 的微服务即时通讯系统学习项目。

## 技术栈

- **框架**: Go-Zero
- **语言**: Go 1.21+
- **数据库**: MySQL 8.0
- **缓存**: Redis 7.2
- **服务发现**: Etcd v3.5.10
- **容器化**: Docker & Docker Compose

## 快速开始

### 开发模式（推荐）

**1. 启动基础设施服务**
```bash
docker-compose up -d
```

**2. 本地运行 user-rpc 服务**
```bash
go run apps/user/rpc/user.go -f apps/user/rpc/etc/user.yaml
```

**3. 修改配置连接本地基础设施**

确保 `apps/user/rpc/etc/user.yaml` 配置：
```yaml
Etcd:
  Hosts:
  - 127.0.0.1:3379  # 本地 Docker 端口
```

### 部署模式

**一键启动所有服务（包括应用）**
```bash
# 取消注释 docker-compose.yaml 中的 user-rpc 配置
docker-compose up -d
```

## 服务端口

| 服务 | 端口 |
|------|------|
| user-rpc | 18080 |
| etcd | 3379 |
| redis | 16379 |
| mysql | 13306 |

## 常用命令

```bash
# 查看日志
docker-compose logs -f

# 重启服务
docker-compose restart

# 停止服务
docker-compose down

# 构建镜像
make user-rpc-dev
```

## 项目结构

```
easy-chat/
├── apps/           # 应用服务
├── deploy/         # 部署配置
├── components/     # 数据目录
└── Makefile        # 构建脚本
```

## 开发计划

- [x] 基础框架搭建
- [x] Docker 容器化
- [ ] 用户服务
- [ ] 消息服务
- [ ] API 网关
- [ ] 前端界面


