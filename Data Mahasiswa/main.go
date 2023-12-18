package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Product struct{
	ID string `json:"id"`
 	Name string `json:"nama"`
	Stock int `json:"stock"`
}

type Student struct{
	NIM string `json:"nim"`
	FirstName string
	LastName string
	Class string
	Age int
}


func main()  {

	e := echo.New()
	Yovi := Student{
		NIM: "424", 
		FirstName: "Yovi",
		LastName: "Meong",
		Class: "I",
		Age: 21,
	}

	e.GET("/awal/", func(c echo.Context) error {return c.JSON(http.StatusOK,Yovi)})
	e.GET("/", func(c echo.Context) error {return c.String(http.StatusOK,"Halaman Awal")})
	// e.GET("/hello", func(c echo.Context)error{return c.String(http.StatusOK,"Hello, There!")})
	e.GET("/hello",hello)
	e.GET("/bye", func (c echo.Context)error{return c.String(http.StatusOK,"Goodbye!")}) 
	e.Start("localhost:1234")

	e.Logger.Fatal(e.Start("localhost:1234"))

}

func hello(c echo.Context) error{ return c.String(http.StatusOK, "Hello There!")}