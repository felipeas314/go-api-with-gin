package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/felipeas314/go/handlers"
	"github.com/felipeas314/go/database"
)

func Setup() *gin.Engine {
	r := gin.Default()
	api := &handlers.APIEnv{
		DB: database.GetDB(),
	}

	r.GET("", api.GetBooks)

	return r
}