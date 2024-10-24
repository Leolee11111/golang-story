# Chapter4: golang函数高级

本章节我们主要学习golang函数高级的内容，包括：

- 函数与闭包
- 函数与递归
- 接口类型
- golang与面向对象
- 鸭子类型

通过本章节的学习，您将对golang函数有一个更深入的了解。我们知道，和Python语言类似，Golang中的函数不仅仅是代码的基本构建块，它们还可以作为一等公民，支持高阶函数和闭包等高级特性。这些特性使得Golang在处理复杂逻辑和实现函数式编程时非常强大和灵活。

## 1. 函数与闭包

### 1.1 函数基础

函数是Golang中的基本构建块，用于代码重用和模块化。一个函数的基本结构包括函数名、参数列表、返回值类型和函数体。下面是一个简单的函数示例：

➡️src/demo1/main.go

```go
package main

import "fmt"

// 定义一个简单的函数: 函数名add 参数列表(a, b int) 返回值类型int
func add(a int, b int) int {
    return a + b
}

func main() {
    result := add(1, 2)
    fmt.Println("Result: ", result) // 输出：Result: 3
}
```

在终端中运行上述代码，可以看到输出结果为`Result: 3`。这里我们定义了一个简单的函数`add`，接收两个整型参数`a`和`b`，返回它们的和。

### 1.2 高阶函数

高阶函数是指接受一个或多个函数作为参数，或返回一个函数的函数。在Golang中，函数是一等公民，称为"一等公民"（First-Class Citizen）是因为它们具有与其他基本数据类型（如整数、字符串等）相同的地位和操作能力。具体来说，函数作为一等公民意味着它们可以：

- 赋值给变量：函数可以像其他数据类型一样赋值给变量。
- 作为参数传递：函数可以作为参数传递给其他函数。
- 作为返回值：函数可以作为其他函数的返回值。
- 存储在数据结构中：函数可以存储在数组、切片、映射等数据结构中。

这些特性使得函数在编程中非常灵活和强大，特别是在实现高阶函数和函数式编程时。这点跟Python语言是一样的，这使得高阶函数成为可能。那返回函数的函数——高阶函数有什么实际用途呢？

- 代码重用：通过将通用逻辑抽象为高阶函数，可以减少代码重复。
- 函数式编程：高阶函数是函数式编程的核心概念，允许更灵活和简洁的代码。
- 回调函数：高阶函数可以用于实现回调机制，处理异步操作或事件驱动编程。

（1）代码复用：使用高阶函数实现不同的数学运算

我们将定义一个高阶函数`operate`，它接受两个整数和一个操作函数作为参数，并返回操作函数的结果。然后，我们可以定义不同的操作函数，如加法、减法、乘法和除法，并将它们传递给`operate`函数。

➡️src/demo2/main.go

```go
package main

import "fmt"

// 定义一个函数类型
type opFunc func(int, int) int

// 定义加法函数
func add(a int, b int) int {
    return a + b
}

// 定义减法函数
func sub(a int, b int) int {
    return a - b
}

// 定义乘法函数
func mul(a int, b int) int {
    return a * b
}

// 定义除法函数
func div(a int, b int) int {
    if b == 0 {
        fmt.Println("Error: Division not be zero")
        return 0
    }
    return a / b
}

// 定义一个高阶函数，接受函数作为参数
func operate(a int, b int, op opFunc) int {
    return op(a, b)
}

func main() {
    a, b := 15, 5

    fmt.Println("Add:", operate(a, b, add))         // 输出：Add: 20
    fmt.Println("Subtract:", operate(a, b, sub)) // 输出：Subtract: 10
    fmt.Println("Multiply:", operate(a, b, mul)) // 输出：Multiply: 75
    fmt.Println("Divide:", operate(a, b, div))     // 输出：Divide: 3
}
```

在上述代码中，我们定义了四个不同的操作函数`add`、`sub`、`mul`和`div`，它们分别实现了加法、减法、乘法和除法运算。然后，我们定义了一个高阶函数`operate`，它接受两个整数和一个操作函数作为参数，并返回操作函数的结果。最后，我们调用`operate`函数，传递不同的操作函数，实现不同的数学运算。

（2）函数式编程：函数式编程是一种编程范式，它强调使用函数来处理数据和操作。高阶函数是函数式编程的核心概念之一，它允许我们将函数作为参数传递或返回，从而实现更灵活和简洁的代码。在函数式编程中有几个经久不衰的案例：Map函数、Filter函数和Reduce函数。

