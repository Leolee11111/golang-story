# Time包

## Time包介绍

Time包提供了时间和日期的功能，包括获取当前时间、时间格式化、时间计算等。

## 常用例子
### Time.Time类型
1. 获取当前时间：
```go
now := time.Now()
fmt.Println("当前时间:", now)
//获取Unix时间
nowUnix := now.Unix()
fmt.Println("当前Unix时间:", nowUnix)
```

2. 时间格式化：
```go
// 时间格式化的定义由来
// 在Go语言中，时间格式化使用一个特定的时间模板 "2006-01-02 15:04:05"。
// 这个模板是基于一个特定的时间点，代表了2006年1月2日15时04分05秒。
// 这个时间点是Go语言的设计者选择的，用于表示时间格式的基准。

formattedTime := now.Format("2006-01-02 15:04:05")
fmt.Println("格式化时间:", formattedTime)
```

3. 时间计算：
```go
// 计算10小时后的时间
futureTime := now.Add(10 * time.Hour)
fmt.Println("10小时后的时间:", futureTime)

// 计算10小时之前的时间
pastTime := now.Add(-10 * time.Hour)
fmt.Println("10小时前的时间:", pastTime)

// 计算两个时间之间的间隔
duration := futureTime.Sub(now)
fmt.Println("当前时间到10小时后的时间间隔:", duration)
//https://play.golang.org/p/GBcTqhIFIw9
// AddDate 对时间进行加减操作，代码分享如上。
```

4. 时间解析：
```go
// 解析字符串为时间
parsedTime, err := time.Parse("2006-01-02 15:04:05", "2023-10-01 12:00:00")
if err != nil {
    fmt.Println("解析时间出错:", err)
} else {
    fmt.Println("解析后的时间:", parsedTime)
}
```

# 定时器与超时处理

## Time.Ticker 周期定时器

```go
// time.Ticker 是一个周期性定时器，每隔指定的时间间隔（此处为1秒）发送一个时间信号。
// 通过 `ticker.C` 可以接收这些信号，适合用于周期性的定时任务。
ticker := time.NewTicker(1 * time.Second)
defer ticker.Stop()

go func() {
    for t := range ticker.C {
        fmt.Println("Tick at", t)
    }
}()
```

## Time.Timer 一次性定时器

```go
// time.Timer 用于设置一个一次性的延时操作，
// 在指定的时间（此处为2秒）后，`timer.C` 会接收到一个信号，适合用于超时处理。
timer := time.NewTimer(2 * time.Second)
<-timer.C
fmt.Println("Timer expired")
```

## Time.After 延时触发器

```go
// time.After 是一个简便的方式来创建一个延时触发器，
// 它会在指定的时间（此处为3秒）后发送一个信号，适合用于简单的延时操作。
// 实际上time.After 会创建一个time.Timer 定时器
go func() {
    <-time.After(3 * time.Second)
    fmt.Println("Executed after 3 seconds")
}()
```

# Strings包
1. Contains
   检查字符串是否包含子字符串。
```go
contains := strings.Contains("hello world", "world") // true
```
2. Count
   计算子字符串在字符串中出现的次数。
```go
count := strings.Count("hello world", "l") // 3
```

3. HasPrefix
```go
hasPrefix := strings.HasPrefix("hello world", "hello") // true
```
检查字符串是否以指定的前缀开始。

4. HasSuffix
```go
hasSuffix := strings.HasSuffix("hello world", "world") // true
```
检查字符串是否以指定的后缀结束。
5. Index
```go
index := strings.Index("hello world", "world") // 6
```
返回子字符串在字符串中第一次出现的位置，如果未找到则返回 -1。
6. Join
```go
joined := strings.Join([]string{"hello", "world"}, ", ") // "hello, world"
```
将字符串切片连接成一个字符串，使用指定的分隔符。
7. Replace
```go
replaced := strings.Replace("hello world", "world", "golang", 1) // "hello golang"
```
替换字符串中的子字符串，返回替换后的字符串。
8. Split
   将字符串按照指定的分隔符分割成字符串切片。
```go
split := strings.Split("hello,world", ",") // []string{"hello", "world"}
```
9. ToLower
   将字符串转换为小写
```go
lower := strings.ToLower("HELLO WORLD") // "hello world"
```
10. ToUpper
    将字符串转换为大写。
```go
upper := strings.ToUpper("hello world") // "HELLO WORLD"
```
11. Trim
    去除字符串符串两端的空白字符或指定字符。
```go
trimmed := strings.Trim("  hello world  ", " ") // "hello world"
```
这些是 Strings 包中一些常用的 API，适用于字符串的处理和操作。

# Strconv包
Strconv 实现基本数据类型的字符串表示的转换。

