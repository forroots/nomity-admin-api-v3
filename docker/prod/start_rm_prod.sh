#!/bin/bash

# コンテナスタート
# docker container start prod_nomity_admin_api

# コンテナストップ
docker container stop prod_nomity_admin_api

# コンテナ削除
docker container rm prod_nomity_admin_api

# イメージを削除
docker rmi prod_nomity_admin_api

# イメージ作成
docker build -f Dockerfile_prod -t prod_nomity_admin_api .

# コンテナ生成&実行
docker run --rm -p 8080:8080 --name prod_nomity_admin_api prod_nomity_admin_api .

# 確認
docker ps
