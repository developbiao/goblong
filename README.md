## Goblog (Simple blog written with golang)
This is simple blog traditional MVC project for practice golang, like php laravel.

### Mind map
- Articles manager
- User manager
- User authorization

### About database tables
All table when project startup will be automatic create.

------------

## Go Web程序打包和部署

## 开始部署
### 1. 创建Web目录 

```shell
$mkdir -p /data/www/app.com
```

### 2. 编译文件
Linux 平台amd64 架构可执行一下命令
Mac 或Linux系统下执行

```shell
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app
```

Windows 依次执行以下四个命令：

```shell
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build -o app
```

查询Go支持的平台使用命令

```shell
$ go tool dist list
```


### 3. 上传包文件

```shell
scp app user@YOUR_SERVER_IP:/data/www/app.com
```

配置项目.env文件
编辑.env文件指定服务器环境的配置。



### 4. 运行项目

4.1 服务器上进入项目目录

```shell
$ cd /data/www/app.com
$ ./app

```

4.2 命令测试是否正常

```shell
curl http://localhost:3000
```


### 5. Nginx配置

```shell
server {
    listen       80;
    server_name app.com;

    access_log   /data/log/nginx/app.com/access.log;
    error_log    /data/log/nginx/app.com/error.log;

    location / {
        proxy_pass                 http://127.0.0.1:3000;
        proxy_redirect             off;
        proxy_set_header           Host             $host;
        proxy_set_header           X-Real-IP        $remote_addr;
        proxy_set_header           X-Forwarded-For  $proxy_add_x_forwarded_for;
    }
}
```














