## Viper是什么？

Viper是一个完整的Go应用程序配置解决方案，包括[12-Factor应用](https://12factor.net/#the_twelve_factors)。
它旨在在应用程序内部工作，可以处理所有类型的配置需求和格式。它支持：

* 设置默认值
* 从JSON、TOML、YAML、HCL、envfile和Java属性配置文件读取
* 实时监视和重新读取配置文件（可选）
* 从环境变量读取
* 从远程配置系统（etcd或Consul）读取，并监视更改
* 从命令行标志读取
* 从缓冲区读取
* 设置显式值

Viper可以被视为您所有应用程序配置需求的注册中心。

## 安装

```shell
go get github.com/spf13/viper
```

**注意：** Viper使用[Go Modules](https://go.dev/wiki/Modules)来管理依赖关系。

## 为什么选择Viper？

在构建现代应用程序时，您不想担心配置文件格式；您想专注于构建出色的软件。
Viper在这里帮助您。

Viper为您做以下事情：

1. 查找、加载和反序列化JSON、TOML、YAML、HCL、INI、envfile或Java属性格式的配置文件。
2. 提供机制为不同的配置选项设置默认值。
3. 提供机制为通过命令行标志指定的选项设置覆盖值。
4. 提供别名系统以轻松重命名参数而不破坏现有代码。
5. 使您能够轻松区分用户提供的命令行或配置文件与默认值相同的情况。

Viper使用以下优先级顺序。每个项目优先于下面的项目：

* 显式调用`Set`
* 标志
* 环境变量
* 配置
* 键/值存储
* 默认值

**重要：** Viper配置键不区分大小写。
关于是否使其可选的讨论仍在进行中。

## 将值放入Viper

### 建立默认值

一个好的配置系统将支持默认值。默认值对于一个键不是必需的，但在未通过配置文件、环境变量、远程配置或标志设置键的情况下，它是有用的。

示例：

```go
viper.SetDefault("ContentDir", "content")
viper.SetDefault("LayoutDir", "layouts")
viper.SetDefault("Taxonomies", map[string]string{"tag": "tags", "category": "categories"})
```

### 读取配置文件

Viper需要最少的配置，以便知道在哪里查找配置文件。
Viper支持JSON、TOML、YAML、HCL、INI、envfile和Java属性文件。Viper可以搜索多个路径，但当前单个Viper实例仅支持单个配置文件。
Viper不默认任何配置搜索路径，将默认决策留给应用程序。

以下是如何使用Viper搜索和读取配置文件的示例。
并非所有特定路径都是必需的，但至少应提供一个路径以期望找到配置文件。

```go
viper.SetConfigName("config") // 配置文件的名称（不带扩展名）
viper.SetConfigType("yaml") // 如果配置文件没有扩展名，则必需
viper.AddConfigPath("/etc/appname/")   // 查找配置文件的路径
viper.AddConfigPath("$HOME/.appname")  // 多次调用以添加多个搜索路径
viper.AddConfigPath(".")               // 可选地在工作目录中查找配置
err := viper.ReadInConfig() // 查找并读取配置文件
if err != nil { // 处理读取配置文件时的错误
	panic(fmt.Errorf("fatal error config file: %w", err))
}
```

您可以处理未找到配置文件的特定情况，如下所示：

```go
if err := viper.ReadInConfig(); err != nil {
	if _, ok := err.(viper.ConfigFileNotFoundError); ok {
		// 配置文件未找到；如有需要可忽略错误
	} else {
		// 找到配置文件但产生了其他错误
	}
}

// 配置文件找到并成功解析
```

*注意[自1.6起]：* 您还可以有一个没有扩展名的文件，并以编程方式指定格式。对于那些位于用户主目录中的没有扩展名的配置文件，如`.bashrc`

### 写入配置文件

从配置文件读取是有用的，但有时您希望存储在运行时所做的所有修改。
为此，有一系列命令可用，每个命令都有其特定目的：

* WriteConfig - 将当前viper配置写入预定义路径（如果存在）。如果没有预定义路径则会出错。如果当前配置文件存在，将覆盖它。
* SafeWriteConfig - 将当前viper配置写入预定义路径。如果没有预定义路径则会出错。如果当前配置文件存在，将不会覆盖它。
* WriteConfigAs - 将当前viper配置写入给定的文件路径。如果存在，将覆盖给定的文件。
* SafeWriteConfigAs - 将当前viper配置写入给定的文件路径。如果存在，将不会覆盖给定的文件。

作为经验法则，所有标记为安全的内容不会覆盖任何文件，而是仅在不存在时创建，而默认行为是创建或截断。

小示例部分：

```go
viper.WriteConfig() // 将当前配置写入由'viper.AddConfigPath()'和'viper.SetConfigName'设置的预定义路径
viper.SafeWriteConfig()
viper.WriteConfigAs("/path/to/my/.config")
viper.SafeWriteConfigAs("/path/to/my/.config") // 将出错，因为它已经被写入
viper.SafeWriteConfigAs("/path/to/my/.other_config")
```

### 监视和重新读取配置文件

Viper支持在运行时让您的应用程序实时读取配置文件的能力。

不再需要重启服务器以使配置生效，使用viper的应用程序可以在运行时读取配置文件的更新而不会错过任何内容。

只需告诉viper实例监视配置。
可选地，您可以提供一个函数，让Viper在每次发生更改时运行。

**确保在调用`WatchConfig()`之前添加所有configPaths**

```go
viper.OnConfigChange(func(e fsnotify.Event) {
	fmt.Println("配置文件已更改:", e.Name)
})
viper.WatchConfig()
```

### 从io.Reader读取配置

Viper预定义了许多配置源，例如文件、环境变量、标志和远程K/V存储，但您并不受限于它们。您还可以实现自己的所需配置源并将其提供给viper。

```go
viper.SetConfigType("yaml") // 或viper.SetConfigType("YAML")

// 任何将此配置引入程序的方法。
var yamlExample = []byte(`
Hacker: true
name: steve
hobbies:
- skateboarding
- snowboarding
- go
clothing:
  jacket: leather
  trousers: denim
age: 35
eyes : brown
beard: true
`)

viper.ReadConfig(bytes.NewBuffer(yamlExample))

viper.Get("name") // 这将是"steve"
```

### 设置覆盖

这些可以来自命令行标志，或来自您自己的应用程序逻辑。

```go
viper.Set("Verbose", true)
viper.Set("LogFile", LogFile)
viper.Set("host.port", 5899)   // 设置子集
```

### 注册和使用别名

别名允许通过多个键引用单个值

```go
viper.RegisterAlias("loud", "Verbose")

viper.Set("verbose", true) // 与下一行相同的结果
viper.Set("loud", true)   // 与前一行相同的结果

viper.GetBool("loud") // true
viper.GetBool("verbose") // true
```

### 使用环境变量

Viper完全支持环境变量。这使得12因素应用程序开箱即用。存在五种方法来帮助处理环境变量：

* `AutomaticEnv()`
* `BindEnv(string...) : error`
* `SetEnvPrefix(string)`
* `SetEnvKeyReplacer(string...) *strings.Replacer`
* `AllowEmptyEnv(bool)`

_在处理环境变量时，重要的是要认识到Viper将环境变量视为区分大小写。_

Viper提供了一种机制来确保环境变量的唯一性。通过使用`SetEnvPrefix`，您可以告诉Viper在读取环境变量时使用前缀。`BindEnv`和`AutomaticEnv`都将使用此前缀。

`BindEnv`接受一个或多个参数。第一个参数是键名，其余是要绑定到此键的环境变量名称。如果提供了多个，它们将按指定顺序优先。环境变量名称是区分大小写的。如果未提供环境变量名称，则Viper将自动假定环境变量与以下格式匹配：前缀 + "_" + 键名（全部大写）。当您显式提供环境变量名称（第二个参数）时，它**不会**自动添加前缀。例如，如果第二个参数是"id"，Viper将查找环境变量"ID"。

在处理环境变量时需要认识到的一件重要事情是，值将在每次访问时读取。Viper在调用`BindEnv`时不会固定值。

`AutomaticEnv`是一个强大的助手，尤其是与`SetEnvPrefix`结合使用时。当调用时，Viper将在每次进行`viper.Get`请求时检查环境变量。它将应用以下规则。它将检查名称与键大写并以`EnvPrefix`（如果设置）为前缀的环境变量。

`SetEnvKeyReplacer`允许您使用`strings.Replacer`对象重写Env键。这在您希望在`Get()`调用中使用`-`或其他内容，但希望您的环境变量使用`_`分隔符时非常有用。使用它的示例可以在`viper_test.go`中找到。

或者，您可以使用`EnvKeyReplacer`与`NewWithOptions`工厂函数一起使用。与`SetEnvKeyReplacer`不同，它接受一个`StringReplacer`接口，允许您编写自定义字符串替换逻辑。

默认情况下，空环境变量被视为未设置，并将回退到下一个配置源。要将空环境变量视为已设置，请使用`AllowEmptyEnv`方法。

#### 环境示例

```go
SetEnvPrefix("spf") // 将自动大写
BindEnv("id")

os.Setenv("SPF_ID", "13") // 通常在应用程序外部完成

id := Get("id") // 13
```

### 使用标志

Viper能够绑定到标志。具体来说，Viper支持在[Cobra](https://github.com/spf13/cobra)库中使用的`Pflags`。

与`BindEnv`一样，绑定方法调用时不会设置值，而是在访问时设置。这意味着您可以尽早绑定，甚至在`init()`函数中。

对于单个标志，`BindPFlag()`方法提供此功能。

示例：

```go
serverCmd.Flags().Int("port", 1138, "运行应用程序服务器的端口")
viper.BindPFlag("port", serverCmd.Flags().Lookup("port"))
```

您还可以绑定现有的pflags（pflag.FlagSet）：

示例：

```go
pflag.Int("flagname", 1234, "flagname的帮助信息")

pflag.Parse()
viper.BindPFlags(pflag.CommandLine)

i := viper.GetInt("flagname") // 从viper而不是pflag检索值
```

在Viper中使用[pflag](https://github.com/spf13/pflag/)并不排除使用标准库中的[flag](https://golang.org/pkg/flag/)包的其他包。pflag包可以通过导入这些标志来处理为flag包定义的标志。通过调用pflag包提供的便利函数AddGoFlagSet()来实现。

示例：

```go
package main

import (
	"flag"
	"github.com/spf13/pflag"
)

func main() {

	// 使用标准库"flag"包
	flag.Int("flagname", 1234, "flagname的帮助信息")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	i := viper.GetInt("flagname") // 从viper检索值

	// ...
}
```

#### 标志接口

Viper提供两个Go接口以绑定其他标志系统，如果您不使用`Pflags`。

`FlagValue`表示单个标志。这是实现此接口的非常简单的示例：

```go
type myFlag struct {}
func (f myFlag) HasChanged() bool { return false }
func (f myFlag) Name() string { return "my-flag-name" }
func (f myFlag) ValueString() string { return "my-flag-value" }
func (f myFlag) ValueType() string { return "string" }
```

一旦您的标志实现了此接口，您可以简单地告诉Viper绑定它：

```go
viper.BindFlagValue("my-flag-name", myFlag{})
```

`FlagValueSet`表示一组标志。这是实现此接口的非常简单的示例：

```go
type myFlagSet struct {
	flags []myFlag
}

func (f myFlagSet) VisitAll(fn func(FlagValue)) {
	for _, flag := range flags {
		fn(flag)
	}
}
```

一旦您的标志集实现了此接口，您可以简单地告诉Viper绑定它：

```go
fSet := myFlagSet{
	flags: []myFlag{myFlag{}, myFlag{}},
}
viper.BindFlagValues("my-flags", fSet)
```

### 远程键/值存储支持

要在Viper中启用远程支持，请对`viper/remote`包进行空白导入：

`import _ "github.com/spf13/viper/remote"`

Viper将从键/值存储（如etcd或Consul）中检索的路径读取配置字符串（格式为JSON、TOML、YAML、HCL或envfile）。这些值优先于默认值，但会被从磁盘、标志或环境变量检索的配置值覆盖。

Viper支持多个主机。要使用，请传递用`;`分隔的端点列表。例如`http://127.0.0.1:4001;http://127.0.0.1:4002`。

Viper使用[crypt](https://github.com/sagikazarmark/crypt)从K/V存储中检索配置，这意味着您可以存储加密的配置值，并在拥有正确的gpg密钥环时自动解密。加密是可选的。

您可以独立于本地配置或与本地配置结合使用远程配置。

`crypt`有一个命令行助手，您可以使用它将配置放入K/V存储。`crypt`默认使用http://127.0.0.1:4001上的etcd。

```bash
$ go get github.com/sagikazarmark/crypt/bin/crypt
$ crypt set -plaintext /config/hugo.json /Users/hugo/settings/config.json
```

确认您的值已设置：

```bash
$ crypt get -plaintext /config/hugo.json
```

有关如何设置加密值或如何使用Consul的示例，请参阅`crypt`文档。

### 远程键/值存储示例 - 未加密

#### etcd
```go
viper.AddRemoteProvider("etcd", "http://127.0.0.1:4001","/config/hugo.json")
viper.SetConfigType("json") // 因为在字节流中没有文件扩展名，支持的扩展名为"json"、"toml"、"yaml"、"yml"、"properties"、"props"、"prop"、"env"、"dotenv"
err := viper.ReadRemoteConfig()
```

#### etcd3
```go
viper.AddRemoteProvider("etcd3", "http://127.0.0.1:4001","/config/hugo.json")
viper.SetConfigType("json") // 因为在字节流中没有文件扩展名，支持的扩展名为"json"、"toml"、"yaml"、"yml"、"properties"、"props"、"prop"、"env"、"dotenv"
err := viper.ReadRemoteConfig()
```

#### Consul
您需要将键设置为Consul键/值存储，值为包含所需配置的JSON值。
例如，创建一个Consul键/值存储键`MY_CONSUL_KEY`，值为：

```json
{
    "port": 8080,
    "hostname": "myhostname.com"
}
```

```go
viper.AddRemoteProvider("consul", "localhost:8500", "MY_CONSUL_KEY")
viper.SetConfigType("json") // 需要明确设置为json
err := viper.ReadRemoteConfig()

fmt.Println(viper.Get("port")) // 8080
fmt.Println(viper.Get("hostname")) // myhostname.com
```

#### Firestore

```go
viper.AddRemoteProvider("firestore", "google-cloud-project-id", "collection/document")
viper.SetConfigType("json") // 配置的格式："json"、"toml"、"yaml"、"yml"
err := viper.ReadRemoteConfig()
```

当然，您也可以使用`SecureRemoteProvider`。

#### NATS

```go
viper.AddRemoteProvider("nats", "nats://127.0.0.1:4222", "myapp.config")
viper.SetConfigType("json")
err := viper.ReadRemoteConfig()
```

### 远程键/值存储示例 - 加密

```go
viper.AddSecureRemoteProvider("etcd","http://127.0.0.1:4001","/config/hugo.json","/etc/secrets/mykeyring.gpg")
viper.SetConfigType("json") // 因为在字节流中没有文件扩展名，支持的扩展名为"json"、"toml"、"yaml"、"yml"、"properties"、"props"、"prop"、"env"、"dotenv"
err := viper.ReadRemoteConfig()
```

### 监视etcd中的更改 - 未加密

```go
// 或者，您可以创建一个新的viper实例。
var runtime_viper = viper.New()

runtime_viper.AddRemoteProvider("etcd", "http://127.0.0.1:4001", "/config/hugo.yml")
runtime_viper.SetConfigType("yaml") // 因为在字节流中没有文件扩展名，支持的扩展名为"json"、"toml"、"yaml"、"yml"、"properties"、"props"、"prop"、"env"、"dotenv"

// 第一次从远程配置读取。
err := runtime_viper.ReadRemoteConfig()

// 反序列化配置
runtime_viper.Unmarshal(&runtime_conf)

// 打开一个goroutine以监视远程更改
go func(){
	for {
		time.Sleep(time.Second * 5) // 每次请求后的延迟

		// 当前仅在etcd支持下测试
		err := runtime_viper.WatchRemoteConfig()
		if err != nil {
			log.Errorf("无法读取远程配置: %v", err)
			continue
		}

		// 将新配置反序列化到我们的运行时配置结构中。您还可以使用通道实现信号以通知系统更改
		runtime_viper.Unmarshal(&runtime_conf)
	}
}()
```

## 从Viper获取值

在Viper中，根据值的类型，有几种方法可以获取值。
以下函数和方法存在：

* `Get(key string) : any`
* `GetBool(key string) : bool`
* `GetFloat64(key string) : float64`
* `GetInt(key string) : int`
* `GetIntSlice(key string) : []int`
* `GetString(key string) : string`
* `GetStringMap(key string) : map[string]any`
* `GetStringMapString(key string) : map[string]string`
* `GetStringSlice(key string) : []string`
* `GetTime(key string) : time.Time`
* `GetDuration(key string) : time.Duration`
* `IsSet(key string) : bool`
* `AllSettings() : map[string]any`

需要认识到的一件重要事情是，如果未找到，则每个Get函数将返回零值。如果给定键存在，`IsSet()`方法已提供以检查该键是否存在。

如果值已设置但无法解析为请求的类型，则也将返回零值。

示例：
```go
viper.GetString("logfile") // 不区分大小写的设置和获取
if viper.GetBool("verbose") {
	fmt.Println("启用详细模式")
}
```
### 访问嵌套键

访问器方法还接受格式化路径以访问深层嵌套键。例如，如果加载了以下JSON文件：

```json
{
    "host": {
        "address": "localhost",
        "port": 5799
    },
    "datastore": {
        "metric": {
            "host": "127.0.0.1",
            "port": 3099
        },
        "warehouse": {
            "host": "198.0.0.1",
            "port": 2112
        }
    }
}
```

Viper可以通过传递以`.`分隔的键路径访问嵌套字段：

```go
GetString("datastore.metric.host") // (返回"127.0.0.1")
```

这遵循上述建立的优先级规则；对路径的搜索将通过剩余的配置注册表级联，直到找到为止。

例如，给定此配置文件，`datastore.metric.host`和`datastore.metric.port`都已定义（并可能被覆盖）。如果此外`datastore.metric.protocol`在默认值中定义，Viper也会找到它。

但是，如果`datastore.metric`被（通过标志、环境变量、`Set()`方法等）用直接值覆盖，则所有`datastore.metric`的子键变为未定义，它们被更高优先级的配置级别“遮蔽”。

Viper可以通过在路径中使用数字访问数组索引。例如：

```jsonc
{
    "host": {
        "address": "localhost",
        "ports": [
            5799,
            6029
        ]
    },
    "datastore": {
        "metric": {
            "host": "127.0.0.1",
            "port": 3099
        },
        "warehouse": {
            "host": "198.0.0.1",
            "port": 2112
        }
    }
}
```

```go
GetInt("host.ports.1") // 返回6029
```

最后，如果存在与分隔键路径匹配的键，则将返回其值。例如：

```jsonc
{
    "datastore.metric.host": "0.0.0.0",
    "host": {
        "address": "localhost",
        "port": 5799
    },
    "datastore": {
        "metric": {
            "host": "127.0.0.1",
            "port": 3099
        },
        "warehouse": {
            "host": "198.0.0.1",
            "port": 2112
        }
    }
}
```

```go
GetString("datastore.metric.host") // 返回"0.0.0.0"
```

### 提取子树

在开发可重用模块时，提取配置的子集并将其传递给模块通常很有用。
这样，模块可以多次实例化，使用不同的配置。

例如，一个应用程序可能会使用多个不同的缓存存储以不同的目的：

```yaml
cache:
  cache1:
    max-items: 100
    item-size: 64
  cache2:
    max-items: 200
    item-size: 80
```

我们可以将缓存名称传递给模块（例如`NewCache("cache1")`），但这将需要奇怪的连接以访问配置键，并且与全局配置的分离程度较低。

因此，不如将表示配置子集的Viper实例传递给构造函数：

```go
cache1Config := viper.Sub("cache.cache1")
if cache1Config == nil { // Sub在找不到键时返回nil
	panic("未找到缓存配置")
}

cache1 := NewCache(cache1Config)
```

**注意：** 始终检查`Sub`的返回值。如果找不到键，它将返回`nil`。

在内部，`NewCache`函数可以直接访问`max-items`和`item-size`键：

```go
func NewCache(v *Viper) *Cache {
	return &Cache{
		MaxItems: v.GetInt("max-items"),
		ItemSize: v.GetInt("item-size"),
	}
}
```

生成的代码易于测试，因为它与主配置结构解耦，并且更易于重用（出于同样的原因）。

### 反序列化

您还可以选择将所有或特定值反序列化到结构、映射等中。

有两种方法可以做到这一点：

* `Unmarshal(rawVal any) : error`
* `UnmarshalKey(key string, rawVal any) : error`

示例：

```go
type config struct {
	Port int
	Name string
	PathMap string `mapstructure:"path_map"`
}

var C config

err := viper.Unmarshal(&C)
if err != nil {
	t.Fatalf("无法解码到结构中，%v", err)
}
```

如果您想反序列化包含点（默认键分隔符）的键的配置，则必须更改分隔符：

```go
v := viper.NewWithOptions(viper.KeyDelimiter("::"))

v.SetDefault("chart::values", map[string]any{
	"ingress": map[string]any{
		"annotations": map[string]any{
			"traefik.frontend.rule.type":                 "PathPrefix",
			"traefik.ingress.kubernetes.io/ssl-redirect": "true",
		},
	},
})

type config struct {
	Chart struct{
		Values map[string]any
	}
}

var C config

v.Unmarshal(&C)
```

Viper还支持反序列化到嵌入式结构中：

```go
/*
示例配置：

module:
    enabled: true
    token: 89h3f98hbwf987h3f98wenf89ehf
*/
type config struct {
	Module struct {
		 Enabled bool

		 moduleConfig `mapstructure:",squash"`
	}
}

// moduleConfig可以在特定模块的包中
type moduleConfig struct {
	Token string
}

var C config

err := viper.Unmarshal(&C)
if err != nil {
	t.Fatalf("无法解码到结构中，%v", err)
}
```

Viper在内部使用[github.com/go-viper/mapstructure](https://github.com/go-viper/mapstructure)进行反序列化，默认使用`mapstructure`标签。

### 解码自定义格式

Viper的一个常见请求功能是添加更多值格式和解码器。
例如，将字符（点、逗号、分号等）分隔的字符串解析为切片。

这已经在Viper中使用mapstructure解码钩子可用。

有关详细信息，请阅读[这篇博文](https://sagikazarmark.hu/blog/decoding-custom-formats-with-viper/)。

### 序列化为字符串

您可能需要将Viper中持有的所有设置序列化为字符串，而不是将它们写入文件。
您可以使用您最喜欢的格式的序列化器与`AllSettings()`返回的配置一起使用。

```go
import (
	yaml "gopkg.in/yaml.v2"
	// ...
)

func yamlStringSettings() string {
	c := viper.AllSettings()
	bs, err := yaml.Marshal(c)
	if err != nil {
		log.Fatalf("无法将配置序列化为YAML: %v", err)
	}
	return string(bs)
}
```

## Viper还是Vipers？

Viper提供了一个全局实例（单例），开箱即用。

尽管它使设置配置变得简单，但通常不建议使用它，因为这会使测试变得更加困难，并可能导致意外行为。

最佳实践是初始化一个Viper实例并在必要时传递它。

全局实例_可能_在未来被弃用。
有关更多详细信息，请参见[#1855](https://github.com/spf13/viper/issues/1855)。

### 使用多个viper

您还可以为应用程序创建多个不同的viper。每个viper将具有自己独特的配置和值集。每个viper可以从不同的配置文件、键值存储等读取。Viper包支持的所有函数在viper上都有镜像。

示例：

```go
x := viper.New()
y := viper.New()

x.SetDefault("ContentDir", "content")
y.SetDefault("ContentDir", "foobar")

//...
```

在使用多个viper时，用户需要跟踪不同的viper。

## 问答

### 为什么叫“Viper”？

答：Viper旨在成为[Cobra](https://github.com/spf13/cobra)的[伴侣](http://en.wikipedia.org/wiki/Viper_(G.I._Joe))。虽然两者可以完全独立操作，但结合在一起，它们构成了处理应用程序基础需求的强大组合。

### 为什么叫“Cobra”？

难道没有更好的名字来形容[指挥官](http://en.wikipedia.org/wiki/Cobra_Commander)吗？

### Viper支持区分大小写的键吗？

**简而言之：** 不支持。

Viper合并来自各种来源的配置，其中许多来源要么不区分大小写，要么使用与其余来源不同的大小写（例如环境变量）。
为了在使用多个来源时提供最佳体验，决定使所有键不区分大小写。

已经有几次尝试实现区分大小写，但不幸的是，这并不是那么简单。我们可能会在[Viper v2](https://github.com/spf13/viper/issues/772)中尝试实现它，但尽管最初有噪音，但似乎并没有那么多请求。

您可以通过填写此反馈表投票支持区分大小写：https://forms.gle/R6faU74qPRPAzchZ9

### 同时读取和写入viper是否安全？

不安全，您需要自己同步对viper的访问（例如使用`sync`包）。并发读取和写入可能会导致恐慌。

## 故障排除

请参阅[TROUBLESHOOTING.md](TROUBLESHOOTING.md)。

## 开发

**为了获得最佳的开发者体验，建议安装[Nix](https://nixos.org/download.html)和[direnv](https://direnv.net/docs/installation.html)。**

_或者，在您的计算机上安装[Go](https://go.dev/dl/)，然后运行`make deps`以安装其余依赖项。_

运行测试套件：

```shell
make test
```

运行linter：

```shell
make lint # 传递-j选项以并行运行它们
```

某些linter违规可以自动修复：

```shell
make fmt
```

## 许可证

该项目根据[MIT许可证](LICENSE)进行许可。