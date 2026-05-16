#!/bin/bash
# Fork 构建：子路径 /mmw/ 前端 + forkserver 后端
set -e

ROOT="$(cd "$(dirname "$0")/.." && pwd)"
cd "$ROOT"

echo "========================================"
echo "Fork 构建（BASE_PATH 默认 /mmw）"
echo "========================================"

BUILD_DIR="build/fork"
FRONTEND_DIR="miaomiaowu"

echo ""
echo "[1/2] 构建前端（base=/mmw/）..."
cd "$FRONTEND_DIR"
if [ ! -d "node_modules" ]; then
  npm install
fi
npx vite build --config vite.config.fork.mts
cd "$ROOT"
echo "前端构建完成 ✓"

echo ""
echo "[2/2] 构建 forkserver..."
mkdir -p "$BUILD_DIR"
CGO_ENABLED=1 go build -trimpath -ldflags="-s -w" -o "${BUILD_DIR}/forkserver" ./fork/cmd/forkserver
echo "后端构建完成 ✓"

echo ""
echo "输出: ${BUILD_DIR}/forkserver"
echo "默认 BASE_PATH=/mmw，Nginx 使用 location /mmw/"
