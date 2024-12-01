# 文件操作

## 文件操作示例

1. **创建文件**
```go
file, err := os.Create("example.txt")
if err != nil {
    fmt.Println("创建文件出错:", err)
    return
}
defer file.Close()
fmt.Println("文件创建成功")
```

2. **写入文件**
```go
_, err = file.WriteString("Hello, World!\n")
if err != nil {
    fmt.Println("写入文件出错:", err)
    return
}
fmt.Println("写入成功")
```

3. **读取文件**
```go
data, err := os.ReadFile("example.txt")
if err != nil {
    fmt.Println("读取文件出错:", err)
    return
}
fmt.Println("文件内容:", string(data))
```

4. **追加内容到文件**
```go
file, err = os.OpenFile("example.txt", os.O_APPEND|os.O_WRONLY, 0644)
if err != nil {
    fmt.Println("打开文件出错:", err)
    return
}
defer file.Close()

_, err = file.WriteString("Appending a new line.\n")
if err != nil {
    fmt.Println("追加内容出错:", err)
    return
}
fmt.Println("追加成功")

我们可以看一下os.Create的实现，其实是封装了os.OpenFile
// Create creates or truncates the named file. If the file already exists,
// it is truncated. If the file does not exist, it is created with mode 0666
// (before umask). If successful, methods on the returned File can
// be used for I/O; the associated file descriptor has mode O_RDWR.
// If there is an error, it will be of type *PathError.
func Create(name string) (*File, error) {
	return OpenFile(name, O_RDWR|O_CREATE|O_TRUNC, 0666)
}
其中 
O_RDWR: 以读写模式打开文件。如果文件成功打开，您可以同时读取和写入文件。
O_CREATE: 如果文件不存在，则创建一个新文件。在这个情况下，您需要提供一个文件的权限模式（通常是以八进制表示的），用于设置新创建文件的权限。
O_TRUNC: 如果文件已存在并成功打开，文件的内容将被截断为零长度。这意味着，原有的文件内容将被删除，仅保留文件的元数据。

0666 表示的是文件权限
文件所有者：可读、可写
同组用户：可读、可写
其他用户：可读、可写



```
而go语言中常见的文件标志位有
```
os.O_WRONLY：只写；
os.O_CREATE：创建文件；
os.O_RDONLY：只读；
os.O_RDWR：读写；
os.O_TRUNC：清空；
os.O_APPEND：追加。
```
5. **删除文件**
```go
err = os.Remove("example.txt")
if err != nil {
    fmt.Println("删除文件出错:", err)
    return
}
fmt.Println("文件删除成功")
6.TODO 进阶内容
```

# Json处理
Go语言的encoding/json包用于处理JSON（JavaScript Object Notation）数据的编码和解码。它提供了将Go数据结构（如结构体、切片、映射等）转换为JSON格式的功能，以及将JSON格式的数据解析回Go数据结构的功能。（对于JSON不太了解的同学，可以借助GPT类的工具自主学习一下。)

1. 序列化和反序列化
    - **序列化**: 将Go数据结构转换为JSON格式。
   ```go
   type Person struct {
       Name string `json:"name"`
       Age  int    `json:"age"`
   }

   person := Person{Name: "Alice", Age: 30}
   jsonData, err := json.Marshal(person)
   if err != nil {
       fmt.Println("序列化出错:", err)
   }
   fmt.Println(string(jsonData)) // 输出: {"name":"Alice","age":30}
   ```

    - **反序列化**: 将JSON格式的数据解析为Go数据结构。
   ```go
   jsonString := `{"name":"Bob","age":25}`
   var person2 Person
   err = json.Unmarshal([]byte(jsonString), &person2)
   if err != nil {
       fmt.Println("反序列化出错:", err)
   }
   fmt.Println(person2) // 输出: {Bob 25}
   ```

2. Json tag的运用

