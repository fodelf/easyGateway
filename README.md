# Easy Gateway,一个基于gonlang的网关工具
### 说明
Easy Gateway（一个开箱即用的网关工具）
重要:南京前端-吴文周-周末兼职（18651892475）
PS:此项目基于gin和Vue开发本人负责需求分析，项目设计，日常开发
前后端端代码开源持续更新，可以实现接口动态转发,服务限流，熔断，配合consul实现健康监测，服务注册等等。
### Links/相关链接

体验地址 

低保真设计 https://modao.cc/app/5ee15024c7e5115c15764a70ed0367dfb8985e40?simulator_type=device&sticky

掘金文档地址 

## start

```
直接执行二进制文件即可
```

### Features

1. pc功能

   - [x] 首页
   - [x] 服务管理
   - [x] 系统管理

#### 技术栈（当前）

1. 前端：[Vue.js]
2. 后端：[Golang]
3. 数据库：[sqlite]

## Project setup （前端）

```
cd  gateway-ui
npm install
```

### 前端项目启动

```
npm run serve
```

## Project setup （后端端）

### 前端项目启动

```
cd  gateway-server
go run main.go
```

