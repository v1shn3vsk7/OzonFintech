package gRPC

import (
	"OzonTestTask/internal/app/data"
	"OzonTestTask/internal/app/data/sqldata"
	"OzonTestTask/internal/app/model"
	"OzonTestTask/internal/app/server"
	pb "OzonTestTask/internal/app/server/gRPC/proto"
	"context"
	"database/sql"
	"errors"
	"google.golang.org/grpc"
	"net"
	"os"
)

type GRPCServer struct {
	Data   data.IData
	server *grpc.Server
	pb.UnimplementedUrlServer
}

func NewgRPCServer(data data.IData, server *grpc.Server) *GRPCServer {
	return &GRPCServer{
		Data: data,
		server: server,
	}
}

func (s *GRPCServer) Serve(d data.IData, lis net.Listener) error {
	pb.RegisterUrlServer(s.server, &GRPCServer{Data: d})

	if err := s.server.Serve(lis); err != nil {
		return err
	}

	return nil
}

func Start(cfg *server.Config) error{
	listener, err := net.Listen("tcp", ":5536")
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()

	var d data.IData

	storeType := os.Getenv("STORE_TYPE")
	if storeType == "inmemory" {
		d = cfg.Data

	} else if storeType == "postgres" {
		db, err := newDb(os.Getenv("DB_URL"))
		if err != nil {
			return err
		}
		defer db.Close()

		d = sqldata.New(db)

		} else {
		return errors.New("no choice for storage type")
	}

	s := NewgRPCServer(d, grpcServer)
	if err := s.Serve(d, listener); err != nil {
		return err
	}

	return nil
}

func newDb(dbUrl string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func (s *GRPCServer) CreateShortUrl(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	link := &model.Link{
		OriginUrl: in.Url,
		ShortUrl: "",
	}

	if err := s.Data.Link().Create(link); err != nil {
		return nil, err
	}

	return &pb.Response{
		Url: link.ShortUrl,
	}, nil
}

func (s *GRPCServer) GetOriginUrl(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	link := &model.Link{
		OriginUrl: "",
		ShortUrl: in.Url,
	}

	if err := s.Data.Link().FindByShortURL(link); err != nil {
		return nil, err
	}

	return &pb.Response{
		Url: link.OriginUrl,
	}, nil
}
