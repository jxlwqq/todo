package todo

import (
	"context"
	"github.com/jxlwqq/todo/api/protobuf"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Server struct {
	protobuf.UnimplementedTodoServer
	r Repository
}

func NewServer(r Repository) protobuf.TodoServer {
	return &Server{
		r: r,
	}
}

func (s Server) Create(ctx context.Context, req *protobuf.CreateRequest) (*protobuf.CreateResponse, error) {
	title := req.Item.Title
	description := req.Item.Description
	remindAt := req.Item.RemindAt.AsTime()

	item := Item{
		Title:       title,
		Description: description,
		RemindAt:    remindAt,
	}

	resp := protobuf.CreateResponse{}

	if err := s.r.Create(&item); err != nil {
		return &resp, status.Errorf(codes.Internal, "create item failed: %v", err)
	}

	resp.Id = item.ID

	return &resp, nil

}

func (s Server) Update(ctx context.Context, req *protobuf.UpdateRequest) (*protobuf.UpdateResponse, error) {
	id := req.Item.Id
	resp := protobuf.UpdateResponse{}
	if _, err := s.r.Get(id); err != nil {
		return &resp, status.Error(codes.NotFound, "item not found")
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
		return &resp, status.Errorf(codes.Internal, "update item failed: %v", err)
	}

	resp.Updated = true

	return &resp, nil
}

func (s Server) Delete(ctx context.Context, req *protobuf.DeleteRequest) (*protobuf.DeleteResponse, error) {
	id := req.Id
	resp := protobuf.DeleteResponse{}
	if _, err := s.r.Get(id); err != nil {
		return &resp, status.Errorf(codes.NotFound, "item not found")
	}
	if err := s.r.Delete(id); err != nil {
		return &resp, status.Errorf(codes.Internal, "delete item failed: %v", err)
	}

	resp.Deleted = true

	return &resp, nil
}

func (s Server) Get(ctx context.Context, req *protobuf.GetRequest) (*protobuf.GetResponse, error) {
	id := req.Id

	resp := protobuf.GetResponse{}

	item, err := s.r.Get(id)
	if err != nil {
		return &resp, status.Errorf(codes.NotFound, "item not found")
	}

	resp.Item = &protobuf.Item{
		Id:          item.ID,
		Title:       item.Title,
		Description: item.Description,
		RemindAt:    timestamppb.New(item.RemindAt),
	}

	return &resp, nil
}

func (s Server) List(ctx context.Context, req *protobuf.ListRequest) (*protobuf.ListResponse, error) {
	resp := protobuf.ListResponse{}

	items, err := s.r.List()
	if err != nil {
		return &resp, status.Errorf(codes.Internal, "list items failed: %v", err)
	}

	for _, item := range items {
		resp.Items = append(resp.Items, &protobuf.Item{
			Id:          item.ID,
			Title:       item.Title,
			Description: item.Description,
			RemindAt:    timestamppb.New(item.RemindAt),
		})
	}

	return &resp, nil
}
