#!/bin/bash
# sync_project.sh
# 自动打包 + 推送 GitHub + 输出路径以供上传给 ChatGPT

# 检查参数
if [ -z "$1" ]; then
  echo "Usage: ./sync_project.sh [go|php]"
  exit 1
fi

# 设置项目名称
if [ "$1" = "go" ]; then
  PROJECT="go_service"
elif [ "$1" = "php" ]; then
  PROJECT="b_service"
else
  echo "Invalid argument. Use 'go' or 'php'."
  exit 1
fi

# 时间戳
DATE=$(date +"%Y%m%d_%H%M%S")

# 压缩文件名
ZIP_FILE="${PROJECT}_${DATE}.zip"

echo "🔧 正在打包项目: $PROJECT"

# 压缩目录（排除无关文件）
zip -r "$ZIP_FILE" "$PROJECT" -x "*.git*" -x "vendor/*" -x "node_modules/*" > /dev/null

echo "✅ 打包完成: $ZIP_FILE"

# 推送到 GitHub（自动 commit）
cd "$PROJECT" || exit
git add .
git commit -m "Auto sync at $DATE"
git push origin main
cd ..

echo "✅ GitHub 已更新到最新版本。"
echo "📦 压缩包路径: $(pwd)/$ZIP_FILE"
echo "👉 你现在可以直接上传这个文件到 ChatGPT 让它分析。"
