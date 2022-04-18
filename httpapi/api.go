package httpapi

import (
	"github.com/labstack/echo/v4"
	"log"
)

func EchoGlobalHandler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		//urlPath := c.Request().RequestURI
		return next(c)
	}
}

// Reg http://127.0.0.1:8888/reg
func Reg(c echo.Context) error {
	realIp := c.RealIP()
	log.Println(realIp)
	t := 0
	for _, ip := range Server.botIps {
		if realIp == ip {
			t = t + 1
			break
		}
	}
	if t == 0 {
		return c.String(200, "0")
	} else {
		return c.String(200, "1")
	}
}

// Add http://127.0.0.1:8888/add?ip=127.0.0.1
func Add(c echo.Context) error {
	ip := c.QueryParam("ip")
	Server.botIps = append(Server.botIps, ip)
	log.Println("add", ip)
	return c.String(200, "ok")
}

// Rm http://127.0.0.1:8888/rm?ip=127.0.0.1
func Rm(c echo.Context) error {
	ip := c.QueryParam("ip")
	for i, ipp := range Server.botIps {
		if ip == ipp {
			Server.botIps = append(Server.botIps[:i], Server.botIps[i+1:]...)
		}
	}
	log.Println("remove", ip)
	return c.String(200, "ok")
}
