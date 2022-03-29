package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"regexp"
)

func RewriteURLMiddleWare() echo.MiddlewareFunc {
	return middleware.RewriteWithConfig(middleware.RewriteConfig{
		RegexRules: map[*regexp.Regexp]string{
			//Ex: https://domain/basic/san-pham-gia-5-trieu-16673458?platform=web
			regexp.MustCompile("^/basic/([0-9a-zA-Z\\=-]+)-([0-9]+)$"): "/basic/$1/$2",
			//Ex: https://domain/full/san-pham-gia-5-trieu-16673458?platform=web
			regexp.MustCompile("^/full/([0-9a-zA-Z\\=-]+)-([0-9]+)$"): "/full/$1/$2",
		},
	})
}