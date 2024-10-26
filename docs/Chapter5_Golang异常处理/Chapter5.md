# Chapter5: golang异常处理

本章节我们主要学习golang异常处理的内容，包括：

- 什么是异常
- defer的妙用
- recover与panic异常捕获


通过本章节的学习，我们可以更好的了解golang异常处理的机制，提高代码的健壮性。

## 5.1 什么是异常

在编程中，异常是指在程序运行过程中发生的非正常情况或错误。异常处理是编程语言提供的一种机制，用于捕获和处理这些错误，从而提高程序的健壮性和可靠性。如果程序没有对异常进行处理，那么程序就会崩溃。下面我们使用一个例子做说明：

➡️src/Chapter5_Golang异常处理/demo1/main.go

```go
package main

import "fmt"

func devide(a, b int) int {
    return a / b
}

func main () {
    res := devide(1, 0)
    fmt.Printf("devide(1, 0) = %d\n", res)
}
```

我们在控制台运行如下语句会得到以下结果：

```shell
$ go run ./src/Chapter5_Golang异常处理/demo1/main.go 
panic: runtime error: integer divide by zero

goroutine 1 [running]:
main.devide(...)
        E:/work/datawhale/golang_datawhale/golang_datawhale/src/Chapter5_Golang异常处理/demo1/main.go:6
main.main()
        E:/work/datawhale/golang_datawhale/golang_datawhale/src/Chapter5_Golang异常处理/demo1/main.go:10 +0xa
exit status 2
```

程序在做`devide(1, 0)`时，由于除数为0，导致程序崩溃。程序崩溃是一件非常严重的事情，因此我们需要对异常进行处理。

在Golang中，异常处理主要通过 `error` 类型和 `panic` 机制来实现。Golang没有传统的 `try-catch` 异常处理机制，而是通过返回错误值和使用 `defer`、`panic`、`recover` 关键字来处理异常。

下面我们来看看如何使用`error`类型以及返回值来处理异常。

➡️src/Chapter5_Golang异常处理/demo2/main.go

```go
package main

import (
    "errors"
    "fmt"
)

func devide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

func main() {
    res, err := devide(1, 0)
    if err != nil {
        fmt.Printf("devide(1, 0) error: %s\n", err)
    } else {
        fmt.Printf("devide(1, 0) = %d\n", res)
    }
}
```

在上述程序中，我们做了如下操作：

- 定义错误：使用 errors.New 函数创建一个新的错误对象。
- 返回错误：在 divide 函数中，如果除数为0，则返回错误对象。
- 检查错误：在 main 函数中，调用 divide 函数并检查返回的错误对象。如果发生错误，则打印错误信息。

## 5.2 defer的妙用

`defer` 语句用于延迟函数的执行，直到包含 `defer` 语句的函数执行完毕。`defer` 语句通常用于资源清理、文件关闭、解锁等操作。我们来看一段包含多个defer语句的代码，请注意defer语句的执行顺序。

➡️src/Chapter5_Golang异常处理/demo3/main.go

```go
package main

import "fmt"

func main() {
    fmt.Println("Start of main function")

    defer fmt.Println("Deferred call 1")
    defer fmt.Println("Deferred call 2")
    defer fmt.Println("Deferred call 3")

    fmt.Println("End of main function")
}
```

在上述程序中，我们在main函数中使用了三个defer语句，分别打印 "Deferred call 1"、"Deferred call 2" 和 "Deferred call 3"。当程序运行时，defer语句会按照后进先出（LIFO）的顺序执行，因此输出结果如下：

```shell
$ go run ./src/Chapter5_Golang异常处理/demo3/main.go
Start of main function
End of main function
Deferred call 3
Deferred call 2
Deferred call 1
```

下图可以解释上述程序的执行过程：

- 开始执行 `main` 函数：打印 `Start of main function`。
- 注册 `defer` 语句：
  - `defer fmt.Println("Deferred call 1")` 被注册，但不会立即执行。
  - `defer fmt.Println("Deferred call 2")` 被注册，但不会立即执行。
  - `defer fmt.Println("Deferred call 3")` 被注册，但不会立即执行。
- 继续执行 `main` 函数：打印 `End of main function`。
- 执行 `defer` 语句：
  - `defer` 语句按照后进先出的顺序执行。
  - 首先执行 `defer fmt.Println("Deferred call 3")`，打印 `Deferred call 3`。
  - 然后执行 `defer fmt.Println("Deferred call 2")`，打印 `Deferred call 2`。
  - 最后执行 `defer fmt.Println("Deferred call 1")`，打印 `Deferred call 1`。

