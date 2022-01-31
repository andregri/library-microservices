package main

import (
	"fmt"
	"net/http"
	"os"

	bookssvc "github.com/andregri/library-microservices/books-svc"
	"github.com/go-kit/log"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type App struct {
	Db     *gorm.DB
	Router *mux.Router
	Logger log.Logger
}

func main() {
	app := App{}
	app.Initialize()
	app.Run()
}

func (a *App) Initialize() {

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)

	var err error
	a.Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	a.Db.AutoMigrate(&bookssvc.Book{})

	bookSvc := bookssvc.BookServiceInstance{
		Db: a.Db,
	}

	a.Router = mux.NewRouter()
	a.Logger = log.NewLogfmtLogger(os.Stderr)

	a.Router.Handle("/", bookssvc.MakeHandler(bookSvc, a.Logger))
}

func (a *App) Run() {
	a.Logger.Log(http.ListenAndServe(":8080", a.Router))
}
