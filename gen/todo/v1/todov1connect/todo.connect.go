// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: todo/v1/todo.proto

package todov1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/somasekimoto/connect-go-todo/gen/todo/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// TodoServiceName is the fully-qualified name of the TodoService service.
	TodoServiceName = "todo.v1.TodoService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// TodoServiceCreateTodoProcedure is the fully-qualified name of the TodoService's CreateTodo RPC.
	TodoServiceCreateTodoProcedure = "/todo.v1.TodoService/CreateTodo"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	todoServiceServiceDescriptor          = v1.File_todo_v1_todo_proto.Services().ByName("TodoService")
	todoServiceCreateTodoMethodDescriptor = todoServiceServiceDescriptor.Methods().ByName("CreateTodo")
)

// TodoServiceClient is a client for the todo.v1.TodoService service.
type TodoServiceClient interface {
	CreateTodo(context.Context, *connect.Request[v1.CreateTodoRequest]) (*connect.Response[v1.CreateTodoResponse], error)
}

// NewTodoServiceClient constructs a client for the todo.v1.TodoService service. By default, it uses
// the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewTodoServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) TodoServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &todoServiceClient{
		createTodo: connect.NewClient[v1.CreateTodoRequest, v1.CreateTodoResponse](
			httpClient,
			baseURL+TodoServiceCreateTodoProcedure,
			connect.WithSchema(todoServiceCreateTodoMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// todoServiceClient implements TodoServiceClient.
type todoServiceClient struct {
	createTodo *connect.Client[v1.CreateTodoRequest, v1.CreateTodoResponse]
}

// CreateTodo calls todo.v1.TodoService.CreateTodo.
func (c *todoServiceClient) CreateTodo(ctx context.Context, req *connect.Request[v1.CreateTodoRequest]) (*connect.Response[v1.CreateTodoResponse], error) {
	return c.createTodo.CallUnary(ctx, req)
}

// TodoServiceHandler is an implementation of the todo.v1.TodoService service.
type TodoServiceHandler interface {
	CreateTodo(context.Context, *connect.Request[v1.CreateTodoRequest]) (*connect.Response[v1.CreateTodoResponse], error)
}

// NewTodoServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewTodoServiceHandler(svc TodoServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	todoServiceCreateTodoHandler := connect.NewUnaryHandler(
		TodoServiceCreateTodoProcedure,
		svc.CreateTodo,
		connect.WithSchema(todoServiceCreateTodoMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/todo.v1.TodoService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case TodoServiceCreateTodoProcedure:
			todoServiceCreateTodoHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedTodoServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedTodoServiceHandler struct{}

func (UnimplementedTodoServiceHandler) CreateTodo(context.Context, *connect.Request[v1.CreateTodoRequest]) (*connect.Response[v1.CreateTodoResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("todo.v1.TodoService.CreateTodo is not implemented"))
}
