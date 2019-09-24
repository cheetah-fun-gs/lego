# goso

一个高度解耦的go语言服务端框架/工具集

# 整体架构

## 演化
![](/docs/jpg/演化.jpg)

## 单元
![](/docs/jpg/单元.jpg)

## 架构
![](/docs/jpg/架构.jpg)

# 说明

## 四层
- 网络 net
    - http、websocket、quic
- 数据 pack
    - http-form、json、grpc、自定义二进制
- 代理 proxy
    - gin-route、consul
- 处理器 handler

## 三模块
- gnet 
    - net + gatePack + proxy + caller
- lnet 
    - net + logicPack + proxy + caller
- soer 
    - handler + proxy + caller

# 注意
- 【强制】模块间不能共享内存，必须通过处理器相互调用
- 【参考】模块和处理是多对多的关系
- 【建议】不要用ctx来传参，仅存（打印日志时抓现场用）或仅取（获取框架层保存的固定字段）

