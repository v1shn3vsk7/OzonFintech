package gRPC

import (
	"OzonTestTask/internal/app/data"
	"OzonTestTask/internal/app/data/inmemory"
	"OzonTestTask/internal/app/data/sqldata"
	"OzonTestTask/internal/app/model"
	pb "OzonTestTask/internal/app/server/gRPC/proto"
	"context"
	"database/sql"
	"errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"log"
	"net"
	"os"
)

type GRPCServer struct {
	data   data.Data
	server *grpc.Server
	pb.UnimplementedUrlServer
}

func NewgRPCServer(data data.Data, server *grpc.Server) *GRPCServer {
	return &GRPCServer{
		data: data,
		server: server,
	}
}

func Start() error{
	listener, err := net.Listen("tcp", ":5536")
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	storeType := os.Getenv("STORE_TYPE")
	if storeType == "inmemory" {
		d := &inmemory.Data{}

		s := NewgRPCServer(d, grpcServer)

		pb.RegisterUrlServer(s.server, &GRPCServer{data: d})

		if err := s.server.Serve(listener); err != nil {
			log.Fatal(err)
		}

	} else if storeType == "postgres" {
		db, err := newDb("postgres://user:password@db:5432/ozontesttask?sslmode=disable") /////////TODO FIX
		if err != nil {
			return err
		}

		defer db.Close()

		d := sqldata.New(db)

		s := NewgRPCServer(d, grpcServer)

		pb.RegisterUrlServer(s.server, &GRPCServer{data: d})

		if err := s.server.Serve(listener); err != nil {
			log.Fatal(err)
		}

		//s.server.Serve(listener)
	} else {
		return errors.New("no choice for storage type")
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

	//panic(s.data)

	if err := s.data.Link().Create(link); err != nil {
		//s.error(w, r, http.StatusUnprocessableEntity, err)
		log.Fatal(err)
	}

	return &pb.Response{
		Message: link.ShortUrl,
		Error:   "",
	}, nil
}

func (s *GRPCServer) GetOriginUrl(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	link := &model.Link{
		OriginUrl: "",
		ShortUrl: in.Url,
	}

	if err := s.data.Link().FindByShortURL(link); err != nil {
		//s.error(w, r, http.StatusNoContent, err)
		log.Fatal(err)
	}

	return &pb.Response{
		Message: link.OriginUrl,
		Error:   "",
	}, nil
}
