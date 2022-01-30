package main

import (
	"fmt"
	"net/http"
	"os"

	bookssvc "github.com/andregri/library-microservices/books-svc"
	"github.com/go-kit/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)
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
