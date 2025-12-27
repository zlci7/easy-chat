# Easy Chat 配置指南

## 环境变量配置

复制以下内容创建 `.env` 文件（不要提交到 Git）：

```bash
# 数据库配置
MYSQL_ROOT_PASSWORD=your_password_here
MYSQL_DATABASE=easy_chat

# Redis 配置
REDIS_PASSWORD=your_redis_password_here

# Etcd 配置
ETCD_LISTEN_CLIENT_URLS=http://0.0.0.0:2379
ETCD_ADVERTISE_CLIENT_URLS=http://etcd:2379

# 阿里云镜像仓库
ALIYUN_REGISTRY=crpi-xxxxx.cn-hangzhou.personal.cr.aliyuncs.com
ALIYUN_NAMESPACE=easy-chat7

# 应用配置
APP_ENV=development
APP_PORT=8080
```

## 本地开发配置

1. **创建数据目录**
```bash
mkdir -p components/{mysql,redis,etcd}/{data,logs}
mkdir -p components/redis/config
```

2. **配置 Redis**（可选）
```bash
touch components/redis/config/redis.conf
```

## 生产环境注意事项

⚠️ **重要：生产环境必须修改默认密码！**

- MySQL root 密码
- Redis 密码
- 修改默认端口（可选）

