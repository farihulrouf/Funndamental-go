package main
import (
	"net/http"
    "encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"fmt"
)

func main() {
	router := gin.Default()
	router.GET("/", rootHandler)
	router.GET("/six", sixHandler)
	router.GET("/books/:id", booksHandler)
	router.GET("query/", queryHandler)
	router.POST("/books", postBookHandler)
	router.Run(":8111")
}
type BookInput struct {
	Title string  `json:"title" binding:"required"`
	Price json.Number `json:"price" binding:"required"`
	//SubTitle string
}
func postBookHandler(c *gin.Context) {
	var bookInput BookInput
	err := c.ShouldBindJSON(&bookInput)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			erroMessage := fmt.Sprintf("Error on filed %s, condiiton: %s", e.Field(), e.ActualTag())
			c.JSON(http.StatusBadRequest, erroMessage)
			//fmt.Println(err)
			return
		}
		
	}
	c.JSON(http.StatusOK, gin.H {
		"title": bookInput.Title,
		"price": bookInput.Price,
		//"SubTitle": bookInput.SubTitle,
	})

}

func queryHandler(c *gin.Context) {
	title := c.Query("title")
	c.JSON(http.StatusOK, gin.H{"title": title})

}

func booksHandler(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"id": id})

}

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H {
		"name": "Farihul Rouf",
		"bio": "A Software engineer & content creator",
	})
}

func sixHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H {
		"name": "Farihul Rouf",
		"bio": "A Front End engineer & content creator",
		"city": "Surabaya",
		"country": "Indonesia",
	})
}