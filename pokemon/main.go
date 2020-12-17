package main

import "github.com/gofiber/fiber/v2"
import "github.com/gofiber/template/html"

func main() {
	engine := html.New("./templates", ".html")
    app := fiber.New(fiber.Config{
		Views: engine,
	})

    app.Get("/", func(c *fiber.Ctx) error {
        return c.Render("index", fiber.Map{
			"Title": "Hello, World!",
		})
    })
    app.Listen(":3000")
}


// package main

// import (
// 	"net/http"
// 	"github.com/gin-gonic/gin"
	
// )
// var router *gin.Engine

// func main(){
// 	gin.SetMode(gin.ReleaseMode)
// 	router = gin.Default()
// 	router.LoadHTMLGlob("templates/*")
// 	router.GET("/", func(c *gin.Context){
// 		c.HTML(
// 			http.StatusOK,
// 			"index.html",
// 			gin.H{
// 				"title" : "Home Page",
// 			},
// 		)
// 	})
// 	router.Run(":8080")
// }