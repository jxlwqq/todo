//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	v1 "github.com/jxlwqq/todo/api/proto/v1"
	"github.com/jxlwqq/todo/internal/todo"
	"github.com/jxlwqq/todo/pkg/dbcontext"
)

func InitTodo(DSN string) (v1.TodoServiceServer, error) {
	wire.Build(todo.NewService, todo.NewRepository, dbcontext.NewDB)
	return todo.Service{}, nil
}
