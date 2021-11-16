package todo

import (
	"context"
	v1 "github.com/jxlwqq/todo/api/proto/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Service struct {
	v1.UnimplementedTodoServiceServer
	r Repository
}

func NewService(r Repository) v1.TodoServiceServer {
	return &Service{
		r: r,
	}
}

func (s Service) Create(ctx context.Context, req *v1.CreateRequest) (*v1.CreateResponse, error) {
	title := req.Todo.Title
	description := req.Todo.Description
	remindAt := req.Todo.RemindAt.AsTime()

	todo := Todo{
		Title:       title,
		Description: description,
		RemindAt:    remindAt,
	}

	resp := v1.CreateResponse{}

	if err := s.r.Create(&todo); err != nil {
		return &resp, nil
	}

	resp.Id = todo.ID

	return &resp, nil

}

func (s Service) Update(ctx context.Context, req *v1.UpdateRequest) (*v1.UpdateResponse, error) {
	id := req.Todo.Id
	title := req.Todo.Title
	description := req.Todo.Description
	remindAt := req.Todo.RemindAt.AsTime()

	todo := Todo{
		ID:          id,
		Title:       title,
		Description: description,
		RemindAt:    remindAt,
	}

	resp := v1.UpdateResponse{}

	if err := s.r.Update(&todo); err != nil {
		return &resp, nil
	}

	resp.Updated = true

	return &resp, nil
}

func (s Service) Delete(ctx context.Context, req *v1.DeleteRequest) (*v1.DeleteResponse, error) {
	id := req.Id

	resp := v1.DeleteResponse{}

	if err := s.r.Delete(id); err != nil {
		return &resp, nil
	}

	resp.Deleted = true

	return &resp, nil
}

func (s Service) Get(ctx context.Context, req *v1.GetRequest) (*v1.GetResponse, error) {
	id := req.Id

	resp := v1.GetResponse{}

	todo, err := s.r.Get(id)
	if err != nil {
		return &resp, nil
	}

	resp.Todo = &v1.Todo{
		Id:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
		RemindAt:    timestamppb.New(todo.RemindAt),
	}

	return &resp, nil
}

func (s Service) List(ctx context.Context, req *v1.ListRequest) (*v1.ListResponse, error) {
	resp := v1.ListResponse{}

	todos, err := s.r.List()
	if err != nil {
		return &resp, nil
	}

	for _, todo := range todos {
		resp.Todos = append(resp.Todos, &v1.Todo{
			Id:          todo.ID,
			Title:       todo.Title,
			Description: todo.Description,
			RemindAt:    timestamppb.New(todo.RemindAt),
		})
	}

	return &resp, nil
}
