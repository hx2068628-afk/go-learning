# Docker 打包和运行问题记录

这份文档只记录本次遇到的 Go 项目打 Docker 镜像、运行容器、读取 `config` 和 `static`、以及 Dockerfile `COPY` 相关问题。

## 1. Docker 运行时报 `exec ./main: no such file or directory`

这个错误不一定是真的没有 `main` 文件。

如果本地直接执行：

```bash
go build -o main .
```

然后把这个 `main` 放进 `alpine` 镜像运行，有可能报：

```text
exec ./main: no such file or directory
```

原因是本地编译出来的 `main` 可能是动态链接的 Linux 程序，依赖 glibc 的 loader，例如：

```text
/lib64/ld-linux-x86-64.so.2
```

但是 `alpine` 默认使用 musl，不带这个 glibc loader，所以文件虽然存在，系统还是会提示找不到。

解决方式是编译成适合 Linux 容器运行的静态二进制：

```bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .
```

参数含义：

```text
CGO_ENABLED=0   禁用 CGO，减少动态链接问题
GOOS=linux      编译给 Linux 系统运行
GOARCH=amd64    编译给 x86_64 架构运行
-o main         输出文件名为 main
.               编译当前目录这个 Go package
```

## 2. `go build main.go` 和 `go build .` 的区别

这两个命令在当前项目里可能都能运行：

```bash
go build main.go
go build .
```

但更推荐：

```bash
go build -o main .
```

区别是：

```text
go build main.go    编译指定的 main.go 文件作为入口
go build .          编译当前目录这个 package
```

如果以后当前目录下还有其他同属于 `package main` 的 `.go` 文件，使用 `go build .` 更稳，不容易漏掉文件。

## 3. 本地开发用 Air，Docker 运行不用 Air

`air` 是本地开发时的热更新工具，适合写代码时使用：

```bash
air
```

Docker 正式运行时不需要 `air`，而是运行编译后的二进制：

```bash
./main
```

也就是说：

```text
本地开发：air
Docker 运行：./main
```

## 4. 本地先编译，再放进 Docker

如果选择本地先编译，Dockerfile 可以写成：

```dockerfile
FROM alpine:latest

WORKDIR /app

COPY main ./main
COPY config ./config
COPY static ./static

EXPOSE 8080

CMD ["./main"]
```

构建和运行命令：

```bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .
docker build -t project01 .
docker run -d --name project01 -p 8080:8080 project01
docker logs project01
```

注意：这种方式要求 Docker 构建上下文里真的有根目录下的 `main` 文件。如果 `.dockerignore` 里写了 `main`，Docker build 时就不会把它复制进镜像。

## 5. Docker build 时自动编译

也可以不在本地手动执行 `go build`，而是在 Docker build 过程中自动编译：

```dockerfile
FROM golang:1.25-alpine AS builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/main .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /out/main ./main
COPY config ./config
COPY static ./static

EXPOSE 8080

CMD ["./main"]
```

这种方式只需要：

```bash
docker build -t project01 .
docker run -d --name project01 -p 8080:8080 project01
```

优点是不用本地手动编译，也不会因为本地编译环境和容器环境不一致导致运行失败。

## 6. `./config` 和 `./static` 是怎么找到的

代码里使用的是相对路径：

```go
r.LoadHTMLGlob("./static/*")
viper.SetConfigFile("./config/config.yaml")
```

这里的 `./` 不是当前源码文件所在目录，而是程序启动时的工作目录。

本地如果这样启动：

```bash
cd /home/xr/goproject/go_code/week08/project01
./main
```

那么 `./config/config.yaml` 实际就是：

```text
/home/xr/goproject/go_code/week08/project01/config/config.yaml
```

Dockerfile 里写了：

```dockerfile
WORKDIR /app
```

这表示容器启动时默认站在 `/app` 目录下。再加上：

```dockerfile
COPY config ./config
COPY static ./static
```

容器里的目录结构就是：

```text
/app
├── main
├── config
│   └── config.yaml
└── static
    ├── fail.html
    ├── hello.html
    └── register.html
```

所以程序运行时：

```text
./config/config.yaml  ->  /app/config/config.yaml
./static/*            ->  /app/static/*
```

只要 `WORKDIR /app` 和 `COPY config ./config`、`COPY static ./static` 对应上，Docker 运行时就能找到配置文件和 HTML 文件。

## 7. `go run /绝对路径/main.go` 为什么也能找到配置

例如：

```bash
go run /home/xr/goproject/go_code/week08/project01/main.go
```

如果执行这条命令时，终端当前目录本来就在项目根目录：

```bash
pwd
# /home/xr/goproject/go_code/week08/project01
```

那么程序的工作目录仍然是项目根目录，所以可以找到：

```text
./config/config.yaml
./static/*
```

但如果换到其他目录再执行：

```bash
cd /tmp
go run /home/xr/goproject/go_code/week08/project01/main.go
```

这时 `./config/config.yaml` 会变成：

```text
/tmp/config/config.yaml
```

就可能找不到配置。

结论：

