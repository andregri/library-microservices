package main

import (
	"log"
	"net/http"

	bookssvc "github.com/andregri/library-microservices/books-svc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=password dbname=library_db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&bookssvc.Book{})

	bookSvc := bookssvc.BookServiceInstance{
		Db: db,
	}

	log.Fatal(http.ListenAndServe(":8080", bookssvc.MakeHandler(bookSvc)))
}
