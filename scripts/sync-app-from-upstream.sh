#!/bin/bash
# 合并源仓后：对比 cmd/server/main.go 与 internal/app/run.go，提示需人工同步的区块
set -e

ROOT="$(cd "$(dirname "$0")/.." && pwd)"
cd "$ROOT"

UPSTREAM_MAIN="${1:-}"
if [ -z "$UPSTREAM_MAIN" ]; then
  if [ -f "cmd/server/main.go" ]; then
    echo "用法: $0 [upstream-main.go路径]"
    echo "或将 upstream 的 cmd/server/main.go 保存后作为第一个参数传入。"
    echo ""
    echo "当前对比: cmd/server/main.go（若已瘦身）与 internal/app/run.go"
    if [ ! -f "internal/app/run.go" ]; then
      echo "错误: 未找到 internal/app/run.go"
      exit 1
    fi
    echo "--- 提示：源仓 main.go 的 mux 注册/后台任务变更需同步到 internal/app/run.go ---"
    wc -l cmd/server/main.go internal/app/run.go 2>/dev/null || true
    exit 0
  fi
  exit 1
fi

echo "对比 upstream main 与 internal/app/run.go ..."
diff -u internal/app/run.go "$UPSTREAM_MAIN" || true
echo "请人工将 upstream main.go 中的路由/启动逻辑合并到 internal/app/run.go"
