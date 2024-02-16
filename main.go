package main

import (
  "context"
  
  "github.com/labstack/echo/v4"
  "github.com/apodicticscott/BRUTAL/templates"
)

func main()  {
  e := echo.New()
  
  // Main Menu
  component := templates.Index()
  e.GET("/", func (c echo.Context) error  {
    return component.Render(context.Background(), c.Response().Writer)
  })
  e.Static("/static", "static")
  e.Static("/css", "css")
  e.Logger.Fatal(e.Start(":3000"))
}