```md
+-----------------------------------+
| Start of main function            |
|                                   |
| +-----------------------------+   |
| | defer fmt.Println("Deferred |   |
| | call 1")                    |   |
| +-----------------------------+   |
|                                   |
| +-----------------------------+   |
| | defer fmt.Println("Deferred |   |
| | call 2")                    |   |
| +-----------------------------+   |
|                                   |
| +-----------------------------+   |
| | defer fmt.Println("Deferred |   |
| | call 3")                    |   |
| +-----------------------------+   |
|                                   |
| End of main function              |
|                                   |
| +-----------------------------+   |
| | Deferred call 3             |   |
| +-----------------------------+   |
|                                   |
| +-----------------------------+   |
| | Deferred call 2             |   |
| +-----------------------------+   |
|                                   |
| +-----------------------------+   |
| | Deferred call 1             |   |
| +-----------------------------+   |
+-----------------------------------+
```

defer 语句在函数结束时按照后进先出的顺序执行，通常用于资源清理、文件关闭、解锁等操作。掌握 defer 语句的使用，可以提高代码的健壮性和可维护性。

在这里给大家说几个defer的注意事项：

（1）被 `defer` 的函数的参数会在 `defer` 声明时求值（而不是在函数实际执行时）。

➡️src/Chapter5_Golang异常处理/demo4/main.go

```go
package main

import "fmt"

func main() {
    var i int = 1
    defer fmt.Println("defer i = ", i)
    i++
}
```

当我们在控制台中运行上述程序时，会得到如下结果：

```shell
$ go run ./src/Chapter5_Golang异常处理/demo4/main.go 
defer i =  1
```

如何解释如上代码呢？这是由于 `defer` 语句在声明时捕获了变量 `i` 的当前值，而不是在函数执行时。因此，即使在 `defer` 语句之后对i进行了递增操作，defer语句仍然打印了 `i` 的初始值。

我们再看另一个更容易混淆的例子：

➡️src/Chapter5_Golang异常处理/demo5/main.go

```go
package main

import "fmt"


func main() {
    var i int = 1
    defer func() {
        fmt.Println("defer i = ", i)
    }()
    i++
}
```

当我们在控制台中运行上述程序时，会得到如下结果：

```shell
$ go run ./src/Chapter5_Golang异常处理/demo5/main.go 
defer i =  2
```

在这个案例中，`defer` 语句中的你匿名函数捕获了变量 `i` 的引用，而不是值。因此，即使在 `defer` 语句之后对 `i` 进行了递增操作，`defer` 语句仍然打印了 `i` 的最终值，但此时 `i` 的值为2。

（2）避免在 `for` 中使用 `defer`：

需要值得关注的是，被 `defer` 的调用会在包含的函数的末尾执行，而不是包含代码块的末尾（如：for）。在 `for` 循环中使用 `defer` 语句时，需要注意以下几点：

- 资源积累：每次循环迭代都会创建一个新的 `defer` 语句，这可能会导致资源积累，直到函数返回时才会释放。例如，如果在循环中打开文件并使用 `defer` 关闭它们，所有文件会在函数结束时才关闭，而不是在每次迭代结束时关闭。
- 顺序问题：`defer` 语句按照后进先出的顺序执行。如果在循环中使用多个 `defer` 语句，它们的执行顺序会与声明顺序相反。

以下是一个示例，展示了在for循环中使用defer可能导致的问题：

```go
package main

import "fmt"

func main() {
    for i := 0; i < 3; i++ {
        defer fmt.Println("defer i =", i)
    }
}
```

运行上述代码时，输出结果为：

```shell
defer i = 2
defer i = 1
defer i = 0
```

解决方法: 如果需要在每次迭代结束时执行某些操作，可以将这些操作放在循环体内，而不是使用 `defer` 。例如：

```go
package main

import "fmt"

func main() {
    for i := 0; i < 3; i++ {
        fmt.Println("i =", i)
        // 其他需要在每次迭代结束时执行的操作
    }
}
```

这样可以确保每次迭代结束时立即执行所需的操作，而不是等到函数返回时才执行。

## 5.3 recover与panic异常捕获

在以上案例中，如果我们提前对除数为`0`的场景做出处理，我们可以避免程序崩溃。在Golang中，`panic` 用于引发异常，`recover` 用于捕获异常。`panic` 通常用于不可恢复的错误，而 `recover` 则用于从 `panic` 中恢复。下面我们来看看如何使用 `defer` 和 `recover` 来处理异常。

