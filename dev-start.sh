#!/bin/bash

# Easy-Chat 开发环境启动脚本
# 特点：前台运行，Ctrl+C 停止所有服务

set -e

PROJECT_ROOT="/root/project/easy-chat"
LOG_DIR="$PROJECT_ROOT/logs"

# 创建日志目录
mkdir -p "$LOG_DIR"

# 颜色输出
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

# 存储所有服务的 PID
declare -a SERVICE_PIDS

# 清理函数：停止所有服务
cleanup() {
    echo ""
    echo -e "${YELLOW}========================================${NC}"
    echo -e "${YELLOW}  收到停止信号，正在关闭所有服务...${NC}"
    echo -e "${YELLOW}========================================${NC}"
    echo ""
    
    # 杀掉所有子进程
    for pid in "${SERVICE_PIDS[@]}"; do
        if ps -p $pid > /dev/null 2>&1; then
            echo -e "${YELLOW}[停止] PID: $pid${NC}"
            kill $pid 2>/dev/null || true
        fi
    done
    
    # 等待所有进程结束
    sleep 2
    
    # 强制杀掉还没停止的进程
    for pid in "${SERVICE_PIDS[@]}"; do
        if ps -p $pid > /dev/null 2>&1; then
            echo -e "${RED}[强制停止] PID: $pid${NC}"
            kill -9 $pid 2>/dev/null || true
        fi
    done
    
    echo ""
    echo -e "${GREEN}✓ 所有服务已停止${NC}"
    echo ""
    exit 0
}

# 捕获 Ctrl+C (SIGINT) 和 SIGTERM 信号
trap cleanup SIGINT SIGTERM

echo -e "${BLUE}========================================${NC}"
echo -e "${BLUE}  Easy-Chat 开发环境启动${NC}"
echo -e "${BLUE}========================================${NC}"
echo ""
echo -e "${YELLOW}提示：按 Ctrl+C 停止所有服务${NC}"
echo ""

# 启动服务的函数
start_service() {
    local name=$1
    local dir=$2
    local config=$3
    local main_file=$4
    
    echo -e "${BLUE}[启动] $name${NC}"
    
    cd "$PROJECT_ROOT/$dir"
    go run "$main_file" -f "$config" > "$LOG_DIR/$name.log" 2>&1 &
    local pid=$!
    SERVICE_PIDS+=($pid)
    
    echo -e "${GREEN}✓ $name 已启动 (PID: $pid)${NC}"
    echo ""
    
    sleep 1
}

# 按顺序启动所有服务
echo -e "${BLUE}[1/5] 启动 User RPC${NC}"
start_service "user-rpc" "apps/user/rpc" "etc/user.yaml" "user.go"

echo -e "${BLUE}[2/5] 启动 Msg RPC${NC}"
start_service "msg-rpc" "apps/msg/rpc" "etc/msg.yaml" "msg.go"

echo -e "${BLUE}[3/5] 启动 User API${NC}"
start_service "user-api" "apps/user/api" "etc/user-api.yaml" "user.go"

echo -e "${BLUE}[4/5] 启动 Msg API${NC}"
start_service "msg-api" "apps/msg/api" "etc/msg-api.yaml" "msg.go"

echo -e "${BLUE}[5/5] 启动 Gateway${NC}"
start_service "gateway" "apps/gateway" "etc/gateway.yaml" "gateway.go"

echo ""
echo -e "${GREEN}========================================${NC}"
echo -e "${GREEN}  所有服务启动完成！${NC}"
echo -e "${GREEN}========================================${NC}"
echo ""
echo -e "${BLUE}服务列表：${NC}"
echo -e "  • User RPC    : 0.0.0.0:8080"
echo -e "  • Msg RPC     : 0.0.0.0:8090"
echo -e "  • User API    : 0.0.0.0:8888"
echo -e "  • Msg API     : 0.0.0.0:8091"
echo -e "  • Gateway     : 0.0.0.0:8889"
echo ""
echo -e "${YELLOW}日志位置：${NC}"
echo -e "  $LOG_DIR/<服务名>.log"
echo ""
echo -e "${YELLOW}按 Ctrl+C 停止所有服务...${NC}"
echo ""

# 保持脚本运行，直到收到 Ctrl+C
while true; do
    # 检查所有服务是否还在运行
    all_running=true
    for pid in "${SERVICE_PIDS[@]}"; do
        if ! ps -p $pid > /dev/null 2>&1; then
            echo -e "${RED}[警告] 服务 PID $pid 已意外停止${NC}"
            all_running=false
        fi
    done
    
    if [ "$all_running" = false ]; then
        echo -e "${RED}检测到服务异常退出，正在停止所有服务...${NC}"
        cleanup
    fi
    
    sleep 5
done

