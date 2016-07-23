package main

import (
	"flag"
	"log"

	"github.com/ericychoi/bloomcache/protobuf"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	isAdd   bool
	isCheck bool
	host    string
	key     string
)

func main() {
	usage := `go run client.go -check -key $KEY -host $HOST`
	flag.BoolVar(&isAdd, "add", false, "perform add")
	flag.BoolVar(&isCheck, "check", false, "perform check")
	flag.StringVar(&key, "key", "", "key to add / check")
	flag.StringVar(&host, "host", "localhost:58080", "host:port")
	flag.Parse()

	if key == "" {
		log.Fatalf("key is required. usage: %s\n", usage)
	}
	if !isAdd && !isCheck {
		log.Fatalf("add or check is required. usage: %s\n", usage)
	}

	// Set up a connection to the server.
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()
	c := protobuf.NewBloomcacheClient(conn)

	// Contact the server and print out its response.
	if isAdd {
		resp, err := c.Add(context.Background(), &protobuf.Request{Key: key})
		if err != nil {
			log.Fatalf("could not add: key: %s, err: %s", key, err)
		}
		if resp.Error != "" {
			log.Fatalf("response contains error: key: %s, err: %s", key, err)
		}
	} else {
		resp, err := c.Check(context.Background(), &protobuf.Request{Key: key})
		if err != nil {
			log.Fatalf("could not check: key: %s, err: %s", key, err)
		}

		if resp.Exists {
			log.Printf("%s exists\n", key)
		} else {
			log.Printf("%s doesn't exist in cache\n", key)
		}
	}
}
