//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/jxlwqq/todo/api/protobuf"
	"github.com/jxlwqq/todo/internal/pkg/dbcontext"
	"github.com/jxlwqq/todo/internal/todo"
)

func InitTodoServer(DSN string) (protobuf.TodoServer, error) {
	wire.Build(todo.NewServer, todo.NewRepository, dbcontext.NewDB)
	return todo.Server{}, nil
}
