# 设计
这是一个使用go语言和Kratos框架搭建的微服务演示项目。(💯💯包含了各种各样的测试💯💯)

## 特性
* 展示mono-repo风格的项目结构（不同于kratos-layout创建的风格）
* 实践多个服务间通信
* 实例展示kratos和多个基础设施服务结合，例如databases
* 这个项目只是模拟微服务的方案，请大胆发挥想象力。
* 完善的测试示例：单元测试，集成测试

## 组件
这一章节描述项目的各个组件。

### `/api/`
所有的 API `.proto` 文件和生成的 `.go` 文件都在此目录。
此目录结构和目录 `/app/` 类似。

### `/app/`
所有服务的源码都放在这个目录。

#### Service: `/app/catalog/service`
##### 作用/功能
此服务管理所有销售的啤酒信息。
##### 特性
* 和Ent框架集成
* 单元测试（包含biz,data,service层的单元测试）
* 集成测试（真实数据库）

#### Service: `/app/user/service`
##### 作用/功能
管理用户
##### 特性
* 集成Ent框架
* 调用catalog service(简单示例)
* 单元测试（包含biz,data,service层的单元测试）
* 集成测试（真实数据库 + mock catalog grpc ）

### `/pkg/`
公共的包，各个服务都可以引用。

### `/deploy/`
dockerfile和部署脚本

### `/test/integration/`
user和catalog的集成测试(端到端)

## 架构
整个项目的架构蓝图[TBD]
