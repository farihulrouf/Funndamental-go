package main
import (
	"log"
	"fmt"
	"github.com/gin-gonic/gin"
	"pustaka-api/handler"
	"pustaka-api/book"
	"gorm.io/driver/sqlite"
  	"gorm.io/gorm"
)

func main() {

    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Db connection error")
	}
	fmt.Println("Database Connection successed")
	db.AutoMigrate(&book.Book{}) //create Migrate Db model to Database
	// TEST CRUD

	// Create
	/*
	book := book.Book{}
	book.Title = "Node Js"
	book.Price = 120000
	book.Rating = 5
	book.Description = "Buku menjelaskan tentang pemrograman nodejs"

	err = db.Create(&book).Error
	if err != nil {

		log.Fatal("==========================")
		log.Fatal("Error creating book record")
		log.Fatal("==========================")
	}
	*/

	// Read get all
	/*
	var books []book.Book
	err = db.Debug().Find(&books).Error
	if err != nil {

		log.Fatal("==========================")
		log.Fatal("Error read book record")
		log.Fatal("==========================")
	}
	for _, b := range books {	
		//fmt.Println("Title :", book.Title)
		fmt.Println("book object %v ", b)
	}
	*/

	// Find Db
	var book book.Book
	err = db.Debug().Where("id = ?", 1).First(&book).Error
	if err != nil {
		fmt.Println("===========================")
		fmt.Println("Error Finding Record")
		fmt.Println("===========================")
	}

	//deleted
	/*
	err = db.Delete(&book).Error
	if err != nil {
		fmt.Println("===========================")
		fmt.Println("Error Deleting Record")
		fmt.Println("===========================")
	}
	*/

	//Updated
	/*
	book.Title = "Man Triger Revisied Edition"
	err = db.Save(&book).Error
	if err != nil {
		fmt.Println("===========================")
		fmt.Println("Error Update Record")
		fmt.Println("===========================")
	}
	*/

	router := gin.Default()
	router.GET("/", handler.RootHandler)
	router.GET("/six", handler.SixHandler)
	router.GET("/books/:id", handler.BooksHandler)
	router.GET("query/", handler.QueryHandler)
	router.POST("/books", handler.PostBookHandler)
	router.Run(":8111")
}



