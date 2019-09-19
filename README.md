# goso

go语言服务端框架/工具集。适用于网关、游戏、网页等。

# 整体架构

## 演化
![](/docs/jpg/演化.jpg)

## 架构
![](/docs/jpg/架构.jpg)

# 说明
- gate 网关服务
    1. 解析gateHead（不解析logicPack）获取目标服务
    2. 生成公共字段和gateHead一起添加进ctx中
    3. 从配置或服务治理中获取服务的地址
    4. 将ctx和logicPack传递给logic服务
- logic 逻辑服务
    1. 解析logicPack
    2. 将ctx和logicStruct传递给handle
- module&handle 模块和处理器
    1. 业务逻辑的核心组成
    2. 模块是一个功能的抽象
    3. 处理器是模块或服务暴露的接口
    4. 模块和处理是多对多的关系
    5. 模块间的调用必须通过处理器

# 注意
- 【建议】不要用ctx来传参，仅存（打印日志时抓现场用）或仅取（获取框架层保存的固定字段）