- Map函数：对列表中的每个元素应用一个函数，并返回结果列表。

```md
输入列表: [1, 2, 3, 4, 5]
          |   |   |   |   |
          v   v   v   v   v
        +---+---+---+---+---+
        | f | f | f | f | f |
        +---+---+---+---+---+
          |   |   |   |   |
          v   v   v   v   v
输出列表: [f(1), f(2), f(3), f(4), f(5)]
```

下面是`map`函数的实现案例：`map`函数接受一个函数和一个列表，并返回一个新的列表，其中每个元素都是通过将函数应用于原列表中的每个元素得到的。

➡️src/demo3/main.go

```go
package main

import "fmt"

// 定义一个函数类型
type transFunc func(int) int

// 定义一个map函数: 输入数据的切片和一个变换函数，返回变换后的切片
func mapFunc(arr []int, f transFunc) []int {
    result := make([]int, len(arr))
    for i, v := range arr {
        result[i] = f(v)
    }
    return result
}

// 定义一个简单的变换函数: y = x^2
func square(x int) int {
    return x * x
}

func main() {
    arr := []int{1, 2, 3, 4, 5}
    squaredRes := mapFunc(arr, square)
    fmt.Println(squaredRes) // 输出：[1 4 9 16 25]
}
```

- Filter函数：从列表中选择满足条件的元素，并返回结果列表。如：`[1, 2, 3, 4, 5, 6]`中选择所有偶数元素为`[2, 4, 6]`。

```md
输入列表: [1, 2, 3, 4, 5, 6]
          |   |   |   |   |   |
          v   v   v   v   v   v
        +---+---+---+---+---+---+
        | p | p | p | p | p | p |
        +---+---+---+---+---+---+
          |   |   |   |   |   |
          v   v   v   v   v   v
        [1] [2] [3] [4] [5] [6]
          |   |   |   |   |   |
          v   v   v   v   v   v
        +---+---+---+---+---+---+
        | F | T | F | T | F | T |
        +---+---+---+---+---+---+
              |       |       |
              v       v       v
输出列表: [2, 4, 6]
```

下面是`filter`函数的实现案例：输入数据的切片slice和一个谓词函数(判断是否为偶数)，返回满足条件的切片slice：

➡️src/demo4/main.go

```go
package main

import "fmt"

// 定义一个函数类型
type predicateFunc func(int) bool

// 定义一个filter函数: 输入数据的切片和一个谓词函数，返回满足条件的切片
func filter(arr []int, f predicateFunc) []int {
    result := make([]int, 0)
    for _, v := range arr {
        if f(v) {
            result = append(result, v)
        }
    }
    return result
}

// 定义一个简单的谓词函数: 判断是否为偶数
func isEven(x int) bool {
    return x % 2 == 0
}

func main() {
    arr := []int{1, 2, 3, 4, 5, 6}
    evenArr := filter(arr, isEven)
    fmt.Println(evenArr) // 输出：[2, 4, 6]
}
```

- Reduce函数：对列表中的元素进行累积操作，返回一个值。如：`[1, 2, 3, 4, 5]`的累积和为`15`。

```md
输入列表: [1, 2, 3, 4, 5]
          |   |   |   |   |
          v   v   v   v   v
初始值:   0
          |
          v
累积过程:  0 + 1 = 1
          1 + 2 = 3
          3 + 3 = 6
          6 + 4 = 10
          10 + 5 = 15
          |
          v
输出值:   15
```

下面是`reduce`函数的实现案例：输入数据的切片和一个累积函数，返回累积后的结果：

➡️src/demo5/main.go

```go
package main

import "fmt"

// 定义一个函数类型
type reduceFunc func(int, int) int

// 定义一个reduce函数: 输入数据的切片、一个累积函数和一个初始值，返回累积的结果
func reduce(arr []int, f reduceFunc, initial int) int {
    result := initial
    for _, v := range arr {
        result = f(result, v)
    }
    return result
}

// 定义一个简单的累积函数
func sum(x, y int) int {
    return x + y
}

func main() {
    arr := []int{1, 2, 3, 4, 5}
    total := reduce(arr, sum, 0)
    fmt.Println(total) // 输出：15
}
```

