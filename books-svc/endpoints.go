package bookssvc

import (
	"context"
	"time"

	"github.com/go-kit/kit/endpoint"
)

// Book represents a single book record
type Book struct {
	ID               int       `json:"id"`
	Title            string    `json:"title"`
	Author           string    `json:"author"`
	AverageRating    float32   `json:"average_rating"`
	ISBN             string    `json:"isbn"`
	ISBN13           string    `json:"isbn13"`
	LanguageCode     string    `json:"language_code"`
	NumPages         int       `json:"num_pages"`
	RatingsCount     int       `json:"ratings_count"`
	TextReviewsCount int       `json:"text_reviews_count"`
	PublicationDate  time.Time `json:"publication_date"`
	Publisher        string    `json:"publisher"`
}

// BookService is a CRUD interface for books
type BooksService interface {
	PostBook(ctx context.Context, b Book) (int, error)
	GetBook(ctx context.Context, id int) (Book, error)
	PutBook(ctx context.Context, id int, b Book) error
	DeleteBook(ctx context.Context, id int) error
}

//
func MakePostBookEndpoint(svc BooksService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(PostBookRequest)
		id, err := svc.PostBook(ctx, req.Book)
		return PostBookResponse{ID: id, Error: err.Error()}, nil
	}
}

//
func MakeGetBookENdpoint(svc BooksService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetBookRequest)
		book, err := svc.GetBook(ctx, req.ID)
		return GetBookResponse{Book: book, Error: err.Error()}, nil
	}
}

//
func MakePutBookEndpoint(svc BooksService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(PutBookRequest)
		err := svc.PutBook(ctx, req.ID, req.Book)
		return PutBookResponse{Error: err.Error()}, nil
	}
}

//
func MakeDeleteBookENdpoint(svc BooksService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteBookRequest)
		err := svc.DeleteBook(ctx, req.ID)
		return DeleteBookResponse{Error: err.Error()}, nil
	}
}
