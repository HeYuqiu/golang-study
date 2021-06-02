#源镜像
FROM golang:alpine
# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
	GOPROXY="https://goproxy.cn,direct"
# 工作目录
WORKDIR /app
# 将代码复制到容器中，ADD也一样
COPY . .
# 将我们的代码编译成二进制可执行文件  可执行文件名为 app
# 这句代码是在构建镜像时执行的，可以理解是在一个临时容器里面，所以，需要设置代理，不然build时拉依赖包很慢
RUN go build -o app .
# 声明服务端口
EXPOSE 8082
# 启动容器时运行的命令，直接执行上一步生成好的可执行文件app
CMD ["/app/app"]
