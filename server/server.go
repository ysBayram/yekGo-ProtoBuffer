package main

import (
	"log"
	"net"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	pb "yekGoProtoBuffer/server/generated"
	"yekGoProtoBuffer/server/handlers"
)

const (
	port = ":50051"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	db, err := gorm.Open("sqlite3", "./db/gorm.db")
	if err != nil {
		log.Fatalf("failed to connect DB : %v", err)
	}
	defer db.Close()

	db.AutoMigrate(&pb.Ogrenci{})

	ogrenciHandler := handlers.NewOgrenciHandler(db)

	pb.RegisterOgrenciHubServer(s, &ogrenciHandler)
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
