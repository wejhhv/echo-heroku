package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo"
)

func main() {

	port := os.Getenv("PORT")

	e := echo.New()
	e.GET("/:num", NumMul)
	e.Logger.Fatal(e.Start(":" + port))

}

func NumMul(c echo.Context) error {
	num := c.Param("num")
	return c.String(http.StatusOK, num+num)
}