（3）回调函数：高阶函数可以用于实现回调机制，处理异步操作或事件驱动编程。在这里，我们还没有讨论Go语言的并发编程，但是回调函数是Go语言中处理并发操作的重要机制之一。通过将函数作为参数传递给其他函数，我们可以实现异步操作、事件处理和并发控制等功能。后续在学习完并发编程以后，再来深入了解回调函数的应用。

### 1.3 匿名函数

匿名函数
匿名函数是没有名字的函数。它们可以在函数内部定义并立即调用，或者赋值给变量以便稍后调用。匿名函数在需要一次性使用的情况下非常有用，特别是在回调函数和闭包中。

匿名函数的特性：

- 没有名字：匿名函数没有名字，因此不能通过名字调用。
- 可以立即调用：匿名函数可以在定义时立即调用。
- 可以赋值给变量：匿名函数可以赋值给变量，以便稍后调用。
- 可以作为参数传递：匿名函数可以作为参数传递给其他函数。

➡️src/demo6/main.go

```go
package main

import "fmt"

func main() {
    // 1.定义匿名函数并立即执行
    result1 := func (a int, b int) int {
        return a + b
    }(1, 2)
    fmt.Println("result1: ", result1)

    // 2.将匿名函数赋值给变量
    addFunc := func (a int, b int) int {
        return a + b
    }
    result2 := addFunc(1, 2)
    fmt.Println("result2: ", result2)

    // 3.匿名函数作为参数传递
    result3 := calc(1, 2, func(a int, b int) int {
        return a + b
    })
    fmt.Println("result3: ", result3)
}

func calc(a int, b int, op func(int, int) int) int {
 return op(a, b)
}
```

### 1.4 闭包

闭包（Closure）是指一个函数与其引用的外部变量环境的组合。闭包允许函数在其词法作用域外调用时，仍然能够访问并操作其词法作用域内的变量。换句话说，闭包可以“记住”并访问创建它的作用域中的变量。

一个闭包是一个函数值，它引用了其函数体之外的变量。这个函数可以访问并赋值这个引用的变量；换句话说，这个函数被“绑定”在这个变量上。

以下是一个简单的闭包示例：

➡️src/demo7/main.go

```go
package main

import "fmt"

// 定义一个外部函数，返回一个闭包函数
func adder() func(int) int {
    sum := 0
    return func(x int) int {
        sum += x
        return sum
    }
}

func main() {
    // 调用外部函数，返回闭包函数
    pos := adder()

    // 多次调用闭包函数
    fmt.Println(pos(1)) // 输出：1
    fmt.Println(pos(2)) // 输出：3
    fmt.Println(pos(3)) // 输出：6
}
```

为了帮助理解闭包的工作原理,以下是一个闭包流程图，展示了闭包如何捕获并访问其外部作用域中的变量，以及为什么变量的值会累计下来.

```md
+-------------------------+
| 外部函数                |
|                         |
| func adder() func(int) int { |
|     sum := 0            |  <--- 外部变量 sum
|     return func(x int) int { |
|         sum += x        |  <--- 捕获并操作 sum
|         return sum      |
|     }                   |
| }                       |
+-------------------------+
            |
            v
+-------------------------+
| 闭包函数                |
|                         |
| func(x int) int {       |
|     sum += x            |  <--- 持续操作 sum
|     return sum          |
| }                       |
+-------------------------+
            |
            v
+-------------------------+
| 调用闭包函数            |
|                         |
| pos := adder()          |  <--- 创建闭包实例
| fmt.Println(pos(1))     |  <--- sum = 0 + 1 = 1
| fmt.Println(pos(2))     |  <--- sum = 1 + 2 = 3
| fmt.Println(pos(3))     |  <--- sum = 3 + 3 = 6
+-------------------------+
```

为什么变量的值会累计下来？在闭包中，匿名函数捕获了外部函数中的局部变量`sum`。即使外部函数`adder`已经返回，闭包函数仍然可以访问并操作`sum`变量。每次调用闭包函数时，`sum`的值都会更新并持久化。这是因为**闭包函数持有对`sum`变量的引用，而不是对其值的拷贝**。因此，`sum`变量的状态在多次调用之间保持一致，导致其值会累计下来。

那闭包函数实际上有什么实际的应用场景呢？闭包函数在Go语言中有很多实际的应用场景，特别是在需要持久化状态或创建函数工厂时。

