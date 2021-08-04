package main

import (
    "fmt"
    "github.com/kataras/iris/v12"
)

func main() {
    app := iris.New()
    app.Get("/", func(ctx iris.Context) {
        ctx.JSON(iris.Map{"message":"good"})
    })
    fmt.Println("hello world")
    app.Run(iris.Addr(":9000"))
}
