package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	pb "yekGoProtoBuffer/client/generated"
)

const (
	address = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewOgrenciHubClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	rInsertOgrenci, err := c.InsertOgrenci(ctx, &pb.OgrenciRequest{Ogrenci: &pb.Ogrenci{Name: "Foo2", SurName: "Bar2", Number: 2}})
	if err != nil {
		log.Fatalf("could not InsertOgrenci: %v", err)
	}
	log.Printf("InsertOgrenci output : %s", rInsertOgrenci.Ogrenci)

	rGetOgrenciByID, err := c.GetOgrenciByID(ctx, &pb.OgrenciID{Id: 2})
	if err != nil {
		log.Fatalf("could not GetOgrenciByID: %v", err)
	}
	log.Printf("GetOgrenciByID output : %s", rGetOgrenciByID.Ogrenci)

	rUpdateOgrenci, err := c.UpdateOgrenci(ctx, &pb.OgrenciRequest{Ogrenci: &pb.Ogrenci{Id: 2, Name: "FooUpdated", SurName: "BarUpdated", Number: 2}})
	if err != nil {
		log.Fatalf("could not UpdateOgrenci: %v", err)
	}
	log.Printf("UpdateOgrenci output : %s", rUpdateOgrenci.Ogrenci)

	rDeleteOgrenciByID, err := c.DeleteOgrenciByID(ctx, &pb.OgrenciID{Id: 2})
	if err != nil {
		log.Fatalf("could not DeleteOgrenciByID: %v", err)
	}
	log.Printf("DeleteOgrenciByID output : %s", rDeleteOgrenciByID.Ogrenci)

	rGetOgrenciAll, err := c.GetOgrenciAll(ctx, &pb.OgrenciID{Id: 0})
	if err != nil {
		log.Fatalf("could not GetOgrenciAll: %v", err)
	}
	log.Printf("GetOgrenciAll output : %s", rGetOgrenciAll.Ogrenci)
}
