package httpapi

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitEcho() *echo.Echo {
	e := echo.New()

	e.HideBanner = true // 隐藏Banner

	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{Level: 5})) // 开启gzip压缩,开启可能会影响性能
	e.Use(EchoGlobalHandler)                                          // 添加全局过滤函数,执行一些鉴权操作
	e.Use(middleware.BodyLimit("20M"))                                // 限制客户端发送的body数据长度为2兆

	//e.Use(middleware.Logger())                                        // 设置日志
	//e.Use(middleware.Recover())                                       // 拦截panic错误并且在控制台打印错误日志，避免echo程序直接崩溃

	e.HTTPErrorHandler = ServerHTTPErrorHandler // 自定义错误,echo自带的404错误处理默认返回json,影响服务器性能

	AddRoutes(e) // 添加路由

	return e
}
