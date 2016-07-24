# bloomcache

Bloomcache is a simple cache server backed by a bloom filter.

It relies on [GRPC](http://www.grpc.io/) to handle requests.  A simple client is provided, but you may use any GRPC client that implements protobuf specified in [github.com/ericychoi/bloomcache/protobuf/bloomcache.proto](https://github.com/ericychoi/bloomcache/blob/master/protobuf/bloomcache.proto).

## Install
```bash
go get github.com/ericychoi/bloomcache
source $GOPATH/src/github.com/ericychoi/bloomcache/development.env && bcserver
```
A test client is provided in `bin/client`

```bash
% cd $GOPATH/src/github.com/ericychoi/bloomcache
% go run bin/client.go -key test -add
% go run bin/client.go -key test -check
2016/07/23 11:49:25 test exists
```




## Development

### Test
```bash
% go test -v -race
=== RUN   TestServer
2016/07/23 23:23:03 received Add(): testKey
2016/07/23 23:23:03 received check(): testKey
2016/07/23 23:23:03 received check(): differentKey
--- PASS: TestServer (0.00s)
PASS
ok  	github.com/ericychoi/bloomcache	8.101s
```

### Benchmark
```bash
% go test -bench=.
bloomcache2016/07/24 00:59:21 received Add(): testKey
bloomcache2016/07/24 00:59:21 received check(): testKey
bloomcache2016/07/24 00:59:21 received check(): differentKey
PASS
BenchmarkAddByK3-8   	 2000000	       983 ns/op
BenchmarkAddByK30-8  	 1000000	      1318 ns/op
BenchmarkAddByK300-8 	  300000	      5416 ns/op
BenchmarkAddByK3000-8	   30000	     45674 ns/op
ok  	github.com/ericychoi/bloomcache	23.114s
```

### How to Build Protobuf
If you modified .proto file, you will need to rebuild it with protobuf compiler

#### Install Protobuf 3
https://github.com/golang/protobuf

#### Compile Protobuf file
```bash
protoc --go_out=plugins=grpc:. protobuf/bloomcache.proto
```

Reference:
