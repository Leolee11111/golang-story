# 二、golang基础入门

接下来，我们来了解一下golang的一些基础知识

## 1. 变量与常量

让我们先来了解一下变量与常量的概念

### 1.1 变量

变量可以理解为**一个存储数据的容器**，我们可以在程序运行过程中改变其值。

#### 1.1.1 变量的声明

在我们使用一个变量之前，我们要先对其进行声明。而声明有显式声明以及隐式声明两种方式。

- 显式声明

  ```go
  // 显式声明
  var age int
  fmt.Println(age)
  ```

- 隐式声明

  ```go
  // 隐式声明
  name := "hello world"
  fmt.Println(name)
  ```

- 批量声明

  ```go
  // 批量声明
  var (
      price int    = 25
      thing string = "box"
  )
  fmt.Println(price)
  fmt.Println(thing)
  ```

#### 1.1.2 变量作用域

在声明了变量之后，我们需要搞清楚我们的变量在什么范围内会生效，这就是我们变量的作用域。

根据变量作用域的不同，我们可以将变量分为全局变量以及局部变量。

- 全局变量：在函数外部声明的变量称为全局变量，可以在整个文件中访问。

  ```go
  package main
  
  var globalVar = "I am a global variable"
  
  func main() {
      fmt.Println(globalVar)
  }
  ```

- 局部变量：在函数内部声明的变量称为局部变量，只能在该函数内部访问。

  ```go
  package main
  
  import "fmt"
  
  func main() {
      localVar := "I am a local variable"
      fmt.Println(localVar)
  }
  ```

### 1.2 常量

与可以改变容器内数据的变量相对的，**常量是不可以重新赋值的容器**

#### 1.2.1 常量的声明

在使用常量之前，我们也要对其进行声明。常量的声明是const关键字

- 单个声明

  ```go
  package main
  
  import "fmt"
  
  const pi float64 = 3.14159
  const maxAttempts = 3
  
  func main() {
  	fmt.Println(pi)
  	fmt.Println(maxAttempts)
  }
  ```

- 批量声明

  ```go
  const (
  	count = 10
  	name = "Alice"
  )
  
  func main() {
  	fmt.Println(count)
  	fmt.Println(name)
  }
  ```

#### 1.2.2 常量的作用

使用常量有以下几点优势：

- **提高代码可读性**
  - 使用常量可以提高代码的可读性和可维护性。
- **避免硬编码**
  - 使用常量可以避免在代码中使用硬编码值。

## 2. 基础数据类型

接下来我们来学习一下go语言的基本数据类型。

### 2.1 整型

整形，即我们平时认为的整数。在go语言中，有多种整数类型，分别有不同的表示范围。

| 类型       | 表示范围                                                     |
| ---------- | ------------------------------------------------------------ |
| **uint8**  | 无符号 8 位整型 (0 到 255)                                   |
| **uint16** | 无符号 16 位整型 (0 到 65535)                                |
| **uint32** | 无符号 32 位整型 (0 到 4294967295)                           |
| **uint64** | 无符号 64 位整型 (0 到 18446744073709551615)                 |
| **int8**   | 有符号 8 位整型 (-128 到 127)                                |
| **int16**  | 有符号 16 位整型 (-32768 到 32767)                           |
| **int32**  | 有符号 32 位整型 (-2147483648 到 2147483647)                 |
| **int64**  | 有符号 64 位整型 (-9223372036854775808 到 9223372036854775807) |

### 2.2 浮点型

浮点型，即我们平时所认为的小数。在go语言中，也有多种浮点数类型，分别有不同的表示精度。

| 类型           | 表示精度              |
| -------------- | --------------------- |
| **float32**    | IEEE-754 32位浮点型数 |
| **float64**    | IEEE-754 64位浮点型数 |
| **complex64**  | 32 位实数和虚数       |
| **complex128** | 64 位实数和虚数       |

### 2.3 布尔型

布尔型是用来表示逻辑值的类型，值只可以是常量 true 或者 false，用来表示真和假。

### 2.4 字符串

字符串用于表示文本数据，由一系列字符组成。

### 2.5 类型转换



## 3. 字符串的格式化