## 常用API

1. **Atoi**
   将字符串转换为整数。
```go
i, err := strconv.Atoi("123") // i = 123, err = nil
```

2. **Itoa**
   将整数转换为字符串。
```go
s := strconv.Itoa(123) // s = "123"
```

3. **FormatFloat**
   将浮点数转换为字符串。
```go
f := 3.14159
s := strconv.FormatFloat(f, 'f', 2, 64) // s = "3.14"
```

4. **ParseFloat**
   将字符串解析为浮点数。
```go
f, err := strconv.ParseFloat("3.14", 64) // f = 3.14, err = nil
```

5. **ParseInt**
   将字符串解析为整数。
```go
i, err := strconv.ParseInt("123", 10, 0) // i = 123, err = nil
```

6. **FormatInt**
   将整数转换为字符串。
```go
s := strconv.FormatInt(123, 10) // s = "123"
```

7. **Quote**
   将字符串转义为安全的字符串格式。
```go
s := strconv.Quote("hello\nworld") // s = "\"hello\\nworld\""
```

8. **Unquote**
   将转义的字符串解析为普通字符串。
```go
s, err := strconv.Unquote("\"hello\\nworld\"") // s = "hello\nworld", err = nil
```

# Reflect包
Go语言中的Reflect包提供了执行反射所需的基本工具。引用go语言圣经这本书的话来讲，`Go 语言提供了一种机制在运行时更新变量和检查它们的值、调用它们的方法，但是在编译时并不知道这些变量的具体类型，这称为反射机制。`这在很多中间件（如json解析，web框架，grpc等）的实现中提供了灵活性和动态性。

1. **TypeOf**
   获取变量的类型。
```go
t := reflect.TypeOf(42) // t = int
```

2. **ValueOf**
   获取变量的值。
```go
v := reflect.ValueOf("hello") // v = "hello"
```

3. **Kind**
   获取类型的种类。
```go
kind := reflect.TypeOf(3.14).Kind() // kind = float64
```

4. **Elem**
   获取指针指向的值。
```go
num := 42
var p *int = &num
v := reflect.ValueOf(p).Elem() // 获取指针p指向的值
```

5. **Field**
   获取结构体字段的值。
```go
type Person struct {
    Name string
    Age  int
}
p := Person{"Alice", 30}
v := reflect.ValueOf(p).Field(0) // v = "Alice"
```

5.1 Field 相关api
```go
// 添加 Field 相关的 API 示例
// 这里是些常用的 Field 相关操作示例
// 获取结构体字段的值
type Person struct {
    Name string
    Age  int
}

p := Person{"Alice", 30}
v := reflect.ValueOf(p).Field(0) // v = "Alice"

// 获取字段的类型
fieldType := reflect.TypeOf(p).Field(0) // fieldType = Name

// 获取字段的标签，struct tag 是go语言通过反射实现的。
tag := fieldType.Tag.Get("json") // tag = "name"

// 设置字段的值（需要可设置的值）
vPtr := reflect.ValueOf(&p).Elem()
vPtr.Field(1).SetInt(31) // 设置Age为31
// 现在p的值为 Person{"Alice", 31}
```
6. **Method**
   获取方法的值。
```go
type MyStruct struct{}
func (m MyStruct) Hello() {
    fmt.Println("Hello")
}
v := reflect.ValueOf(MyStruct{}).Method(0) // 获取Hello方法
```

7. **Call**
   调用方法。
```go
v.Method(0).Call(nil) // 调用Hello方法
```

8. **Set**
   设置值（需要可设置的值）。
```go
v := reflect.ValueOf(&p).Elem()
v.Field(1).SetInt(31) // 设置Age为31
```

9. **Interface**
   获取接口值。
```go
i := v.Interface() // 获取原始值
```

10. **NumField**
    获取结构体字段数量。
```go
num := reflect.TypeOf(p).NumField() // num = 2
```
11. **DeepEqual 和 Copy**
- **DeepEqual**: 用于比较两个值是否深度相等。它会检查值的所有字段和元素适用于复杂数据结构的比较。
```go
a := struct{ Name string }{"Alice"}
b := struct{ Name string }{"Alice"}
equal := reflect.DeepEqual(a, b) // true
```

- **reflect.Copy**: 用于复制值。对于切片、映射和通道等引用类型，Copy 会创建一个新的副本。
```go
import "reflect"

src := []int{1, 2, 3}
dst := make([]int, len(src))
reflect.Copy(reflect.ValueOf(dst), reflect.ValueOf(src)) // dst 现在是 [1, 2, 3]，并且reflect.Copy会返回拷贝的数量
```
这些是 Reflect 包中一些常用的 API，适用于动态类型检查和操作。