➡️src/Chapter5_Golang异常处理/demo6/main.go

```go
package main

import "fmt"

// 定义一个函数，使用 panic 引发异常
func devide(a, b int) int {
    if b == 0 {
        panic("division by zero")
    }
    return a / b
}

func main() {
    defer func() {
        if err := recover(); err != nil {
            fmt.Println("panic occurred:", err)
        }
    }()
    res := devide(10, 0)
    fmt.Printf("devide(10, 0) = %d\n", res)
}
```

在上述程序中，我们做了如下操作：

- 定义 `devide` 函数：

  - 该函数接受两个整数参数 `a` 和 `b`，并返回它们的商。
  - 如果 `b` 为 0，则使用 `panic` 引发异常，提示 `"division by zero"`。
  - 否则，返回 a / b 的结果。
- 在 `main` 函数中使用 `defer` 和 `recover`：
  - 使用 `defer` 定义一个匿名函数，该函数在 `main` 函数返回之前执行。
  - 在匿名函数中，使用 `recover` 捕获可能的 `panic`，并打印错误信息。
  - 调用 `devide(10, 0)`，由于 `b` 为 `0`，会引发 `panic`。
  - `recover` 捕获到 `panic` 后，打印 `"panic occurred: division by zero"`。

上述程序可以使用如下图解释：

```md
+-----------------------------------+
| main()                            |
|                                   |
| +-----------------------------+   |
| | defer func() {              |   |
| |     if err := recover();    |   |
| |     err != nil {            |   |
| |         fmt.Println("panic  |   |
| |         occurred:", err)    |   |
| |     }                       |   |
| | }()                         |   |
| +-----------------------------+   |
|                                   |
| +-----------------------------+   |
| | res := devide(10, 0)        |   |
| |                             |   |
| | +-------------------------+ |   |
| | | devide(a, b int)        | |   |
| | | if b == 0 {             | |   |
| | |     panic("division by  | |   |
| | |     zero")              | |   |
| | | }                       | |   |
| | | return a / b            | |   |
| | +-------------------------+ |   |
| +-----------------------------+   |
|                                   |
| +-----------------------------+   |
| | fmt.Printf("devide(10, 0) = |   |
| | %d\n", res)                 |   |
| +-----------------------------+   |
+-----------------------------------+
            |
            v
+-----------------------------------+
| panic: division by zero           |
|                                   |
| +-----------------------------+   |
| | defer func() {              |   |
| |     if err := recover();    |   |
| |     err != nil {            |   |
| |         fmt.Println("panic  |   |
| |         occurred:", err)    |   |
| |     }                       |   |
| | }()                         |   |
| +-----------------------------+   |
|                                   |
|                                   |
| End                               |
|                                   |
|                                   |
+-----------------------------------+
```

- 开始执行 `main` 函数：

  - 打印 `Start of main function`。
- 注册 `defer` 语句：

  - `defer func() { if err := recover(); err != nil { fmt.Println("panic occurred:", err) } }()` 被注册，但不会立即执行。
- 调用 `devide` 函数：

  - 调用 `devide(10, 0)`，进入 `devide` 函数。
- 引发 `panic`：

  - 在 `devide` 函数中，检查到 `b == 0`，引发 `panic("division by zero")`。
- 捕获 `panic`：

  - `panic` 被引发后，程序跳转到 `defer` 注册的匿名函数。
  - 在匿名函数中，使用 `recover` 捕获 `panic`，并将错误信息赋值给 `err`。
- 打印错误信息：

  - 检查到 `err != nil`，打印 `panic occurred: division by zero`。
- 继续执行 `main` 函数：

  - `defer` 语句执行完毕，程序结束。

因此，如果在shell中运行上述程序，会得到如下结果：

```shell
$ go run ./src/Chapter5_Golang异常处理/demo4/main.go
panic occurred: division by zero
```

`panic` 和 `recover` 的区别:

- 触发和处理：`panic` 用于触发异常，而 `Recover` 用于捕获和处理异常。
- 使用位置：`panic` 可以在任何地方触发，但`recover' 只能在延迟函数中使用。
- 效果和行为：`panic` 会立即终止当前函数的执行并展开堆栈，而 `recover` 可以恢复程序的执行并返回 `panic` 的值。
- 使用场景：`panic` 用于处理无法恢复的错误或异常, `recover` 用于防止程序崩溃并采取措施处理异常情况。
