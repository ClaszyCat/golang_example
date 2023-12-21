package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Product struct{
	ID string `json:"id"`
 	Name string `json:"nama"`
	Stock int `json:"stok"`
}

// type Student struct{
// 	NIM string `json:"nim"`
// 	FirstName string
// 	LastName string
// 	Class string
// 	Age int
// }

var products = make (map[string]Product)

func getAllProducts(c echo.Context) error{
	productsList :=make ([]Product, 0, len(products))
	for _,v := range products{
		productsList = append(productsList, v)
	}
	return c.JSON(http.StatusOK, productsList)
}

func getProducts(c echo.Context) error{
	id := c.Param("id")
	product, found := products[id]
	if !found{
		return c.String(http.StatusNotFound, "Product not found")
	}
	return c.JSON(http.StatusOK,product)
}

func updateProducts(c echo.Context) error{
	id := c.Param("id")
	if _, found := products[id]; !found{
		return c.String(http.StatusNotFound, "Product not found")
	}

	updateProducts := new(Product)
	if err := c.Bind(updateProducts); err !=nil{
		return err
	}
	products[id]= *updateProducts
	return c.JSON(http.StatusOK, "Product Updated")
}

func createProducts(c echo.Context) error{
	newProducts :=new(Product)
	fmt.Println(c)
	if err := c.Bind(newProducts); err !=nil{
		return err
	}

	if _, found := products[newProducts.ID]; found{
		return c.String (http.StatusBadRequest,"Product already exists")
	}
	products[newProducts.ID]= *newProducts
	return c.JSON(http.StatusCreated, newProducts)
}

func deleteProduct(c echo.Context) error{
	id := c.Param("id")
	if _, found := products[id]; !found{
		return c.String(http.StatusNotFound, "Product not found")
	}
	delete(products, id)
	return c.JSON(http.StatusOK, "Data Adios")
}

func main()  {

	// e := echo.New()
	// Yovi := Student{
	// 	NIM: "424", 
	// 	FirstName: "Yovi",
	// 	LastName: "Meong",
	// 	Class: "I",
	// 	Age: 21,
	// }

	// e.GET("/awal/", func(c echo.Context) error {return c.JSON(http.StatusOK,Yovi)})
	// e.GET("/", func(c echo.Context) error {return c.String(http.StatusOK,"Halaman Awal")})
	// // e.GET("/hello", func(c echo.Context)error{return c.String(http.StatusOK,"Hello, There!")})
	// e.GET("/hello",hello)
	// e.GET("/bye", func (c echo.Context)error{return c.String(http.StatusOK,"Goodbye!")}) 
	// e.Start("localhost:1234")

	// e.Logger.Fatal(e.Start("localhost:1234"))

	e:= echo.New()
	e.GET("/products", getAllProducts)
	e.GET("/products/:id", getProducts)
	e.POST("/products", createProducts)
	e.PUT("/products/:id", updateProducts)
	e.DELETE("/products/:id", deleteProduct)

	e.Logger.Fatal(e.Start("localhost:2345"))
}

// func hello(c echo.Context) error{ return c.String(http.StatusOK, "Hello There!")}