//TODO 如何通过反射实现
- JSON标签用于指定在序列化和反序列化时使用的字段名称。
   ```go
   type Product struct {
       ID    int     `json:"id"`
       Name  string  `json:"name"`
       Price float64 `json:"price"`
   }
   ```
2.1 特殊的json tag
```go
// 字段在 JSON 中显示为键 "myName"。  
Field int `json:"myName"`  

// 字段在 JSON 中显示为键 "myName"，  
// 如果字段的值为空，则该字段在对象中将被省略，如上所定义。  
Field int `json:"myName,omitempty"`  

// 字段在 JSON 中显示为键 "Field"（默认），  
// 但如果为空则该字段被跳过。  
// 注意前面的逗号。  
Field int `json:",omitempty"`  

// 该字段被该包忽略。  
Field int `json:"-"`  

// 字段在 JSON 中显示为键 "-"。  
Field int `json:"-,"`
```
2.2 例子
```go
package main

import (
    "encoding/json"
    "fmt"
)

// 定义一个结构体，使用 JSON 标签
type User struct {
    Username string `json:"username" validate:"required"` // 必填字段，这里的validate是一个用于验证的参数的第三方包。通过struct tag，我们可以节省if else的代码。对参数的范围验证，数据校验等。此处的含义为Username为非空的字符串。
    Password string `json:"password" sql:"type:varchar(100)"`
    Age      int    `json:"age,omitempty"` // 如果为空则省略
    Email    string `json:"email,omitempty"` // 如果为空则省略
}

func main() {
    // 创建一个 User 实例
    user := User{
        Username: "Alice",
        Password: "secret",
        Age:      30,
        // Email 字段为空
    }

    // 序列化为 JSON
    jsonData, err := json.Marshal(user)
    if err != nil {
        fmt.Println("序列化出错:", err)
        return
    }
    fmt.Println("序列化结果:", string(jsonData)) // 输出: {"username":"Alice","password":"secret","age":30}

    // 反序列化 JSON
    var newUser User
    jsonString := `{"username":"Bob","password":"12345","age":25}`
    err = json.Unmarshal([]byte(jsonString), &newUser)
    if err != nil {
        fmt.Println("反序列化出错:", err)
        return
    }
    fmt.Println("反序列化结果:", newUser) // 输出: {Bob 12345 25}
}
```
3. struct tag
    - 除了JSON标签，Go的结构体还可以使用其他标签，如`validate`和`sql`标签，用于数据验证和数据库操作。
   ```go
   type User struct {
       Username string `json:"username" validate:"required"`
       Password string `json:"password" sql:"type:varchar(100)"`
   }
   ```

4. 综合例子
    - 下面是一个完整的示例，展示如何使用encoding/json包进行序列化和反序列化。
   ```go
   package main

   import (
       "encoding/json"
       "fmt"
   )

   type Employee struct {
       ID   int    `json:"id"`
       Name string `json:"name"`
   }

   func main() {
       emp := Employee{ID: 1, Name: "John Doe"}
       jsonData, _ := json.Marshal(emp)
       fmt.Println("序列化:", string(jsonData))

       var emp2 Employee
       json.Unmarshal(jsonData, &emp2)
       fmt.Println("反序列化:", emp2)
   }
   ```
5. json包中如何运用的反射

TODO

# viper配置文件管理

Viper 是一个 Go 语言的配置管理库，支持从多种来源读取配置，包括 JSON、YAML、TOML、HCL、ENV 等格式。它提供了简单的 API 来读取和管理应用程序的配置。

## 安装 Viper

可以通过以下命令从 GitHub 安装 Viper：

```bash
go get github.com/spf13/viper
```

## Viper 支持的功能

- 支持多种配置文件格式（JSON、YAML、TOML 等）。
- 支持环境变量和命令行参数的读取。
- 支持热加载配置文件。
- 支持嵌套结构体的解析。
- 支持默认值设置。

## Viper 的常用功能及示例

### 1. 读取配置文件

