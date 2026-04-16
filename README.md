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
cp emailsend.example.yaml emailsend.yaml
```

然后填充该配置文件中的内容。

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

直接打开对应的端口的网站即可使用前端。