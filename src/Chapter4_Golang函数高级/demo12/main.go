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