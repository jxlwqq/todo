package todo

import (
	"context"
	v1 "github.com/jxlwqq/todo/api/todo/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Server struct {
	v1.UnimplementedTodoServer
	r Repository
}

func NewServer(r Repository) v1.TodoServer {
	return &Server{
		r: r,
	}
}

func (s Server) Create(ctx context.Context, req *v1.CreateRequest) (*v1.CreateResponse, error) {
	title := req.Item.Title
	description := req.Item.Description
	remindAt := req.Item.RemindAt.AsTime()

	item := Item{
		Title:       title,
		Description: description,
		RemindAt:    remindAt,
	}

	resp := v1.CreateResponse{}

	if err := s.r.Create(&item); err != nil {
		return &resp, nil
	}

	resp.Id = item.ID

	return &resp, nil

}

func (s Server) Update(ctx context.Context, req *v1.UpdateRequest) (*v1.UpdateResponse, error) {
	id := req.Item.Id
	resp := v1.UpdateResponse{}
	if _, err := s.r.Get(id); err != nil {
		resp.Updated = false
		return &resp, err
	}
	title := req.Item.Title
	description := req.Item.Description
	remindAt := req.Item.RemindAt.AsTime()

	item := Item{
		ID:          id,
		Title:       title,
		Description: description,
		RemindAt:    remindAt,
	}

	if err := s.r.Update(&item); err != nil {
		return &resp, nil
	}

	resp.Updated = true

	return &resp, nil
}

func (s Server) Delete(ctx context.Context, req *v1.DeleteRequest) (*v1.DeleteResponse, error) {
	id := req.Id
	resp := v1.DeleteResponse{}
	if _, err := s.r.Get(id); err != nil {
		resp.Deleted = false
		return &resp, err
	}
	if err := s.r.Delete(id); err != nil {
		return &resp, nil
	}

	resp.Deleted = true

	return &resp, nil
}

func (s Server) Get(ctx context.Context, req *v1.GetRequest) (*v1.GetResponse, error) {
	id := req.Id

	resp := v1.GetResponse{}

	item, err := s.r.Get(id)
	if err != nil {
		return &resp, nil
	}

	resp.Item = &v1.Item{
		Id:          item.ID,
		Title:       item.Title,
		Description: item.Description,
		RemindAt:    timestamppb.New(item.RemindAt),
	}

	return &resp, nil
}

func (s Server) List(ctx context.Context, req *v1.ListRequest) (*v1.ListResponse, error) {
	resp := v1.ListResponse{}

	items, err := s.r.List()
	if err != nil {
		return &resp, nil
	}

	for _, item := range items {
		resp.Items = append(resp.Items, &v1.Item{
			Id:          item.ID,
			Title:       item.Title,
			Description: item.Description,
			RemindAt:    timestamppb.New(item.RemindAt),
		})
	}

	return &resp, nil
}
