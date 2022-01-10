package main

import (
	"encoding/xml"

	"github.com/gin-gonic/gin"
)

type Person struct {
	XMLName xml.Name `xml:"person"`
	FirstName string `xml:"firstName,attr"`
	LastName string `xml:"lastName,attr"`
}

func main() {
	router := gin.Default()
	// router.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "hello world",
	// 	})
	// })
	router.GET("/", IndexHandler2)
	router.Run()
}

func IndexHandler(c *gin.Context) {
	name := c.Param("name")
	c.JSON(200, gin.H{
		"message": "hello " + name,
	})
}

func IndexHandler2(c *gin.Context) {
	c.XML(200,Person{FirstName:"Tran Van",
					LastName:"Quac"})
}
