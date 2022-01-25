package bookssvc

import (
	"net/http"

	"github.com/go-kit/log"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func MakeHandler(svc BooksService, logger log.Logger) http.Handler {

	options := []httptransport.ServerOption{
		httptransport.ServerErrorLogger(logger),
		httptransport.ServerErrorEncoder(EncodeErrorResponse),
	}

	postBookHandler := httptransport.NewServer(
		MakePostBookEndpoint(svc),
		DecodePostBookRequest,
		EncodeBookResponse,
		options...,
	)

	getBookHandler := httptransport.NewServer(
		MakeGetBookENdpoint(svc),
		DecodeGetBookRequest,
		EncodeBookResponse,
		options...,
	)

	putBookHandler := httptransport.NewServer(
		MakePutBookEndpoint(svc),
		DecodePutBookRequest,
		EncodeBookResponse,
		options...,
	)

	deleteBookHandler := httptransport.NewServer(
		MakeDeleteBookENdpoint(svc),
		DecodeDeleteBookRequest,
		EncodeBookResponse,
		options...,
	)

	r := mux.NewRouter()

	r.Handle("/book", postBookHandler).Methods("POST")
	r.Handle("/book/{id:[0-9]*}", getBookHandler).Methods("GET")
	r.Handle("/book/{id:[0-9]*}", putBookHandler).Methods("PUT")
	r.Handle("/book/{id:[0-9]*}", deleteBookHandler).Methods("DELETE")

	return r
}
