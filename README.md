# 📊Golang总动员

> **课程编号**：Golang 0️⃣0️⃣1️⃣ \
> **英文名**：Golang-Story

👋欢迎来到Golang总动员课程！本课程旨在帮助学员全面掌握Golang编程语言的基础知识和高级应用（微服务与云原生）。无论你是编程新手还是有经验的开发者，本课程都将带你深入了解Golang的强大功能和最佳实践。通过系统的学习和实战项目，你将能够使用Golang构建高效、可靠的应用程序，并为微服务和云原生开发打下坚实的基础。<img src="./golang.png" alt="golang" width="30" height="28"/> 加入我们，一起开启Golang编程之旅吧！

## 📆课程大纲
> 预计每日学习时长1~2小时.


### 🚀 第一期：golang微服务基础 ➡️（22天）
<details>
<summary>✏️ Chapter1: golang开发环境搭建➡️【1天】</summary>

- ✅安装golang与环境设置
- ✅第一个Hello World程序
- ✅import包与package
- ✅go get与第三方包
- ✅go mod包管理

</details>

<details>
<summary>✏️ Chapter2: golang基础入门➡️【1天】</summary>

- ✅变量与常量
- ✅基础数据类型
- ✅字符串的格式化
- ✅条件与循环结构
- ✅函数入门

</details>

<details>
<summary>✏️ Chapter3: golang数据结构➡️【2天】</summary>

- ✅数组
- ✅切片
- ✅map
- ✅结构体
- ✅接口与动态类型
- ✅指针类型与nil

</details>

<details>
<summary>✏️ Chapter4: golang函数高级➡️【2天】</summary>

- ✅函数与闭包
- ✅函数与递归
- ✅接口类型
- ✅golang与面向对象
- ✅鸭子类型

</details>

<details>
<summary>✏️ Chapter5: golang异常处理➡️【2天】</summary>

- ✅什么是异常
- ✅defer的妙用
- ✅recover与panic异常捕获

</details>

<details>
<summary>✏️ Chapter6: golang常用工具包➡️【2天】</summary>

- ✅time包
- ✅strings包
- ✅strconv包
- ✅reflect包
- ✅科学计算
- ✅文件操作
- ✅Json文件处理
- ✅viper配置文件管理

</details>

<details>
<summary>✏️ Chapter7: golang并发编程➡️【2天】</summary>

- ✅goroutine与channel
- ✅使用goroutine与channel实现程序的优雅启停
- ✅select与超时处理
- ✅waitgroup并发控制
- ✅互斥锁与读写锁
- ✅context上下文与信息传递
- ✅协程池与任务调度
- ✅实战案例：生产者与消费者模型

</details>

<details>
<summary>✏️ Chapter8: golang网络编程（net/http）➡️【2天】</summary>

- ✅网络编程概述：TCP/IP协议基础、套接字编程基础
- ✅TCP编程：TCP服务器与TCP客户端
- ✅UDP编程：UDP服务器与UDP客户端
- ✅HTTP编程概述：HTTP协议基础
- ✅HTTP编程：http服务端与http客户端
- ✅RESTful API服务
- ✅RPC编程概述：RPC与HTTP的区别
- ✅实战案例：使用HTTP手写实现简单的RPC HelloWorld

</details>

<details>
<summary>✏️ Chapter9: golang操作数据库（GORM）➡️【2天】</summary>

- ✅ORM与GORM的介绍
- ✅数据库连接：MySQL、PostgreSQL、SQLite等
- ✅GORM日志和调试
- ✅模型定义：表结构定义（模型结构体、模型标签、字段类型与约束）
- ✅GORM迁移
- ✅GORM基本操作：CRUD增删改查
- ✅GORM高级操作：条件查询、链式操作、分页排序、关联查询与预加载、原生SQL查询、钩子函数
- ✅GORM事务处理：开启、提交和回滚事务
- ✅实战案例：使用net/http与GORM实现一个简单的RESTful API服务

</details>

<details>
<summary>✏️ Chapter10: golang web框架（gin）➡️【2天】</summary>

