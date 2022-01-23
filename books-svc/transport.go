package bookssvc

import (
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func MakeHandler(svc BooksService) http.Handler {
	postBookHandler := httptransport.NewServer(
		MakePostBookEndpoint(svc),
		DecodePostBookRequest,
		EncodeBookResponse,
	)

	getBookHandler := httptransport.NewServer(
		MakeGetBookENdpoint(svc),
		DecodeGetBookRequest,
		EncodeBookResponse,
	)

	putBookHandler := httptransport.NewServer(
		MakePutBookEndpoint(svc),
		DecodePutBookRequest,
		EncodeBookResponse,
	)

	deleteBookHandler := httptransport.NewServer(
		MakeDeleteBookENdpoint(svc),
		DecodeDeleteBookRequest,
		EncodeBookResponse,
	)

	r := mux.NewRouter()

	r.Handle("/book", postBookHandler).Methods("POST")
	r.Handle("/book/{id:[a-zA-Z0-9]*}", getBookHandler).Methods("GET")
	r.Handle("/book/{id:[0-9]*}", putBookHandler).Methods("PUT")
	r.Handle("/book/{id:[0-9]*}", deleteBookHandler).Methods("DELETE")

	return r
}
