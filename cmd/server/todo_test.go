package main

import (
	"context"
	"net/http/httptest"
	"testing"

	"connectrpc.com/connect"
	todov1 "github.com/somasekimoto/connect-go-todo/gen/todo/v1"        // generated by protoc-gen-go
	"github.com/somasekimoto/connect-go-todo/gen/todo/v1/todov1connect" // generated by protoc-gen-co
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestCreateTodo(t *testing.T) {
	t.Parallel()
	mux := server()
	server := httptest.NewUnstartedServer(mux)
	server.EnableHTTP2 = true
	server.StartTLS()
	t.Cleanup(server.Close)
	cases := []struct {
		scenario string
		name     string
		want     *todov1.TodoItem
	}{
		{
			scenario: "Twitterのユーザー名",
			name:     "heacet43",
			want: &todov1.TodoItem{
				Id:        1,
				Name:      "heacet43",
				Completed: wrapperspb.Bool(false),
			},
		},
		{
			scenario: "GitHubのユーザー名",
			name:     "Hirochon",
			want: &todov1.TodoItem{
				Id:        1,
				Name:      "Hirochon",
				Completed: wrapperspb.Bool(false),
			},
		},
		{
			scenario: "今の気持ち",
			name:     "お腹すいた",
			want: &todov1.TodoItem{
				Id:        1,
				Name:      "お腹すいた",
				Completed: wrapperspb.Bool(false),
			},
		},
	}
	for _, c := range cases {
		c := c
		t.Run(c.scenario, func(t *testing.T) {
			t.Parallel()
			client := todov1connect.NewTodoServiceClient(
				server.Client(),
				server.URL,
			)
			res, err := client.CreateTodo(context.Background(), connect.NewRequest(&todov1.CreateTodoRequest{
				Name: c.name,
			}))
			if err != nil {
				t.Error(err)
			}
			got := res.Msg.GetItem()
			if got.Id != c.want.Id || got.Name != c.want.Name || got.Completed.GetValue() != c.want.Completed.GetValue() {
				t.Errorf("todo got: %+v, want: %+v", got, c.want)
			}
		})
	}
}

func TestUpdateTodo(t *testing.T) {
	t.Parallel()
	mux := server()
	server := httptest.NewUnstartedServer(mux)
	server.EnableHTTP2 = true
	server.StartTLS()
	t.Cleanup(server.Close)
	cases := []struct {
		scenario string
		name     string
		want     *todov1.TodoItem
	}{
		{
			scenario: "Twitterのユーザー名",
			name:     "heacet43",
			want: &todov1.TodoItem{
				Id:        1,
				Name:      "heacet43",
				Completed: wrapperspb.Bool(false),
			},
		},
		{
			scenario: "GitHubのユーザー名",
			name:     "Hirochon",
			want: &todov1.TodoItem{
				Id:        1,
				Name:      "Hirochon",
				Completed: wrapperspb.Bool(false),
			},
		},
		{
			scenario: "今の気持ち",
			name:     "お腹すいた",
			want: &todov1.TodoItem{
				Id:        1,
				Name:      "お腹すいた",
				Completed: wrapperspb.Bool(false),
			},
		},
	}
	for _, c := range cases {
		c := c
		t.Run(c.scenario, func(t *testing.T) {
			t.Parallel()
			client := todov1connect.NewTodoServiceClient(
				server.Client(),
				server.URL,
			)
			res, err := client.UpdateTodo(context.Background(), connect.NewRequest(&todov1.UpdateTodoRequest{
				Name: c.name,
			}))
			if err != nil {
				t.Error(err)
			}
			got := res.Msg.GetItem()
			if got.Id != c.want.Id || got.Name != c.want.Name || got.Completed.GetValue() != c.want.Completed.GetValue() {
				t.Errorf("todo got: %+v, want: %+v", got, c.want)
			}
		})
	}
}

func TestDeleteTodo(t *testing.T) {
	t.Parallel()
	mux := server()
	server := httptest.NewUnstartedServer(mux)
	server.EnableHTTP2 = true
	server.StartTLS()
	t.Cleanup(server.Close)
	cases := []struct {
		scenario string
		id       uint64
		want     *todov1.TodoItem
	}{
		{
			scenario: "Twitterのユーザー名",
			id:       1,
			want: &todov1.TodoItem{
				Id:        1,
				Name:      "heacet43",
				Completed: wrapperspb.Bool(false),
			},
		},
		{
			scenario: "GitHubのユーザー名",
			id:       2,
			want: &todov1.TodoItem{
				Id:        2,
				Name:      "Hirochon",
				Completed: wrapperspb.Bool(false),
			},
		},
		{
			scenario: "今の気持ち",
			id:       3,
			want: &todov1.TodoItem{
				Id:        3,
				Name:      "お腹すいた",
				Completed: wrapperspb.Bool(false),
			},
		},
	}
	for _, c := range cases {
		c := c
		t.Run(c.scenario, func(t *testing.T) {
			t.Parallel()
			client := todov1connect.NewTodoServiceClient(
				server.Client(),
				server.URL,
			)
			res, err := client.DeleteTodo(context.Background(), connect.NewRequest(&todov1.DeleteTodoRequest{
				Id: c.id,
			}))
			if err != nil {
				t.Error(err)
			}
			got := res.Msg.Id
			if got != c.want.Id {
				t.Errorf("todo got: %+v, want: %+v", got, c.want)
			}
		})
	}
}
