package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	bookssvc "github.com/andregri/library-microservices/books-svc"
)

var app App

func TestMain(m *testing.M) {
	app = App{}
	app.Initialize()
	code := m.Run()
	clearTable()
	os.Exit(code)
}

func clearTable() {
	app.Db.Migrator().DropTable(&bookssvc.Book{})
}

func TestGetNonExistentBook(t *testing.T) {
	clearTable()

	req, _ := http.NewRequest("GET", "/book/1", nil)
	rr := httptest.NewRecorder()
	app.Router.ServeHTTP(rr, req)
	t.Errorf("Expected %q, got %q\n", "", rr.Body)
}
