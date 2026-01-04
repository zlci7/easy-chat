#!/bin/bash

# Easy-Chat 开发环境强制停止脚本
# 用途：当 dev-start.sh 的 Ctrl+C 没有正确清理时使用

# 颜色输出
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m'

echo -e "${YELLOW}正在查找并停止所有 Easy-Chat 服务...${NC}"
echo ""

# 停止服务的函数 - 同时匹配 go run 和编译后的二进制
stop_service() {
    local name=$1
    local pattern=$2
    local binary_name=$3
    
    # 先尝试匹配 go run 启动的进程
    if pkill -f "$pattern" 2>/dev/null; then
        echo -e "${GREEN}✓ $name 已停止 (go run)${NC}"
        return 0
    fi
    
    # 再尝试通过端口找到 PID 并停止
    local stopped=false
    if [ ! -z "$4" ]; then
        local port=$4
        local pids=$(ss -tlnp | grep ":$port " | grep -oP 'pid=\K[0-9]+' | sort -u)
        if [ ! -z "$pids" ]; then
            for pid in $pids; do
                kill $pid 2>/dev/null && stopped=true
            done
        fi
    fi
    
    # 最后尝试通过二进制名称停止
    if pkill -x "$binary_name" 2>/dev/null; then
        stopped=true
    fi
    
    if [ "$stopped" = true ]; then
        echo -e "${GREEN}✓ $name 已停止${NC}"
    else
        echo "  $name 未运行"
    fi
}

# 停止所有服务
stop_service "User RPC" "apps/user/rpc/user.go" "user" "8080"
stop_service "Msg RPC" "apps/msg/rpc/msg.go" "msg" "8090"
stop_service "User API" "apps/user/api/user.go" "user" "8888"
stop_service "Msg API" "apps/msg/api/msg.go" "msg" "8091"
stop_service "Gateway" "apps/gateway/gateway.go" "gateway" "8889"

sleep 2

# 检查端口是否释放
echo ""
echo -e "${YELLOW}检查端口占用情况...${NC}"
if ss -tlnp | grep -qE ':(8080|8090|8888|8091|8889)'; then
    echo -e "${YELLOW}以下端口仍被占用：${NC}"
    ss -tlnp | grep -E ':(8080|8090|8888|8091|8889)'
    echo ""
    
    # 提取所有 PID
    pids=$(ss -tlnp | grep -E ':(8080|8090|8888|8091|8889)' | grep -oP 'pid=\K[0-9]+' | sort -u | tr '\n' ' ')
    
    if [ ! -z "$pids" ]; then
        echo -e "${YELLOW}强制停止这些进程? [y/N]${NC}"
        read -t 10 -n 1 -r response || response="n"
        echo ""
        
        if [[ $response =~ ^[Yy]$ ]]; then
            echo -e "${YELLOW}正在强制停止...${NC}"
            for pid in $pids; do
                kill -9 $pid 2>/dev/null && echo -e "${GREEN}✓ 已停止 PID: $pid${NC}" || echo -e "${RED}✗ 无法停止 PID: $pid${NC}"
            done
            sleep 1
            echo ""
            if ss -tlnp | grep -qE ':(8080|8090|8888|8091|8889)'; then
                echo -e "${RED}部分端口仍被占用${NC}"
                ss -tlnp | grep -E ':(8080|8090|8888|8091|8889)'
            else
                echo -e "${GREEN}✓ 所有端口已释放${NC}"
            fi
        else
            echo -e "${YELLOW}如需手动强制停止，请执行：${NC}"
            echo "  kill -9 $pids"
        fi
    fi
else
    echo -e "${GREEN}✓ 所有端口已释放${NC}"
fi

echo ""