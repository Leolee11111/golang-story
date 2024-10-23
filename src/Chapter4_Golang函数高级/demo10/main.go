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