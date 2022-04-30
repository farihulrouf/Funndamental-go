package main
import (
	"github.com/gin-gonic/gin"
	"pustaka-api/handler"
)

func main() {
	router := gin.Default()
	router.GET("/", handler.RootHandler)
	router.GET("/six", handler.SixHandler)
	router.GET("/books/:id", handler.BooksHandler)
	router.GET("query/", handler.QueryHandler)
	router.POST("/books", handler.PostBookHandler)
	router.Run(":8111")
}



