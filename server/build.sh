#!/bin/bash

##############################################################
# 用途：Stellar 构建脚本
# 作者：DK <ezops.cn@gmail.com>
##############################################################

##############################################################
# 版本号生成
##############################################################
# Commit ID
commit_id=$(git rev-parse --short HEAD 2>/dev/null || echo "unknown")
# 日期
date=$(date +%Y%m%d%H%M%S)
# 版本号
version=${commit_id}-${date}

# 构建，并设置版本号，需要使用绝对路径，否则无法替换值
go build -ldflags "-X stellar/common.SystemVersion=${version}" -o stellar main.go

# 镜像运行
# docker run -p 13306:3306 --name mysql-dev -e MYSQL_ROOT_PASSWORD=123456 -d mysql
# docker run -p 16379:6379 --name redis-dev -e REDIS_PASSWORD=123456 -d redis