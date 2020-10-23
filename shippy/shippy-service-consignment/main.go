package main

import (
	"sync"

	pb "example.com/user/shippy-service-consignment/proto/consignment"
)

const (
	port = ":50051"
)

type repository interface {
	create(*pb.Consignment) (*pb.Consignment, error)
}

// Repository - dummy repository that will be replaced later
type Repository struct {
	my           sync.RWMutex
	consignments []*pb.Consignment
}

func main() {

}
