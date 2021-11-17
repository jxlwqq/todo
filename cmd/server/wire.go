//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	v1 "github.com/jxlwqq/todo/api/todo/v1"
	"github.com/jxlwqq/todo/internal/todo"
	"github.com/jxlwqq/todo/pkg/dbcontext"
)

func InitTodoServer(DSN string) (v1.TodoServer, error) {
	wire.Build(todo.NewServer, todo.NewRepository, dbcontext.NewDB)
	return todo.Server{}, nil
}
