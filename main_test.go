package main

import (
	"log"
	"os"
	"testing"

	bookssvc "github.com/andregri/library-microservices/books-svc"
	"gorm.io/gorm"
)

var db *gorm.DB

func TestMain(m *testing.M) {
	ensureTableExists()
	code := m.Run()
	clearTable()
	os.Exit(code)
}

func ensureTableExists() {
	db = initDb()
}

func clearTable() {
	db.Migrator().DropTable(&bookssvc.Book{})
}

func TestA(t *testing.T) {
	log.Println("TestA running")
}
