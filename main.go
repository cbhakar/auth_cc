package main

import (
"github.com/labstack/echo"
	"net/http"
)


// using hard coded users instead of DB
var users = []string {"rajat", "chandra"}

func main() {
	e := echo.New()
	e.GET("/auth",authHandler)
	e.Logger.Fatal(e.Start(":9991"))
}

func authHandler(c echo.Context) error {
	u := c.QueryParam("q")
	if len(u) < 1 {
		return c.JSON(http.StatusBadRequest, map[string]string{"err":"header 'Username' can't be nil"})
	}
	for _, user := range users{
		if u == user{
			return c.JSON(http.StatusOK, map[string]string{"status":"ok"})
		}
	}
	return c.JSON(http.StatusUnauthorized, map[string]string{"status":"err"})
}
