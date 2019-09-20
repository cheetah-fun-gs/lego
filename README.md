# goso

go语言服务端框架/工具集，适用于网关、游戏、网页等。

# 主要特性

网络、数据、路由和业务（模块和处理器）解耦
    - 网络：http、websocket、quic
    - 数据：http-form、json、grpc、自定义二进制
    - 路由：gin-route、consul

# 整体架构

## 演化
![](/docs/jpg/演化.jpg)

## 
![](/docs/jpg/架构.jpg)

# 说明
- gate 网关
    1. gate-client gate对用户
        1. client传递gatePack（gateHead+logicPack）
        2. 解析gateHead（不解析logicPack）获取目标logic
        3. 从logicProxy中获取logic的地址
        4. 生成公共字段和gateHead一起添加进ctx中
        5. 将ctx和logicPack传递给logic服务
    2. gate-logic gate对逻辑
        1. 接收logic请求的ctx和logicPack
        2. 从ctx中获取目标client
        3. 将logicPack传递给client
- logic 逻辑
    1. gate-client传递ctx和logicPack
    1. 解析logicPack
    2. 将ctx和req传递给handler
    3. 获取handler的resp
    4. 从gateProxy中获取gate-logic的地址
    5. 将ctx和resp传递给gate-logic
- module&handle 模块和处理器
    1. 业务逻辑的核心组成
    2. 模块是一个功能的抽象
    3. 处理器是模块或服务暴露的接口
    4. 模块和处理是多对多的关系
    5. 模块间的调用必须通过处理器

# 注意
- 【建议】不要用ctx来传参，仅存（打印日志时抓现场用）或仅取（获取框架层保存的固定字段）
