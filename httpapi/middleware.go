package httpapi

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func ServerHTTPErrorHandler(err error, c echo.Context) {
	he, ok := err.(*echo.HTTPError)
	if ok {
		if he.Internal != nil {
			if herr, ok := he.Internal.(*echo.HTTPError); ok {
				he = herr
			}
		}
	} else {
		he = &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
		}
	}

	code := he.Code
	message := he.Message

	// Send response
	if !c.Response().Committed {
		//err = c.String(code, "")
		log.Println(c.RealIP(), c.Request().RequestURI, code, message)
		err = c.NoContent(he.Code)
	}
}
