package handler
import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"pustaka-api/book"
)

func PostBookHandler(c *gin.Context) {
	var bookInput book.BookInput
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
func QueryHandler(c *gin.Context) {
	title := c.Query("title")
	c.JSON(http.StatusOK, gin.H{"title": title})

}

func BooksHandler(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"id": id})

}

func RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H {
		"name": "Farihul Rouf",
		"bio": "A Software engineer & content creator",
	})
}

func SixHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H {
		"name": "Farihul Rouf",
		"bio": "A Front End engineer & content creator",
		"city": "Surabaya",
		"country": "Indonesia",
	})
}