package main

import (
	"context"
	"log"
	"net/http"

	"connectrpc.com/connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	todov1 "github.com/somasekimoto/connect-go-todo/gen/todo/v1"        // generated by protoc-gen-go
	"github.com/somasekimoto/connect-go-todo/gen/todo/v1/todov1connect" // generated by protoc-gen-connect-go
)

type TodoServer struct{}

func (s *TodoServer) CreateTodo(
	ctx context.Context,
	req *connect.Request[todov1.CreateTodoRequest],
) (*connect.Response[todov1.CreateTodoResponse], error) {
	log.Println("Request headers: ", req.Header())
	res := connect.NewResponse(&todov1.CreateTodoResponse{
		Item: &todov1.TodoItem{
			Id:        1,
			Name:      req.Msg.Name,
			Completed: false,
		},
	})
	res.Header().Set("Todo-Version", "v1")
	return res, nil
}

func main() {
	greeter := &TodoServer{}
	mux := http.NewServeMux()
	path, handler := todov1connect.NewTodoServiceHandler(greeter)
	mux.Handle(path, handler)
	http.ListenAndServe(
		"localhost:8080",
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
}
