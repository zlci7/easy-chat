.PHONY: build-test tag-test publish-test release-test clean

VERSION=latest

SERVER_NAME=user
SERVER_TYPE=rpc

# 测试环境配置
# docker的镜像发布地址（个人版实例）
DOCKER_REPO_TEST=crpi-hej3i56qxeb503xt.cn-hangzhou.personal.cr.aliyuncs.com/easy-chat7/${SERVER_NAME}-${SERVER_TYPE}-dev
# 测试版本
VERSION_TEST=$(VERSION)
# 编译的程序名称
APP_NAME_TEST=easy-im-${SERVER_NAME}-${SERVER_TYPE}-test

# 测试下的编译文件
DOCKER_FILE_TEST=./deploy/dockerfile/Dockerfile_${SERVER_NAME}_${SERVER_TYPE}_dev

# 编译Go程序
compile:
	@echo '开始编译 ${SERVER_NAME}-${SERVER_TYPE}...'
	@mkdir -p bin
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bin/${SERVER_NAME}-${SERVER_TYPE} ./apps/${SERVER_NAME}/${SERVER_TYPE}/${SERVER_NAME}.go
	@echo '编译完成'

# 测试环境的编译发布
build-test: compile
	@echo '开始构建 Docker 镜像...'
	docker build . -f ${DOCKER_FILE_TEST} -t ${APP_NAME_TEST}
	@echo 'Docker 镜像构建完成'

# 镜像的测试标签
tag-test:

	@echo 'create tag ${VERSION_TEST}'
	docker tag ${APP_NAME_TEST} ${DOCKER_REPO_TEST}:${VERSION_TEST}

publish-test:
	@echo '推送镜像 ${VERSION_TEST} 到 ${DOCKER_REPO_TEST}'
	@echo '提示: 如果推送失败，请先登录阿里云镜像仓库:'
	@echo '      docker login --username=aliyun0163842886 crpi-hej3i56qxeb503xt.cn-hangzhou.personal.cr.aliyuncs.com'
	docker push $(DOCKER_REPO_TEST):${VERSION_TEST}
	@echo '推送完成！'

# 清理编译产物和本地镜像
clean:
	@echo '清理编译产物...'
	rm -f bin/${SERVER_NAME}-${SERVER_TYPE}
	docker rmi ${APP_NAME_TEST} 2>/dev/null || true
	docker rmi ${DOCKER_REPO_TEST}:${VERSION_TEST} 2>/dev/null || true
	@echo '清理完成'

release-test: build-test tag-test publish-test