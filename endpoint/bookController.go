package books

import (
	"context"
	"fmt"

	"github.com/graniticio/granitic/v2/ws"
)

// BookCreateIFace Writable Interface
type BookCreateIFace interface {
	Write(string) (string, error)
}

// BookCreate struct
type BookCreate struct {
	Store BookCreateIFace
}

// ProcessPayload Create Action for book endpoint
func (pC *BookCreate) ProcessPayload(ctx context.Context, req *ws.Request, res *ws.Response, p *AddBookRequest) {
	d := fmt.Sprintf("%s,%s\n", p.Name, p.Author)
	if _, err := pC.Store.Write(d); err != nil {
		res.HTTPStatus = 500
		res.Body = "Something went wrong"
	} else {
		res.HTTPStatus = 201
		res.Body = "Book info was succesfully registered"
	}
}

type BookIndexIFace interface {
	Read() ([][]string, error)
}

type BookIndex struct {
	Store BookIndexIFace
}

// Process Index handler for Book
func (pI *BookIndex) Process(ctx context.Context, req *ws.Request, res *ws.Response) {
	data, err := pI.Store.Read()
	if err != nil {
		res.HTTPStatus = 500
		return
	}
	Book := make([]Book, len(data))
	for i, d := range data {
		Book[i].UUID = d[0]
		Book[i].Name = d[1]
		Book[i].Author = d[2]
	}

	res.Body = Book
	res.HTTPStatus = 200
}
