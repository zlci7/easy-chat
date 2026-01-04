#!/bin/bash

# Easy-Chat 服务状态查看脚本

PROJECT_ROOT="/root/project/easy-chat"
PID_DIR="$PROJECT_ROOT/pids"

# 颜色输出
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

echo -e "${BLUE}========================================${NC}"
echo -e "${BLUE}  Easy-Chat 服务状态${NC}"
echo -e "${BLUE}========================================${NC}"
echo ""

# 检查服务状态的函数
check_service() {
    local name=$1
    local port=$2
    local pid_file="$PID_DIR/$name.pid"
    
    printf "%-15s" "$name"
    
    if [ -f "$pid_file" ]; then
        local pid=$(cat "$pid_file")
        
        if ps -p $pid > /dev/null 2>&1; then
            echo -e "${GREEN}运行中${NC}  PID: $pid  端口: $port"
        else
            echo -e "${RED}已停止${NC}  (PID 文件存在但进程不存在)"
        fi
    else
        echo -e "${YELLOW}未启动${NC}"
    fi
}

echo -e "${BLUE}RPC 服务：${NC}"
check_service "user-rpc" "8080"
check_service "msg-rpc" "8090"
echo ""

echo -e "${BLUE}API 服务：${NC}"
check_service "user-api" "8888"
check_service "msg-api" "8091"
echo ""

echo -e "${BLUE}网关服务：${NC}"
check_service "gateway" "8889"
echo ""

echo -e "${BLUE}========================================${NC}"
echo ""
echo -e "${YELLOW}提示：${NC}"
echo -e "  启动服务: ./start-all.sh"
echo -e "  停止服务: ./stop-all.sh"
echo -e "  查看日志: tail -f logs/<服务名>.log"
echo ""

