package main

import (
    // "database/sql"
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
    // _ "github.com/mattn/go-sqlite3"
)

func main(){
    e:= echo.New()

    //middleware
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    //definindo as rotas HTTP
    e.GET("/polls", func(c echo.Context) error{
        return c.JSON(200, "PUT Polls")
    })

    e.PUT("/polls", func(c echo.Context) error{
        return  c.JSON(200, "PUT Polls")
    })

    e.PUT("/polls", func(c echo.Context) error{
        return c.JSON(200, "UPDATE Poll " + c.Param("id"))
    })

    //startando servidor
    e.Logger.Fatal(e.Start(":9000"))
}