（1）持久化状态：缓存是一种存储机制，用于临时存储计算结果，以便在后续请求中快速返回结果。
（2）创建函数工厂：函数工厂是一种设计模式，它使用闭包来生成具有特定行为的函数。函数工厂可以根据不同的参数生成不同的函数实例，这些实例可以共享相同的逻辑，但具有不同的初始状态或配置。以下是一个详细的案例，展示了如何使用闭包创建一个函数工厂，并应用于统计不同商品的销售数量：

现有两种商品x和y，每次购买商品时，计数器会自动增加。我们可以使用闭包函数来实现这个计数器，从而计算每种商品的销售数量，进而判断商品的热度。

➡️src/demo8/main.go

```go
package main

import "fmt"

type counterFunc func() int

func main() {
    goods := []string{"x", "y", "x", "x", "x", "y", "x"}
    xCountVal, yCountVal := 0, 0
    goodsCounterMap := map[string]counterFunc {
        "x": counter(),
        "y": counter(),
    }
    for _, good := range goods {
        if good == "x" {
            xCountVal = goodsCounterMap[good]()
        }
        if good == "y" {
            yCountVal = goodsCounterMap[good]()
        }
    }
    fmt.Println("xCountVal:", xCountVal)
    fmt.Println("yCountVal:", yCountVal)
}

func counter() counterFunc {
    x := 0
    return func() int {
        x += 1
        return x
    }
}
```

## 2. 函数与递归

### 2.1 什么是递归？

递归是指函数调用自身的一种编程技巧。递归通常用于解决分治问题，如计算阶乘、斐波那契数列、树的遍历等。递归函数必须包含一个基准条件（base case），用于终止递归调用。下面，我们以计算阶乘为例，介绍递归的基本概念。
阶乘是一个经典的递归问题。阶乘的定义如下：

- $0!=1$
- $n!=n *(n-1)!$

➡️src/demo9/main.go

```go
package main

import "fmt"

// 定义一个递归函数计算阶乘
func factorial(n int) int {
    if n == 0 {
        return 1
    }
    return n * factorial(n-1)
}

func main() {
    fmt.Println(factorial(5)) // 输出：120
}
```

1. 初始调用：当我们调用 factorial（5）时，函数会检查 n 是否等于 0 。如果不是，则执行 return n * factorial(n-1)。

2. 递归展开: 函数会继续调用自身，直到基准条件 $n=0$ 满足为止。递归调用的展开过程如下:

- factorial(5) 调用 factorial(4)
- factorial(4) 调用 factorial(3)
- factorial(3) 调用 factorial(2)
- factorial(2) 调用 factorial(1)
- factorial(1) 调用 factorial(0)

1. 基准条件: 当 factorial(0) 被调用时，基准条件 $\mathrm{n}==0$ 满足，函数返回 1 。

2. 递归回归：函数开始回归，每一层调用都会返回计算结果，直到初始调用 factorial(5) 返回最终结果。回归过程如下:

- factorial(0) 返回 1
- factorial(1) 返回 1 * $1=1$
- factorial(2) 返回 $2 * 1=2$
- factorial(3) 返回 $3 * 2=6$
- factorial(4) 返回 $4 * 6=24$
- factorial(5) 返回 $5 * 24=120$

下图展示了递归函数 factorial 如何计算阶乘，以及递归调用和回归的过程：

```md
调用链: factorial(5)
+-------------------------+
| factorial(5)            |
| if n == 0 { return 1 }  |  <--- 基准条件不满足
| return 5 * factorial(4) |  <--- 递归调用
+-------------------------+
            |
            v
+-------------------------+
| factorial(4)            |
| if n == 0 { return 1 }  |  <--- 基准条件不满足
| return 4 * factorial(3) |  <--- 递归调用
+-------------------------+
            |
            v
+-------------------------+
| factorial(3)            |
| if n == 0 { return 1 }  |  <--- 基准条件不满足
| return 3 * factorial(2) |  <--- 递归调用
+-------------------------+
            |
            v
+-------------------------+
| factorial(2)            |
| if n == 0 { return 1 }  |  <--- 基准条件不满足
| return 2 * factorial(1) |  <--- 递归调用
+-------------------------+
            |
            v
+-------------------------+
| factorial(1)            |
| if n == 0 { return 1 }  |  <--- 基准条件不满足
| return 1 * factorial(0) |  <--- 递归调用
+-------------------------+
            |
            v
+-------------------------+
| factorial(0)            |
| if n == 0 { return 1 }  |  <--- 基准条件满足，返回 1
+-------------------------+
            |
            v
回归链: 返回值
+-------------------------+
| factorial(1)            |
| return 1 * 1 = 1        |  <--- 返回 1
+-------------------------+
            |
            v
+-------------------------+
| factorial(2)            |
| return 2 * 1 = 2        |  <--- 返回 2
+-------------------------+
            |
            v
+-------------------------+
| factorial(3)            |
| return 3 * 2 = 6        |  <--- 返回 6
+-------------------------+
            |
            v
+-------------------------+
| factorial(4)            |
| return 4 * 6 = 24       |  <--- 返回 24
+-------------------------+
            |
            v
+-------------------------+
| factorial(5)            |
| return 5 * 24 = 120     |  <--- 返回 120
+-------------------------+
```

