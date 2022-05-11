package books

import "github.com/graniticio/granitic/v2/types"

type Book struct {
	UUID   string `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
}

type AddBookRequest struct {
	Name   *types.NilableString `json:"name"`
	Author *types.NilableString `json:"author"`
}
