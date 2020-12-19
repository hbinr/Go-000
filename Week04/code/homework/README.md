# 作业
> 按照自己的构想，写一个项目满足基本的目录结构和工程，代码需要包含对数据层、业务层、API 注册，以及 main 函数对于服务的注册和启动，信号处理，使用 Wire 构建依赖。可以使用自己熟悉的框架。

## 项目结构
```go
├── api                   # 对外接口定义
│   └── user              
│       └── v1            # v1版本
│           └── proto     # 存放proto文件
├── cmd
│   └── user-interface    # 主程序入口
├── conf                  # 存放配置文件
├── internal              # 业务目录入口，不希望外界调用
│   ├── biz               # 领域服务层，类似DDD中的domain
│   │   ├── entity        # 存放实体
│   │   ├── factory       # 工厂，创建对象
│   │   ├── logic         # 核心业务逻辑
│   │   └── repository    # 仓储，定义数据访问接口
│   │       └── po        # 存放持久化对象
│   ├── data              # 实现了repository定义的接口
│   │   ├── mysql         # mysql实现
│   │   └── redis         # redis实现
│   ├── pkg               # 业务使用的公共目录，与业务强耦合
│   │   └── codec         # 业务码+错误码
│   └── service           # 应用服务层，主要用来服务编排，调biz服务
│       └── assemble      # 对象组合，dto<->do
├── pkg                   # 公共目录，与业务没有任何关系，可被第三方引用
│   ├── cache             # 缓存初始化
│   ├── conf              # 配置定义及初始化
│   ├── database          # 数据库初始化
│   └── tool              # 工具
│       ├── hash          # 加密
│       └── snowflake     # 唯一ID生成
└── scripts               # sql脚本
```

## 运行
**前置条件：**nacos-sever、myql-server、redis-server启动
### 1.配偶环境变量

启动前需配置环境变量：

```sh
export CONF_PROVIDER_FILE_PATH="../../conf/server.yml"
export APP_LOG_CONF_FILE="../../conf/log.yml"
```

### 2.修改myConf.yml的相关配置
#### mysql
主要改 `dsn` 配置
#### redis
如果是本地启动，使用默认配置就行

### 3.在`user_service` 执行 `go run .`
`go run mian.go` 会报错，主要是因为 `wire_gen`目录下的内容读不到。所以执行`go run .`