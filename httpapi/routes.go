package httpapi

import "github.com/labstack/echo/v4"

func AddRoutes(e *echo.Echo) {

	e.Static("/", "static") //配置静态文件路径

	e.GET("/reg", Reg)
	e.GET("/add", Add)
	e.GET("/rm", Rm)
	//e.POST("/sub", Submit)
}
