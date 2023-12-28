# urlnat

这是一个网址转换项目，使用Gin和Gorm构建，结合缩短网址和DDNS功能，方便内网穿透服务使用。[DEMO](https://urlnat.onrender.com)

## 主要功能

- 增删改查URL转换记录
- 重定向后拼接URL路径参数

## 运行环境

- Go
- MySQL

## 快速开始

```
go mod tidy
go run main.go
```

## 接口文档

- 增加

```
POST /add
Content-Type: application/json
{
    "code": "",
    "note": "",
    "url": "",
    "key": "",
    "date": ""
}
```

- 删除

```
POST /delete
Content-Type: application/json
{
    "code": "",
    "key": ""
}
```

- 更新

```
POST /update
Content-Type: application/json
{
    "code": "",
    "note": "",
    "url": "",
    "key": "",
    "date": ""
}
```

- 查询

```
GET /code/<[]?[]=[]>
```
