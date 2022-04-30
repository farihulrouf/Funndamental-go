package main
import (
	"net/http"
	"github.com/gin-gonic/gin"
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
type BoolInput struct {
	Title string
	Price int
	SubTitle string
}
func postBookHandler(c *gin.Context) {
	var bookInput BoolInput
	err := c.ShouldBindJSON(&bookInput)
	if err != nil {
		//log.Fatal(err)
	}
	c.JSON(http.StatusOK, gin.H {
		"title": bookInput.Title,
		"price": bookInput.Price,
		"SubTitle": bookInput.SubTitle,
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