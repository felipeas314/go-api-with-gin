package main

import (
	"github.com/gin-gonic/gin"
	"github.com/felipeas314/go/api/models"
	"github.com/felipeas314/go/api/controllers"
)

func main() {
   r := gin.Default()

   models.ConnectDatabase()
   
   r.GET("/books", controllers.FindBooks)

   r.Run()
}