接下来我们来讲一下字符串格式化，这在处理用户输入、日志记录和输出时非常常用。Go语言提供了强大的字符串格式化功能，主要通过fmt包来实现。

fmt包中常用Sprintf函数来进行字符串的格式化

- Sprintf函数讲解：

  - 格式：fmt.Sprintf(格式化样式, 参数列表…)

    - **格式化样式**：字符串形式，格式化符号以%开头， %s字符串格式，%d十进制的整数格式。
    - **参数列表**：多个参数以逗号分隔，个数必须与格式化样式中的个数一一对应，否则运行时会报错。

  - 样例：

    ```go
    package main
    
    import "fmt"
    
    func main() {
    	name := "Alice"
    	age := 30
    	message := fmt.Sprintf("Name: %s, Age: %d", name, age)
    	fmt.Println(message)
    }
    ```

- 格式化动词：

  fmt包使用格式化动词来控制输出的格式。以下是一些常用的格式化动词：

  | 格式化动词 |        表示类型        |
  | :--------: | :--------------------: |
  |   **%d**   |       十进制整数       |
  |   **%x**   |      十六进制整数      |
  |   **%f**   |         浮点数         |
  |   **%e**   | 科学记数法表示的浮点数 |
  |   **%s**   |         字符串         |
  |   **%q**   |  双引号括起来的字符串  |
  |   **%v**   |        默认格式        |
  |   **%t**   |         布尔值         |
  |   **%T**   |        值的类型        |
  |   **%%**   |       百分号本身       |

- 其余的常见用法：

  除了基本的格式化动词，fmt包还提供了许多格式化选项，用于进一步控制输出格式

  - 字段宽度

    - **%wd**：指定字段宽度，其中w是一个整数

    - 样例：

      ```go
      number := 42
      fmt.Printf(fmt.Sprintf("|%5d|\n", number))
      fmt.Printf(fmt.Sprintf("|%-5d|\n", number))
      ```

  - 精度

    - **%.wf**：指定浮点数的小数位数，其中w是一个整数。

    - 样例：

      ```go
      number := 3.14159
      fmt.Printf(fmt.Sprintf("%.2f\n", number))
      ```

  - 填充字段

    - **%w.f**：指定填充字符，默认是空格，可以使用#来指定其他字符

    - 样例：

      ```go
      number := 42
      fmt.Printf(fmt.Sprintf("|%05d|\n", number))
      ```

## 4. 条件与循环结构

### 4.1 条件语句

当我们想要在特定条件下执行某段代码时该如何实现呢？这是我们就应该使用条件语句了。

条件语句用于根据不同的条件执行不同的代码块。Go语言提供了 if 和 switch两种主要的条件语句。

#### 4.1.1 if语句

if语句是最常用的条件语句，用于根据布尔表达式的值决定是否执行某个代码块。

- 基本用法

  ```go
  package main
  
  import "fmt"
  
  func main() {
  	age := 25
  
  	if age >= 18 {
  		fmt.Println("You are an adult.")
  	} else {
  		fmt.Println("You are a minor.")
  	}
  }
  ```

- 初始化语句

  if语句可以包含一个初始化语句

  ```go
  package main
  
  import "fmt"
  
  func main() {
  	if age := 25; age >= 18 {
  		fmt.Println("You are an adult.")
  	} else {
  		fmt.Println("You are a minor.")
  	}
  }
  ```



#### 4.1.2 switch语句

switch语句用于根据不同的值执行不同的代码块。它可以更简洁地处理多个条件分支。

- 基本语法

  ```go
  package main
  
  import "fmt"
  
  func main() {
  	day := "Monday"
  
  	switch day {
  	case "Monday":
  		fmt.Println("It's Monday.")
  	case "Tuesday":
  		fmt.Println("It's Tuesday.")
  	case "Wednesday":
  		fmt.Println("It's Wednesday.")
  	default:
  		fmt.Println("It's another day.")
  	}
  }
  ```

- 多个条件分支

  可以在一个case里面列出多个值

  ```go
  package main
  
  import "fmt"
  
  func main() {
      day := "Monday"
  
      switch day {
      case "Monday", "Tuesday", "Wednesday":
          fmt.Println("It's a weekday.")
      case "Thursday", "Friday":
          fmt.Println("It's almost the weekend.")
      case "Saturday", "Sunday":
          fmt.Println("It's the weekend.")
      default:
          fmt.Println("Invalid day.")
      }
  }
  ```

