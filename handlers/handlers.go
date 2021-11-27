package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
  	"github.com/jinzhu/gorm"
  	"github.com/felipeas314/go/database"
	"fmt"
)

type APIEnv struct {
	DB *gorm.DB
}

func (a *APIEnv) GetBooks(c *gin.Context) {
	books, err := database.GetBooks(a.DB)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, books)
}