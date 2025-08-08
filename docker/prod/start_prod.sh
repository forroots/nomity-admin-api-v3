#!/bin/bash

# コンテナスタート
# docker container start prod_nomity_admin_api

# コンテナストップ
# docker container stop prod_nomity_admin_api

# コンテナ削除
# docker container rm prod_nomity_admin_api

# イメージを削除
docker rmi prod_nomity_admin_api

# イメージ作成
docker build -f Dockerfile_prod -t prod_nomity_admin_api .

# コンテナ生成&実行
docker run -d \
  -p 8080:8080 \
  --restart always \
  --name prod_nomity_admin_api_20250808 \
  --log-driver json-file \
  --log-opt max-size=10m \
  --log-opt max-file=10 \
  prod_nomity_admin_api

# 確認
docker ps
