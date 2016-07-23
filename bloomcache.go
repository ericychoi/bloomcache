package bloomcache

import (
	"log"

	"github.com/ericychoi/bloomcache/protobuf"
	"github.com/willf/bloom"
	"golang.org/x/net/context"
)

// Bloomcache Server
type Server struct {
	BF *bloom.BloomFilter
}

//TODO: need AddBatch and CheckBatch (InList call takes a list)
func (s *Server) Add(ctx context.Context, in *protobuf.Request) (*protobuf.Response, error) {
	log.Printf("received Add(): %s", in.Key)
	s.BF.AddString(in.Key)
	return &protobuf.Response{Error: ""}, nil
}

func (s *Server) Check(ctx context.Context, in *protobuf.Request) (*protobuf.CheckResponse, error) {
	log.Printf("received check(): %s", in.Key)
	b := s.BF.TestString(in.Key)
	return &protobuf.CheckResponse{Exists: b}, nil
}
