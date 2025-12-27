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

### 1. 克隆项目
```bash
git clone <your-repo-url>
cd easy-chat
```

### 2. 启动服务
```bash
docker-compose up -d
```

### 3. 查看状态
```bash
docker-compose ps
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