- ✅web框架概述：web框架的作用、web框架的分类
- ✅路由管理：定义路由、路由分组、路由参数和查询参数
- ✅请求处理：处理GET、POST、PUT、PATCH、DELETE请求、处理请求数据（JSON与表单）、请求上下文（Context）
- ✅响应处理：返回JSON、自定义响应状态码、文件下载与重定向
- ✅中间件：中间件概述、gin集成zap日志库、gin集成gorm数据库操作
- ✅数据验证：gin集成validator数据验证
- ✅会话和认证：session机制与不足，gin集成JWT认证
- ✅实战案例：使用gin与GORM实现一个简单的用户管理系统

</details>

<details>
<summary>✏️ Chapter11: 微服务概念与架构设计➡️【2天】</summary>

- ✅单体服务与微服务的概述
- ✅以电商系统为例，讲解单体服务的设计与架构
- ✅以电商系统为例，讲解微服务的设计与架构
- ✅单体服务与微服务的优缺点对比
- ✅为什么微服务要用rpc而不是http

</details>

<details>
<summary>✏️ Chapter12: protobuf与grpc➡️【2天】</summary>

- ✅重新审视使用HTTP手写实现简单的RPC HelloWorld案例的不足
- ✅protobuf与grpc的介绍
- ✅protobuf与grpc的安装与验证
- ✅protobuf的核心概念
- ✅protobuf与grpc的快速体验
- ✅grpc流模式
- ✅protobuf的数据类型
- ✅grpc metadata机制
- ✅gprc拦截器与验证器
- ✅grpc的错误处理与超时控制
- ✅实战案例：使用gin与GORM与grpc实现一个简单的用户管理微服务

</details>

### 🚀 第二期：golang微服务实战
...敬请期待
### 🚀 第三期：golang云原生实战
...敬请期待


## 🛠️我们将用到的工具
<p align='center'>
<img src="https://img.shields.io/badge/language-golang 1.23.0-brightgreen"> 
<img src="https://img.shields.io/badge/package-gin v1.10.0-ff69b4">
<img src="https://img.shields.io/badge/package-grpc v1.67.0-ff69b4">
<br>
<img src="https://img.shields.io/badge/package-protobuf v1.34.2-ff69b4">
<img src="https://img.shields.io/badge/package-viper v1.19.0-ff69b4">
<img src="https://img.shields.io/badge/package-validator v10.22.1-ff69b4">
</p>

## 🧑‍💻贡献者
### 🧑‍🔧课程开发
| 昵称 | 内容           | 联系方式                                  |
| ------ | ---------------- | --------------------------------------------- |
| 李祖贤 | ▶️Chapter4: golang函数高级<br>▶️Chapter5: golang异常处理<br>▶️Chapter7: golang并发编程<br>▶️Chapter8: golang网络编程（net/http）<br>▶️Chapter9: golang操作数据库（GORM）<br>▶️Chapter10: golang web框架（gin） |1028851587@qq.com|
| 李梦吉 | ▶️Chapter6: golang常用工具包<br>▶️Chapter12: protobuf与grpc     |[https://github.com/voidspiral](https://github.com/voidspiral)|
| 于沼懿 | ▶️Chapter1: golang开发环境搭建<br>▶️Chapter2: golang基础入门<br>▶️Chapter3: golang数据结构    |[https://github.com/yzy-fulture](https://github.com/yzy-fulture)|
| 王嘉鹏 | ▶️Chapter11: 微服务概念与架构设计     |[https://github.com/DaPengDiting](https://github.com/DaPengDiting)|

### 🧩改进优化
<h2 align='left'><b>📫 欢迎沟通鸭！</b></h2>

- 📧 **邮箱**：1028851587@qq.com
- 📺 **bilibili**：[二次元的Datawhale](https://space.bilibili.com/431850986)

## 📚学习资料

## :memo:License:sparkling_heart:
如果您喜欢`Golang总动员`的开拓者，由Lizuxian发起，喜欢❤️请收藏给一个赞吧👍

Copyright :copyright:2024 [Lizuxian](https://github.com/Leolee11111)