<div align="center">
    <h1>TimeEmail</h1>
    <p>the email only from you to some one on the time.</p>
</div>

# Introduce

该项目是特地为定时发送邮件而准备的一款仅限私人使用的应用。

# Quick Start

## 直接部署

### 1 克隆项目

```bash
# clone
git clone https://github.com/pykelysia/email-send.git
cd email-send
```

### 2 配置文件

```bash
# email-send >
touch develop.yaml
```

内容如下：

```yaml
UserConfig: 
  UserEmail: youremail@example.com
  EmailPsw: youremail-smtp-password
  EmailHost: smtp.you.use.com:port

EmailTo:
  Addresses:
    - sendto@example.com

LogConfig:
  LogPath: where-you-want-to-log
# 如果你不需要打印日志至文件，可不填

RouteConfig:
  Host: host
  Port: port
```

### 3 运行

```bash
# 直接运行
go run .

# 或构建二进制文件
go build .
```

## `Docker` 部署

在第二步之后：

### 3 构建镜像与容器

```bash
# 构建镜像
docker build -f Dockerfile -t image/name:tag .

# 构建容器
docker run -d image.name:tag
```

## 使用

暂无前端内容（跪）。只要向 `host:port/send` 请求格式满足以下即可：
```json
{
    "time": "yyyy-mm-dd-hour-min-sec-ms",
    "subject": "your email subject",
    "body": "your email body.",
}
```

# TODO

- [ ] 写一个能用的前端。