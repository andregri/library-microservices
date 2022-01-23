package bookssvc

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//
func DecodePostBookRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request PostBookRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

//
func DecodeGetBookRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request GetBookRequest

	id, ok := mux.Vars(r)["id"]
	if !ok {
		return nil, errors.New("bad route")
	}
	request.ID, _ = strconv.Atoi(id)

	return request, nil
}

//
func DecodePutBookRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request PutBookRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	id, ok := mux.Vars(r)["id"]
	if !ok {
		return nil, errors.New("bad route")
	}
	request.ID, _ = strconv.Atoi(id)

	return request, nil
}

//
func DecodeDeleteBookRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request DeleteBookRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	id, ok := mux.Vars(r)["id"]
	if !ok {
		return nil, errors.New("bad route")
	}
	request.ID, _ = strconv.Atoi(id)

	return request, nil
}

//
func EncodeBookResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