### 2.2 递归的优缺点

优点:

- 简洁性：递归函数通常比迭代函数更简洁，代码更易读。
- 自然性：递归函数更符合某些问题的自然定义，如树的遍历和分治算法。

缺点:

- 性能问题：递归函数可能会导致大量的重复计算，影响性能。
- 栈溢出：递归调用会占用栈空间，递归深度过大会导致栈溢出。

## 3. 接口类型

### 3.1 什么是接口？

在现实生活中，接口的概念无处不在。例如：

- 电源插座：电源插座是一个接口，任何符合插座标准的电器都可以插入并使用电源。
- USB接口：USB接口是一个标准接口，任何符合USB标准的设备都可以连接到计算机并进行数据传输。
- 遥控器：遥控器是一个接口，任何符合遥控器信号标准的设备都可以被遥控器控制。

通过这些现实中的例子，可以更直观地理解接口的概念和作用。接口定义了一组行为，而不关心这些行为是如何实现的。任何实现了接口的类型都可以通过接口进行操作，从而实现代码的解耦和多态。

想象一下，我们有一个家庭，里面有各种各样的电器，如电视、空调、洗衣机等。每种电器都有不同的功能，但它们都有一个共同的行为：开机和关机。我们可以定义一个接口来表示这种共同的行为。

➡️src/demo10/main.go

```go
package main

import "fmt"

// 定义一个接口
type Appliance interface {
    TurnOn()
    TurnOff()
}

// 定义一个结构体表示电视
type TV struct{}

// 实现接口方法
func (t TV) TurnOn() {
    fmt.Println("TV is now ON")
}

func (t TV) TurnOff() {
    fmt.Println("TV is now OFF")
}

// 定义另一个结构体表示空调
type AirConditioner struct{}

// 实现接口方法
func (a AirConditioner) TurnOn() {
    fmt.Println("Air Conditioner is now ON")
}

func (a AirConditioner) TurnOff() {
    fmt.Println("Air Conditioner is now OFF")
}

func main() {
    // 创建一个家用电器的切片
    appliances := []Appliance{TV{}, AirConditioner{}}

    // 遍历所有电器并调用它们的开关机方法
    for _, appliance := range appliances {
        appliance.TurnOn()
        appliance.TurnOff()
    }
}
```

通过这个UML类图和代码示例，可以更直观地理解接口的定义和实现，以及接口在多态中的作用:

```md
+---------------------+
|     <<interface>>   |
|      Appliance      |
+---------------------+
| + TurnOn()          |
| + TurnOff()         |
+---------------------+
          ^
          |
+---------------------+       +---------------------+
|        TV           |       |   AirConditioner    |
+---------------------+       +---------------------+
| + TurnOn()          |       | + TurnOn()          |
| + TurnOff()         |       | + TurnOff()         |
+---------------------+       +---------------------+
```

1. 定义接口: type Appliance interface \{ TurnOn(); TurnOff() \} 定义了一个接口 Appliance, 包含两个方法 TurnOn 和 TurnOff
2. 实现接口：TV 和 AirConditioner 结构体分别实现了 TurnOn 和 TurnOff 方法，因此它们都实现了 Appliance 接口。
3. 多态: 在 main 函数中，我们创建了一个 Appliance 类型的切片，包含 TV 和 AirConditioner 实例。通过遍历切片，我们可以调用每个电器的 TurnOn 和 TurnOff 方法，而不需要关心具体的电器类型。这点我们在面向对象中将会再次提到。

### 3.2 为什么面向接口编程？

接口在编程中有许多重要的优点，以下是一些关键的优点，解释为什么要面向接口编程：

1. 解耦和模块化：