```yaml
# config.yaml
app:
  name: MyApp
  version: 1.0.0
  description: 这是一个示例应用程序

database:
  host: localhost
  port: 5432
  user: dbuser
  password: dbpassword
  name: mydatabase

server:
  port: 8080
  timeout: 30

logging:
  level: debug
  file: app.log 
```
```go
package main

import (
    "fmt"
    "log"

    "github.com/spf13/viper"
)

func main() {
    viper.SetConfigName("config") // 配置文件名 (不带扩展名)
    viper.SetConfigType("yaml")   // 配置文件类型
    viper.AddConfigPath(".")      // 配置文件路径
    viper.AddConfigPath("/etc")   // 同时可以拥有多个配置路径
    viper.SetConfigFile("config.yaml") //同时指定路径和文件名以及文件类型。

    if err := viper.ReadInConfig(); err != nil {
        log.Fatalf("读取配置文件出错: %s", err)
    }

    // 获取配置值
    appName := viper.GetString("app.name")
    fmt.Println("应用名称:", appName)
}
```

### 2. 设置默认值

```go
viper.SetDefault("app.name", "MyApp")
```

### 3. 监听配置变化

```go
viper.WatchConfig()
viper.OnConfigChange(func(e fsnotify.Event) {
    fmt.Println("配置文件已更改:", e.Name)
})
```

### 4. 从环境变量读取配置

```go
viper.AutomaticEnv() // 自动从环境变量读取配置 
```

### 9. 支持配置文件的合并

```go
viper.SetConfigName("config") // 基础配置文件
viper.AddConfigPath(".")      // 添加路径
viper.MergeInConfig()          // 合并配置文件
```

### 10. 获取所有配置项

```go
allSettings := viper.AllSettings() // 获取所有配置项
fmt.Println("所有配置项:", allSettings)
```

### 11. 使用环境变量作为配置的默认值

```go
viper.SetDefault("app.port", 8080) // 设置默认端口
viper.BindEnv("app.port", "APP_PORT") // 绑定环境变量
```

### 12. 读取 YAML 配置文件���例

```yaml
# config.yaml
app:
  name: MyApp
  version: 1.0.0
  description: 这是一个示例应用程序
```

```go
viper.SetConfigName("config") // 配置文件名
viper.SetConfigType("yaml")   // 配置文件类型
viper.AddConfigPath(".")      // 配置文件路径

if err := viper.ReadInConfig(); err != nil {
    log.Fatalf("读取配置文件出错: %s", err)
}
```

### 13. 处理复杂数据结构

```go
type ServerConfig struct {
    Port    int    `mapstructure:"port"`
    Timeout int    `mapstructure:"timeout"`
}

var serverConfig ServerConfig
if err := viper.UnmarshalKey("server", &serverConfig); err != nil {
    log.Fatalf("无法解析服务器配置: %s", err)
}
fmt.Printf("服务器配置: %+v\n", serverConfig)
```

### 14. 通过反射获取配置

```go
type Config struct {
    AppName string `mapstructure:"app.name"`
    Port    int    `mapstructure:"app.port"`
}

var config Config
if err := viper.Unmarshal(&config); err != nil {
    log.Fatalf("无法解析配置: %s", err)
}
fmt.Printf("应用名称: %s, 端口: %d\n", config.AppName, config.Port)
```

这些是 Viper 的部分功能和示例，展示了如何处理复杂数据结构、合并配置文件、获取所有配置项以及通过反射获取配置。对于想要完整了解viper更多功能的同学。我们提供了官网文档的翻译版本供大家自学。

在day2的demo文件中，演示了如何嵌套定义struct tag，以及配置文件的热更新。
```shell
go run main.go
修改config.yaml 字段如
```
```yaml
app:
  name: MyService
  version: 1.0.1 //修改为1.0.0
  port: 8080

修改完成后保存文件
```

```shell
export APP_PORT=9090
export DB_PASSWORD=newpassword
go run main.go
观察差异
```


