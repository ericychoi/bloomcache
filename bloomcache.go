package bloomcache

import (
	"log"
	"os"

	"github.com/ericychoi/bloomcache/protobuf"
	"github.com/willf/bloom"
	"golang.org/x/net/context"
)

var Logger *log.Logger = log.New(os.Stdout, `[bloomcache] `, log.LstdFlags)

// Server contains a bloomfilter and handles adding and checking
type Server struct {
	bf *bloom.BloomFilter
}

//TODO: need AddBatch and CheckBatch (InList call takes a list)

// Add adds a key to the bloomfilter
func (s *Server) Add(ctx context.Context, in *protobuf.Request) (*protobuf.Response, error) {
	Logger.Printf("received Add(): %s", in.Key)
	s.bf.AddString(in.Key)
	return &protobuf.Response{Error: ""}, nil
}

// Check checks whether a given key is in the bloom filter or not
func (s *Server) Check(ctx context.Context, in *protobuf.Request) (*protobuf.CheckResponse, error) {
	Logger.Printf("received check(): %s", in.Key)
	b := s.bf.TestString(in.Key)
	return &protobuf.CheckResponse{Exists: b}, nil
}

// New returns a new Server instance given m and k for the bloom filter.
func New(m int, k int) *Server {
	return &Server{bf: bloom.New(uint(m), uint(k))}
}
