# bloomcache

## Development

### Install
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

### How to Build Protobuf
If you modified .proto file, you will need to rebuild it with protobuf compiler

#### Install Protobuf 3
https://github.com/golang/protobuf

#### Compile Protobuf file
```bash
protoc --go_out=plugins=grpc:. protobuf/bloomcache.proto
```

Reference:
