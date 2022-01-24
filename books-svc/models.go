package bookssvc

// PostBookRequest is RPC request message to post a new book
type PostBookRequest struct {
	Book Book `json:"book"`
}

// PostBookResponse is a RPC response message to the post request
type PostBookResponse struct {
	ID int `json:"id"`
}

// GetBookRequest is the rpc request message to get a book by id
type GetBookRequest struct {
	ID int `json:"id"`
}

// GetBookResponse is the rpc response of the get request
type GetBookResponse struct {
	Book Book `json:"book"`
}

// PutBookRequest is the rpc request to update an existing book
type PutBookRequest struct {
	ID   int  `json:"id"`
	Book Book `json:"book"`
}

// DeleteBookRequest is the rpc request to delete a book record
type DeleteBookRequest struct {
	ID int `json:"id"`
}

// ErrorResponse is sent when the endpoint could not complete the request
type ErrorResponse struct {
	Error string `json:"error"`
}
