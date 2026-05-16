#!/bin/bash
# Fork 构建：子路径 /mmw/ 前端 + forkserver 后端
set -e

ROOT="$(cd "$(dirname "$0")/.." && pwd)"
cd "$ROOT"

echo "========================================"
echo "Fork 构建（BASE_PATH 默认 /mmw）"
echo "========================================"

if [ ! -f "$ROOT/.env" ]; then
  echo ""
  echo "提示: 未找到 .env，将使用 Vite 内置默认 VITE_BASE_PATH=/mmw/"
  echo "      自定义子路径请: cp .env.example .env  并修改 BASE_PATH / VITE_BASE_PATH"
fi

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
echo "配置: 见 .env.example（cp 为 .env 后生效）；默认 BASE_PATH=/mmw，Nginx location /mmw/"
