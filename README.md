# Easy Gateway,一个基于gonlang的网关工具
### 说明
Easy Gateway（一个开箱即用的网关工具）
重要:南京前端-吴文周-周末兼职（18651892475）
PS:此项目基于gin和Vue开发本人负责需求分析，项目设计，日常开发
前后端端代码开源持续更新，可以实现接口动态转发,服务限流，熔断，配合consul实现健康监测，服务注册等等。

### 技术讨论
ng分析:  ng纵然性能卓越功能强大
        弊端如下新增模块例如https需重新编译下载
        运维服务端增加多种配置全靠注释维护迁移人员流失等场景根本不知道配置意义
        配置多功能也就复杂也带来了学习成本增加
eG分析:  使用golang高性能得以保持
        同样得益于golang 跨平台使用底层支持
        可视化配置简单
        数据持久不受代码注释人员流失影响
        配合consul使用完全适配前端微服务的场景，无需重启即可实现动态转发                    

## start

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

