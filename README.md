# golang微服务与云原生
## 1. 课程简介

## 2. 课程大纲
### 第一期：golang微服务基础
#### 1. golang开发环境搭建
- 安装golang与环境设置
- 第一个Hello World程序
- import包与package
- go get与第三方包
- go mod包管理
#### 2. golang基础热门
- 变量与常量
- 基础数据类型
- 字符串的格式化
- 条件与循环结构
- 函数入门
#### 3. golang数据结构
- 数组
- 切片
- map
- 结构体
- 接口与动态类型
- 指针类型与nil
#### 4. golang函数高级
- 函数与闭包
- golang与面向对象
- 函数与递归
- 接口与鸭子类型
#### 5. golang异常处理
- defer的妙用
- recover与异常捕获
- error类型与panic异常处理
#### 6. golang常用工具包
- time包
- strings包
- strconv包
- reflect包
- 科学计算
- 文件操作
- Json文件处理
- viper配置文件管理
#### 7. golang并发编程
- goroutine与channel
- 使用goroutine与channel实现程序的优雅启停
- select与超时处理
- waitgroup并发控制
- 互斥锁与读写锁
- context上下文与信息传递
- 协程池与任务调度
- 实战案例：生产者与消费者模型
#### 8. golang网络编程（net/http）
- 网络编程概述：TCP/IP协议基础、套接字编程基础
- TCP编程：TCP服务器与TCP客户端
- UDP编程：UDP服务器与UDP客户端
- HTTP编程概述：HTTP协议基础
- HTTP编程：http服务端与http客户端
- RESTful API服务
- RPC编程概述：RPC与HTTP的区别
- 实战案例：使用HTTP手写实现简单的RPC HelloWorld
#### 9. golang操作数据库（GORM）
- ORM与GORM的介绍
- 数据库连接：MySQL、PostgreSQL、SQLite等
- GORM日志和调试
- 模型定义：表结构定义（模型结构体、模型标签、字段类型与约束）
- GORM迁移
- GORM基本操作：CRUD增删改查
- GORM高级操作：条件查询、链式操作、分页排序、关联查询与预加载、原生SQL查询、钩子函数
- GORM事务处理：开启、提交和回滚事务
- 实战案例：使用net/http与GORM实现一个简单的RESTful API服务
#### 10. golang web框架（gin）
- web框架概述：web框架的作用、web框架的分类
- 路由管理：定义路由、路由分组、路由参数和查询参数
- 请求处理：处理GET、POST、PUT、PATCH、DELETE请求、处理请求数据（JSON与表单）、请求上下文（Context）
- 响应处理：返回JSON、自定义响应状态码、文件下载与重定向
- 中间件：中间件概述、gin集成zap日志库、gin集成gorm数据库操作
- 数据验证：gin集成validator数据验证
- 会话和认证：session机制与不足，gin集成JWT认证
- 实战案例：使用gin与GORM实现一个简单的用户管理系统
#### 11. 微服务概念与架构设计
- 单体服务与微服务的概述
- 以电商系统为例，讲解单体服务的设计与架构
- 以电商系统为例，讲解微服务的设计与架构
- 单体服务与微服务的优缺点对比
- 为什么微服务要用rpc而不是http
#### 12. protobuf与grpc
- 重新审视使用HTTP手写实现简单的RPC HelloWorld案例的不足
- protobuf与grpc的介绍
- protobuf与grpc的安装与验证
- protobuf的核心概念
- protobuf与grpc的快速体验
- grpc流模式
- protobuf的数据类型
- grpc metadata机制
- gprc拦截器与验证器
- grpc的错误处理与超时控制
- 实战案例：使用gin与GORM与grpc实现一个简单的用户管理微服务
### 第二期：golang微服务实战
...敬请期待
### 第三期：golang云原生
...敬请期待