接口将行为的定义与实现分离，使得代码更加模块化和易于维护。通过接口，调用者不需要知道具体的实现细节，只需要知道接口定义的行为。

```go
package main

import "fmt"

// 定义一个接口
type Appliance interface {
    TurnOn()
    TurnOff()
}

// 定义一个结构体表示电视
type TV struct{}

// 实现接口方法
func (t TV) TurnOn() {
    fmt.Println("TV is now ON")
}

func (t TV) TurnOff() {
    fmt.Println("TV is now OFF")
}

// 定义另一个结构体表示空调
type AirConditioner struct{}

// 实现接口方法
func (a AirConditioner) TurnOn() {
    fmt.Println("Air Conditioner is now ON")
}

func (a AirConditioner) TurnOff() {
    fmt.Println("Air Conditioner is now OFF")
}

func main() {
    // 创建一个家用电器的切片
    appliances := []Appliance{TV{}, AirConditioner{}}

    // 遍历所有电器并调用它们的开关机方法
    for _, appliance := range appliances {
        appliance.TurnOn()
        appliance.TurnOff()
    }
}
```

在这个示例中，接口 `Appliance` 将 `TurnOn` 和 `TurnOff` 方法的定义与具体实现分离，使得代码更加模块化和易于维护。

2. 多态和灵活性：

接口允许不同类型实现相同的接口，从而在运行时表现出不同的行为。这种特性称为多态。多态使得代码更加灵活，可以处理不同类型的对象，而不需要知道它们的具体类型。

```go
package main

import "fmt"

// 定义一个接口
type Speaker interface {
    Speak() string
}

// 定义一个结构体
type Dog struct{}

// 实现接口方法
func (d Dog) Speak() string {
    return "Woof!"
}

// 定义另一个结构体
type Cat struct{}

// 实现接口方法
func (c Cat) Speak() string {
    return "Meow!"
}

func main() {
    var s Speaker

    s = Dog{}
    fmt.Println(s.Speak()) // 输出：Woof!

    s = Cat{}
    fmt.Println(s.Speak()) // 输出：Meow!
}
```

在这个示例中，接口 `Speaker` 允许 `Dog` 和 `Cat` 实现相同的方法 `Speak`，从而在运行时表现出不同的行为。

3. 可替换性和可扩展性：

接口使得代码更加灵活，可以轻松地替换或扩展实现，而不需要修改接口的定义。通过接口，可以在不改变现有代码的情况下添加新的实现。

```go
package main

import "fmt"

// 定义一个接口
type Logger interface {
    Log(message string)
}

// 定义一个结构体表示控制台日志
type ConsoleLogger struct{}

// 实现接口方法
func (c ConsoleLogger) Log(message string) {
    fmt.Println("Console Log:", message)
}

// 定义另一个结构体表示文件日志
type FileLogger struct{}

// 实现接口方法
func (f FileLogger) Log(message string) {
    fmt.Println("File Log:", message)
}

func main() {
    var logger Logger

    logger = ConsoleLogger{}
    logger.Log("This is a console log message.") // 输出：Console Log: This is a console log message.

    logger = FileLogger{}
    logger.Log("This is a file log message.") // 输出：File Log: This is a file log message.
}
```

在这个示例中，接口 `Logger` 允许我们轻松地替换或扩展日志的实现，而不需要修改接口的定义。

4. 代码复用和抽象：

接口允许我们编写更加通用和灵活的代码，通过接口类型参数实现代码复用。通过接口，可以编写适用于多种类型的通用函数和方法。

```go
package main

import "fmt"

// 定义一个接口
type Shape interface {
    Area() float64
}

// 定义一个结构体
type Circle struct {
    Radius float64
}

// 实现接口方法
func (c Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}

// 定义另一个结构体
type Rectangle struct {
    Width, Height float64
}

// 实现接口方法
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// 计算总面积
func totalArea(shapes ...Shape) float64 {
    var area float64
    for _, shape := range shapes {
        area += shape.Area()
    }
    return area
}

func main() {
    c := Circle{Radius: 5}
    r := Rectangle{Width: 4, Height: 6}
    fmt.Println("Total Area:", totalArea(c, r)) // 输出：Total Area: 103.5
}
```

在这个示例中，接口 `Shape` 允许我们编写一个通用的 `totalArea` 函数，可以接受任意实现了 `Shape` 接口的类型。

## 4. golang与面向对象

