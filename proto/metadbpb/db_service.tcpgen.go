// Code generated by protoc-gen-tcpgen
// source: db_service.proto
// DO NOT EDIT!

/*
Package metadbpb is a http proxy.
*/

package metadbpb

import (
	"context"

	"github.com/gogo/protobuf/proto"
	"github.com/qiwitech/tcprpc"
)

func RegisterMetaDBServiceHandlers(s *tcprpc.Server, prefix string, srv MetaDBServiceInterface) {

	s.Handle(prefix+"Put", tcprpc.NewHandler(
		func() proto.Message { return new(PutRequest) },
		func(ctx context.Context, inp proto.Message) (proto.Message, error) {
			args := inp.(*PutRequest)
			return srv.Put(ctx, args)
		}))

	s.Handle(prefix+"Get", tcprpc.NewHandler(
		func() proto.Message { return new(GetRequest) },
		func(ctx context.Context, inp proto.Message) (proto.Message, error) {
			args := inp.(*GetRequest)
			return srv.Get(ctx, args)
		}))

	s.Handle(prefix+"GetMulti", tcprpc.NewHandler(
		func() proto.Message { return new(GetMultiRequest) },
		func(ctx context.Context, inp proto.Message) (proto.Message, error) {
			args := inp.(*GetMultiRequest)
			return srv.GetMulti(ctx, args)
		}))

	s.Handle(prefix+"Delete", tcprpc.NewHandler(
		func() proto.Message { return new(DeleteRequest) },
		func(ctx context.Context, inp proto.Message) (proto.Message, error) {
			args := inp.(*DeleteRequest)
			return srv.Delete(ctx, args)
		}))

	s.Handle(prefix+"Search", tcprpc.NewHandler(
		func() proto.Message { return new(SearchRequest) },
		func(ctx context.Context, inp proto.Message) (proto.Message, error) {
			args := inp.(*SearchRequest)
			return srv.Search(ctx, args)
		}))

}

type TCPRPCMetaDBServiceClient struct {
	cl   *tcprpc.Client
	pref string
}

func NewTCPRPCMetaDBServiceClient(cl *tcprpc.Client, pref string) TCPRPCMetaDBServiceClient {
	return TCPRPCMetaDBServiceClient{
		cl:   cl,
		pref: pref,
	}
}

func (cl TCPRPCMetaDBServiceClient) Put(ctx context.Context, args *PutRequest) (*PutResponse, error) {
	var resp PutResponse
	err := cl.cl.Call(ctx, cl.pref+"Put", args, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cl TCPRPCMetaDBServiceClient) Get(ctx context.Context, args *GetRequest) (*GetResponse, error) {
	var resp GetResponse
	err := cl.cl.Call(ctx, cl.pref+"Get", args, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cl TCPRPCMetaDBServiceClient) GetMulti(ctx context.Context, args *GetMultiRequest) (*GetMultiResponse, error) {
	var resp GetMultiResponse
	err := cl.cl.Call(ctx, cl.pref+"GetMulti", args, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cl TCPRPCMetaDBServiceClient) Delete(ctx context.Context, args *DeleteRequest) (*DeleteResponse, error) {
	var resp DeleteResponse
	err := cl.cl.Call(ctx, cl.pref+"Delete", args, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cl TCPRPCMetaDBServiceClient) Search(ctx context.Context, args *SearchRequest) (*SearchResponse, error) {
	var resp SearchResponse
	err := cl.cl.Call(ctx, cl.pref+"Search", args, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type MetaDBServiceInterface interface {
	Put(context.Context, *PutRequest) (*PutResponse, error)

	Get(context.Context, *GetRequest) (*GetResponse, error)

	GetMulti(context.Context, *GetMultiRequest) (*GetMultiResponse, error)

	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)

	Search(context.Context, *SearchRequest) (*SearchResponse, error)
}
