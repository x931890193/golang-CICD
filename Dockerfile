FROM scratch
# 最小镜像，可以有效减少 image 文件大小

ENV TZ=Asia/Shanghai \
    GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    PROGRAM_ENV=pro

WORKDIR /src/build

# 复制构建应用程序所需的代码
COPY ./build .

ADD ./ca-certificates.crt /etc/ssl/certs/

EXPOSE 8011

# 启动服务
CMD ["./main-cicd"]