### 4.2 循环结构

当我们想要在特定条件下重复执行某段代码时该如何实现呢？这是我们就应该使用循环语句了。

循环结构用于重复执行某段代码，直到满足某个条件为止。Go语言提供了for和range两种主要的循环结构。

#### 4.2.1 for循环

- 基本用法

  ```go
  package main
  
  import "fmt"
  
  func main() {
  	for i := 0; i < 5; i++ {
  		fmt.Println(i)
  	}
  }
  ```

- 无限循环、

  当省略for循环的条件部分时，可以创建无限循环。

  ```go
  package main
  
  import "fmt"
  
  func main() {
  	for {
  		fmt.Println("This is an infinite loop.")
  	}
  }
  ```

- 不带初始化和条件的for循环

  我们可以省略初始化和条件部分，只保留循环内部更新的部分。

  ```go
  package main
  
  import "fmt"
  
  func main() {
  	i := 0
  	for i < 5 {
  		fmt.Println(i)
  		i++
  	}
  }
  ```

#### 4.2.2 range循环

注意：range循环用于遍历数组、切片、字符串、映射和通道。

- 遍历数组

  ```go
  package main
  
  import "fmt"
  
  func main() {
  	numbers := []int{1, 2, 3, 4, 5}
  
  	for index, value := range numbers {
  		fmt.Printf("Index: %d, Value: %d\n", index, value)
  	}
  }
  ```

我们这个章节以数组为例简单了解一下，更详细的内容我们后面再讲。



### 4.2 循环控制

当我们想要控制循环语句时，比如打断它或者其他操作，该如何实现呢？

Go语言提供了break, continue, 和 goto 语句来控制循环的执行。

#### 4.2.1 break语句

break语句用于立即退出循环。

```go
package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}
}
```

#### 4.2.2 continue语句

continue语句用于**跳过当前循环的剩余部分，继续下一次循环**，即这次循环执行到这里为止，这个循环的后面就不执行了，直接执行下个循环。

```go
package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
}
```

#### 4.2.3 goto语句

goto语句用于无条件跳转到标签所在的位置。（**较少使用，了解即可**）

```go
package main

import "fmt"

func main() {
    i := 0
    for {
        if i > 5 {
            goto END
        }
        fmt.Println(i)
        i++
    }

END:
    fmt.Println("Loop ended.")
}
```



## 5. 函数入门

## 5.1 函数定义

函数是**一段可重用的代码**，用于执行特定的任务。在Go语言中，函数使用func关键字定义。

- 基本语法

  ```go
  func functionName(parameters) returnType {
      // 函数体
  }
  ```

  - functionName：函数的名称
  - parameters：函数的参数列表
  - returnType：函数的返回类型

- 示例

  ```go
  package main
  
  import "fmt"
  
  func greet(name string) string {
  	return "Hello, " + name + "!"
  }
  
  func main() {
  	message := greet("Alice")
  	fmt.Println(message)
  }
  ```



### 5.2 参数

函数**可以接受零个或多个参数**。参数列表中**每个参数都需要指定类型**。

#### 5.2.1 单个参数

5.1的示例就是单个参数的。



#### 5.2.2 多个参数

```go
package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func main() {
	result := add(3, 5)
	fmt.Println(result) // 输出: 8
}
```



#### 5.2.3 可变参数

```go
package main

import "fmt"

func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	result := sum(1, 2, 3, 4, 5)
	fmt.Println(result) // 输出: 15
}
```



### 5.3 返回值

返回值就是函数返回的结果。

函数**可以返回零个或多个值**。返回值的类型需要在函数定义时指定。



#### 5.3.1 单个返回值

5.2.2的示例即为单个返回值的样例。



#### 5.3.2 多个返回值

```go
package main

import "fmt"

func divide(a int, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

func main() {
	q, r := divide(10, 3)
	fmt.Printf("Quotient: %d, Remainder: %d\n", q, r) // 输出: Quotient: 3, Remainder: 1
}
```



以上就是我们golang的基础内容部分了，下面我们来讲golang中的一些数据结构。