```text
./ 永远表示程序启动时所在的工作目录，不是 main.go 所在目录。
```

## 8. `COPY static ./static` 是什么意思

Dockerfile 里的：

```dockerfile
COPY static ./static
```

意思是：

```text
把本地项目里的 static 目录，复制到镜像当前工作目录下的 static 目录。
```

如果前面写了：

```dockerfile
WORKDIR /app
```

那么 `./static` 实际就是：

```text
/app/static
```

也就是：

```text
本地项目：
/home/xr/goproject/go_code/week08/project01/static

复制到容器里：
/app/static
```

复制后容器里大概是：

```text
/app
├── main
└── static
    ├── fail.html
    ├── hello.html
    └── register.html
```

代码里写：

```go
r.LoadHTMLGlob("./static/*")
```

程序启动时在 `/app` 目录，所以它会去找：

```text
/app/static/*
```

正好能找到 Dockerfile 复制进去的 HTML 文件。

## 9. `COPY` 可以复制文件，也可以复制文件夹

`COPY` 的基本格式是：

```dockerfile
COPY 源路径 目标路径
```

复制文件：

```dockerfile
COPY main ./main
```

意思是：

```text
本地 main 文件 -> 容器 /app/main
```

复制文件夹：

```dockerfile
COPY static ./static
```

意思是：

```text
本地 static 文件夹 -> 容器 /app/static
```

复制单个配置文件：

```dockerfile
COPY config/config.yaml ./config/config.yaml
```

意思是：

```text
本地 config/config.yaml -> 容器 /app/config/config.yaml
```

一次复制多个文件：

```dockerfile
COPY go.mod go.sum ./
```

意思是：

```text
本地 go.mod -> 容器 /app/go.mod
本地 go.sum -> 容器 /app/go.sum
```

这里的 `.` 或 `./` 表示当前 `WORKDIR`。如果 `WORKDIR /app`，那 `.` 就是 `/app`。

注意：`COPY` 的源路径只能来自 Docker build 的上下文目录，也就是执行下面命令时最后这个 `.` 代表的目录：

```bash
docker build -t project01 .
```

不能随便复制系统任意位置的文件。

## 10. 容器里为什么不能继续用 `127.0.0.1` 连接 MySQL、Redis、RabbitMQ

`127.0.0.1` 永远表示当前程序所在的环境自己。

本地运行 Go 程序时：

```text
Go 程序在宿主机
127.0.0.1 = 宿主机自己
```

如果 MySQL、Redis、RabbitMQ 容器做了端口映射：

```text
宿主机 3306 -> mysql8 容器 3306
宿主机 6379 -> redis 容器 6379
宿主机 5672 -> rabbitmq 容器 5672
```

那么宿主机上的 Go 程序访问下面地址可以成功：

```text
127.0.0.1:3306
127.0.0.1:6379
127.0.0.1:5672
```

但是 Go 程序放进 `project01` 容器后：

```text
Go 程序在 project01 容器
127.0.0.1 = project01 容器自己
```

这时访问：

```text
127.0.0.1:3306
127.0.0.1:6379
127.0.0.1:5672
```

访问的是 `project01` 容器自己，不是 `mysql8`、`redis`、`rabbitmq` 容器。

所以容器访问另一个容器，不能继续写 `127.0.0.1`，应该写容器名：

```text
mysql8:3306
redis:6379
rabbitmq:5672
```

## 11. `docker network` 和 `docker network connect` 的作用

Docker network 的作用是让容器之间可以互相通信，并且可以用容器名解析到对应容器。

例如创建一个网络：

```bash
docker network create app-net
```

把已有容器加入这个网络：

```bash
docker network connect app-net mysql8
docker network connect app-net redis
docker network connect app-net rabbitmq
docker network connect app-net project01
```

加入同一个 network 后，`project01` 容器里就可以用容器名访问其他容器：

```text
mysql8:3306
redis:6379
rabbitmq:5672
```

也可以在创建 `project01` 容器时直接加入网络：

```bash
docker rm -f project01
docker run -d --name project01 --network app-net -p 8080:8080 project01
```

这两种方式的区别：

```text
docker run --network app-net        创建容器时就加入 app-net
docker network connect app-net xxx  容器已经存在后，再加入 app-net
```

所以已有容器时，可以不用删除重建，直接执行：

```bash
docker network connect app-net project01
```

但要注意：加入 network 后，配置文件还是不能继续写 `127.0.0.1`。

`docker network` 解决的是：

```text
容器名 -> 容器 IP
```

不是把：

```text
127.0.0.1:3306
```

自动映射到：

```text
mysql8:3306
```

所以配置仍然要改成：

```yaml
mysql:
  dsn: "root:123456@(mysql8:3306)/test?charset=utf8mb4"
redis:
  addr: "redis:6379"
  password: ""
  db: 0
```

RabbitMQ 连接地址也应该从：

```go
amqp091.Dial("amqp://guest:guest@localhost:5672/")
```

改成：

```go
amqp091.Dial("amqp://guest:guest@rabbitmq:5672/")
```
