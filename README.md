# lego基础项目

## 目录说明
遵循[project-layout](https://github.com/golang-standards/project-layout)，尽量不要冲突。
```
.
├── assets                  # 资源目录
├── build                   # 构建部署
├── cmd                     # 启动参数解析
├── configs                 # 配置目录
│   ├── dev                 #   开发/联调/测试环境
│   ├── local               #   本地调试环境
│   └── prod                #   生产环境
├── internal                # 项目私有代码
│   ├── biz                 #   业务目录
│   │   ├── aop             #       aop层 规划中 占坑
│   │   ├── dao             #       数据层 规划中 占坑
│   │   ├── handler         #       处理器包
│   │   ├── module          #       模块包
│   │   └── plugin          #       插件包 规划中 占坑
│   ├── common              # 项目公共代码
│   │   ├── common.go       
│   │   ├── configer.go     #   配置管理
│   │   ├── logger.go       #   日志器管理
│   │   ├── mgo.go          #   mongo管理
│   │   ├── redigo.go       #   redis管理
│   │   └── sql.go          #   sqldb管理
│   ├── generated   # 生成代码目录（不要手工修改）
│   │   └── handler #   处理器
│   └── svc         # 服务管理目录
│       ├── gin     #   gin框架
│       └── svc.go
├── main.go
├── pkg             # 共有代码（可做为二方库开放）
└── vendor
```
## 使用说明
依赖legoctl
```
go get github.com/cheetah-fun-gs/legoctl
```
生成框架代码
```
go generate
```
直接构建
```
go build -mod vendor
```
### 创建处理器（接口）
```internal/biz/handler```  
*强制，所有lego项目的核心，请务必遵循以下协议*  

方法1：
1. 拷贝```internal/handler/common/ping.go```到目标路径
2. 修改包名、文件名，替换```Ping```为目标处理器名  

方法2（规划中）：
```
legoctl handler new <path>
```
注意：
1. 一个功能一个子包，只支持一层子包
2. 一个处理器一个文件
3. 处理器逻辑尽量简单，复杂逻辑放到模块包中
4. 不能引用其他处理器包
5. 不能引用插件包

### 模块包
```internal/biz/module```
1. 一个功能一个子包
2. 不能引用处理器包
3. 不能引用插件包
4. 项目最复杂的部分，可以自由发挥，尽量层次清晰

### 配置管理
```internal/common/configer.go```  
已注册```configs```下所有文件，可使用以下库获取配置：
```golang
import (
	mconfiger "github.com/cheetah-fun-gs/goplus/multier/multiconfiger"
)
```
在```default.yml```中添加配置
```golang
    a, err := mconfiger.Get("a")
    b, err := mconfiger.Get("a.b")
    c, err := mconfiger.GetBool("c")
    d := mconfiger.GetBoolD("d", true)
```
添加```demo.yml```, ```abc/test```文件  
*配置名忽略文件后缀，且唯一*  
```golang
    a, err := mconfiger.GetN("demo", "a")
    b, err := mconfiger.GetN("demo", "a.b")
    c, err := mconfiger.GetBoolN("test", "c")
    d := mconfiger.GetBoolDN("test", "d", true)
```
### 日志器和db管理
日志器  
```internal/common/logger.go```  
```configs/xxx/logger.yml```  
redis  
```internal/common/reigo.go```  
```configs/xxx/reigo.yml```  
mongo  
```internal/common/mgo.go```  
```configs/xxx/mgo.yml```  
sqldb  
```internal/common/sql.go```  
```configs/xxx/sql.yml```  
已注册对应配置，可使用以下库获取：
```golang
import (
	mlogger "github.com/cheetah-fun-gs/goplus/multier/multilogger"
    mredigopool "github.com/cheetah-fun-gs/goplus/multier/multiredigopool"
	mmgodb "github.com/cheetah-fun-gs/goplus/multier/multimgodb"
	msqldb "github.com/cheetah-fun-gs/goplus/multier/multisqldb"
)

    mlogger.Debug("%v %v", "a", "b") // default日志器
    mlogger.Debugc(ctx, "%v %v", "a", "b")
    mlogger.DebugN("demo", "%v %v", "a", "b") // demo日志器
    mlogger.DebugcN(ctx, "demo", "%v %v", "a", "b")

    conn := mredigopool.Get() // default redis连接
    defer conn.Close()
    
    conn := mredigopool.GetN("demo") // demo redis连接
    defer conn.Close()
    
    _, err := msqldb.Exec(query, args...) // default sqldb
     _, err := msqldb.ExecN("demo", query, args...) // demo sqldb

    session, err := mmgodb.Clone() // default mgo session
    defer session.Close()

    session, err := mmgodb.CloneN("demo") // demo mgo session
    defer session.Close()
```
