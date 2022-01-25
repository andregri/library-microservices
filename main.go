package main

import (
	"net/http"
	"os"

	bookssvc "github.com/andregri/library-microservices/books-svc"
	"github.com/go-kit/log"
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

	logger := log.NewLogfmtLogger(os.Stderr)

	bookHandler := bookssvc.MakeHandler(bookSvc, logger)

	http.Handle("/", bookHandler)
	logger.Log(http.ListenAndServe(":8080", nil))
}
