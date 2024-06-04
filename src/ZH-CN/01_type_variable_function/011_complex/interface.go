package main

import "fmt"

// 定义接口类型
type Shape interface {
    Area() float64
}

// 实现Shape接口的矩形结构
type Rectangle struct {
    Width  float64
    Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func main() {
    rect := Rectangle{Width: 5, Height: 3}

    // 使用接口
    var shape Shape = rect
    fmt.Println(shape.Area())
}