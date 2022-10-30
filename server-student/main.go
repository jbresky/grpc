package main

import (
	"log"
	"net"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"github.com/jbresky/grpc/database"
	"github.com/jbresky/grpc/server"
	"github.com/jbresky/grpc/studentpb"
)

func main(){
	list, err := net.Listen("tcp", ":5060")

	if err != nil {
		log.Fatal(nil)
	}

	repo, err := database.NewPostgresRepository("postgres://postgres:postgres@localhost:54321/postgres?sslmode=disable")

	server := server.NewStudentServer(repo)

	if err != nil {
		log.Fatal(nil)
	}

	s := grpc.NewServer()
	studentpb.RegisterStudentServiceServer(s, server)

	reflection.Register(s)

	if err := s.Serve(list); err != nil {
		log.Fatal(err)
	}
}