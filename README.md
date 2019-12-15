# goso

一个高度解耦的go语言服务端框架/工具集

# 特性

- 业务层和框架层解耦，业务模块可复用
- 框架层代码可通过[工具](https://github.com/cheetah-fun-gs/lego-cli)生成，让开发人员更专注业务
- 可展开成微服务架构或收敛为单一进程
- 目录遵循[golang-project-layout](https://github.com/golang-standards/project-layout)
<!-- - 默认支持三种网络协议（http/websocket/quic），用户也可以自定义
- 默认支持两种数据格式（http-json/gRPC），用户也可以自定义
- 默认支持两种服务发现（consul/k8s），用户也可以自定义 -->

# 工具
[goso-cli](https://github.com/cheetah-fun-gs/lego-cli)

# 说明

## 演化
![](/docs/jpg/演化.jpg)

## 单元
![](/docs/jpg/单元.jpg)

## 架构
![](/docs/jpg/架构.jpg)

## 横向划分
- 网络层 net
    - http、websocket、quic
- 数据层 pack
    - http-form、json、grpc、自定义二进制
- 代理层 proxy
    - gin-route、consul
- 业务层 handler

## 纵向划分
- gnet 对用户服务
    - net => gatePack => proxy => caller/handler
- lnet 对服务服务
    - net => logicPack => handler

# 注意
- 【强制】模块间不能共享内存，必须通过处理器相互调用
- 【参考】模块和处理器是多对多的关系
- 【建议】不要用context来传参，仅存（打印日志时抓现场用）或仅取（获取框架层保存的固定字段）