面向对象编程（Object-Oriented Programming，简称OOP）是一种编程范式，它通过将数据和操作数据的方法封装在一起，形成对象，从而实现代码的模块化和重用性。面向对象编程的核心概念包括类、对象、封装、继承和多态。虽然Golang不是传统的面向对象编程语言，但它提供了结构体和接口，使得我们可以实现面向对象的设计模式。

### 4.1 类和对象

类：类是对象的蓝图或模板，定义了对象的属性和行为。在Golang中，类的概念由结构体（struct）实现。

对象：对象是类的实例，包含具体的数据和方法。

想象一下，类是一个建筑蓝图，而对象是根据这个蓝图建造的具体房子。每个房子都有自己的颜色、大小和房间数量，但它们都遵循同一个蓝图。结构体就是Golang中的蓝图，而结构体实例就是具体的房子。

➡️src/demo11/main.go

```go
package main

import "fmt"

// 定义一个结构体
type Person struct {
    Name string
    Age  int
}

// 定义一个方法
func (p Person) Greet() {
    fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}

func main() {
    // 创建一个结构体实例
    p := Person{Name: "Alice", Age: 30}
    p.Greet() // 输出：Hello, my name is Alice and I am 30 years old.
}
```

以上代码可以理解为：

（1）定义结构体：

```go
type Person struct {
    Name string
    Age  int
}
```

`Person` 是一个结构体类型，类似于面向对象编程中的类。
结构体 Person 有两个字段：`Name` 和 `Age`，分别表示人的名字和年龄。

(2) 定义方法：

```go
func (p Person) Greet() {
    fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}
```

`Greet` 是一个方法，绑定到 Person 类型。
方法的接收者是 `p Person`，表示 `Greet` 方法可以访问 `Person` 结构体的字段。
`Greet` 方法使用 `fmt.Printf` 打印出 `Person` 的名字和年龄。

(3) 创建结构体实例：

```go
func main() {
    p := Person{Name: "Alice", Age: 30}
    p.Greet() // 输出：Hello, my name is Alice and I am 30 years old.
}
```

在 `main` 函数中，我们创建了一个 `Person` 结构体的实例 `p`，并初始化 `Name` 为 `"Alice"`，`Age` 为 `30`。
调用 `p.Greet()` 方法，输出 `Hello, my name is Alice and I am 30 years old.`。

通过这个代码示例和解释，可以更直观地理解类和对象的概念，以及如何在Golang中实现这些概念。结构体用于定义类，而结构体实例则是具体的对象。方法绑定到结构体类型，使得对象可以执行特定的行为。

### 4.2 封装

封装是指将数据和操作数据的方法封装在一起，隐藏对象的内部实现细节。在Golang中，封装通过结构体和方法实现。在上述代码中，我们定义了一个 `Person` 结构体和一个 `Greet` 方法，将数据（`Name` 和 `Age`）和操作数据的方法`Greet`封装在一起。

我们来看一个更复杂的示例，展示如何使用结构体和方法实现封装：

➡️src/demo12/main.go

```go
package main

import "fmt"

// 定义一个结构体
type Account struct {
    balance float64
}

// 定义一个方法，增加余额
func (a *Account) Deposit(amount float64) {
    if amount > 0 {
        a.balance += amount
    }
}

// 定义一个方法，获取余额
func (a *Account) GetBalance() float64 {
    return a.balance
}

func main() {
    acc := Account{}
    acc.Deposit(100.50)
    fmt.Println("Balance:", acc.GetBalance()) // 输出：Balance: 100.5
}
```

在这个示例中，我们定义了一个 `Account` 结构体，包含一个 `balance` 字段表示账户余额。我们定义了两个方法：`Deposit` 方法用于存款，`GetBalance` 方法用于获取余额。这样，我们可以通过方法来操作账户的余额，而不需要直接访问 `balance` 字段。

封装就像是一个银行账户，你可以存钱和取钱，但你不能直接访问账户的内部余额。你只能通过银行提供的存取款方法来操作账户。这样可以保护账户的内部状态不被外部直接修改。

### 4.3 继承

继承是指一个类（子类）继承另一个类（父类）的属性和方法。在Golang中，继承通过嵌入式结构体实现。嵌入式结构体是一种结构体嵌套的方式，子结构体可以访问父结构体的字段和方法。

➡️src/demo13/main.go

```go
package main

type Animal struct {
    Name string 
}

func (a Animal) Speak() string {
    return "..."
}

type Dog struct {
    Animal
    Age int
}

func (d Dog) Speak() string {
    return "Wang!"
}

func main() {
    dog := Dog{Animal: Animal{Name: "little dog"}, Age: 3}
    println(dog.Name)
    println(dog.Speak())
}
```

在这个示例中，我们定义了一个 `Animal` 结构体，包含一个 `Name` 字段和一个 `Speak` 方法。我们定义了一个 `Dog` 结构体，嵌入了 `Animal` 结构体，并添加了一个 `Age` 字段。`Dog` 结构体重写了 `Speak` 方法，返回 `Wang!`。

继承就像是孩子继承了父母的特征。比如，父母有黑色的头发，孩子也可能会有黑色的头发。在编程中，子类继承了父类的属性和方法，可以直接使用，也可以重写这些方法。

### 4.4 多态

多态是指同一个接口可以由不同的类实现，从而在运行时表现出不同的行为。在Golang中，多态通过接口实现。

➡️src/demo14/main.go

```go
package main

import "fmt"

type Speaker interface {
    Speak() string
}

type Dog struct {}

func (d Dog) Speak() string {
    return "Wang!"
}

type Cat struct {}

func (c Cat) Speak() string {
    return "Miao!"
}

func main() {
    var s Speaker
    s = Dog{}
    fmt.Println(s.Speak()) // 输出：Wang!

    s = Cat{}
    fmt.Println(s.Speak()) // 输出：Miao!
}
```

在这个示例中，我们定义了一个 `Speaker` 接口，包含一个 `Speak` 方法。我们定义了 `Dog` 和 `Cat` 结构体，分别实现了 `Speak` 方法。在 `main` 函数中，我们创建了一个 `Speaker` 类型的变量 `s`，并分别赋值为 `Dog` 和 `Cat` 结构体实例。通过 `s.Speak()` 方法，我们可以在运行时表现出不同的行为。

多态就像是一个遥控器可以控制不同的设备。你可以用同一个遥控器控制电视、空调和音响，只要这些设备实现了遥控器的接口。在编程中，不同的类实现同一个接口，可以在运行时表现出不同的行为。

## 5. 鸭子类型

鸭子类型（Duck Typing）是一种动态类型的风格，是多态性的一种风格。在鸭子类型中，一个对象的语义完全由它的方法决定。这个概念的名字来源于“走起来像鸭子、叫起来像鸭子，那么它就是鸭子”。

以下是一个解释鸭子类型的图，展示了不同对象如何通过实现相同的方法来表现出相同的行为。

```md
+---------------------+       +---------------------+
|       Dog           |       |       Cat           |
+---------------------+       +---------------------+
| + Speak() string    |       | + Speak() string    |
+---------------------+       +---------------------+
            \                     /
             \                   /
              \                 /
               \               /
                \             /
                 \           /
                  \         /
                   \       /
                    \     /
                     \   /
                      \ /
                +---------------------+
                |     Speaker         |
                +---------------------+
                | + Speak() string    |
                +---------------------+
```

(1) 接口： `Speaker` 接口定义了一个方法 `Speak`，任何实现了这个方法的类型都可以被视为 `Speaker`。

(2) 实现接口的结构体：`Dog` 和 `Cat` 结构体分别实现了 `Speak` 方法，因此它们都实现了 `Speaker` 接口。

(3) 鸭子类型：在鸭子类型中，只要一个对象实现了某个方法，它就可以被视为具有该方法的类型。无论是 `Dog` 还是 `Cat`，只要它们实现了 `Speak` 方法，它们就可以被视为 `Speaker`。

让我们沿用上面的示例，展示如何使用鸭子类型在Golang中实现多态：

```go
package main

import "fmt"

type Speaker interface {
    Speak() string
}

type Dog struct {}

func (d Dog) Speak() string {
    return "Wang!"
}

type Cat struct {}

func (c Cat) Speak() string {
    return "Miao!"
}

func main() {
    var s Speaker
    s = Dog{}
    fmt.Println(s.Speak()) // 输出：Wang!

    s = Cat{}
    fmt.Println(s.Speak()) // 输出：Miao!
}
```

通过这个图和代码示例，可以更直观地理解鸭子类型的概念。鸭子类型是一种动态类型的风格，只要一个对象实现了某个方法，它就可以被视为具有该方法的类型。鸭子类型使得代码更加灵活和通用，允许不同类型的对象在运行时表现出相同的行为。

