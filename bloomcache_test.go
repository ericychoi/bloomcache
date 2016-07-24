package bloomcache

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/ericychoi/bloomcache/protobuf"
	"github.com/nu7hatch/gouuid"
	"github.com/stretchr/testify/assert"
	"github.com/willf/bloom"
	"golang.org/x/net/context"
)

func TestServer(t *testing.T) {
	m := 1000
	k := 5
	key1, key2 := "testKey", "differentKey"
	s := &Server{bf: bloom.New(uint(m), uint(k))}

	resp, _ := s.Add(context.Background(), &protobuf.Request{Key: key1})
	if resp.Error != "" {
		t.Fatalf("Add returned error: %s\n", resp.Error)
	}

	cResp, _ := s.Check(context.Background(), &protobuf.Request{Key: key1})
	if cResp.Error != "" {
		t.Fatalf("Check returned error: %s\n", resp.Error)
	}
	assert.Equal(t, true, cResp.Exists, fmt.Sprintf("%s should exist in the BF", key1))

	cResp, _ = s.Check(context.Background(), &protobuf.Request{Key: key2})
	if cResp.Error != "" {
		t.Fatalf("Check returned error: %s\n", resp.Error)
	}
	assert.Equal(t, false, cResp.Exists, fmt.Sprintf("%s should not exist in the BF", key2))
}

func benchmarkAddByK(k int, b *testing.B) {
	m := 1000000
	s := &Server{bf: bloom.New(uint(m), uint(k))}
	Logger.SetOutput(ioutil.Discard)

	var keys []string
	for i := 0; i < b.N; i++ {
		u4, err := uuid.NewV4()
		if err != nil {
			b.Fatal("couldn't produce uuid v4")
		}
		keys = append(keys, u4.String())
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Add(context.Background(), &protobuf.Request{Key: keys[i]})
	}
}

func BenchmarkAddByK3(b *testing.B)    { benchmarkAddByK(3, b) }
func BenchmarkAddByK30(b *testing.B)   { benchmarkAddByK(30, b) }
func BenchmarkAddByK300(b *testing.B)  { benchmarkAddByK(300, b) }
func BenchmarkAddByK3000(b *testing.B) { benchmarkAddByK(3000, b) }

func benchmarkCheckByK(k int, b *testing.B) {
	m := 1000000
	s := &Server{bf: bloom.New(uint(m), uint(k))}
	Logger.SetOutput(ioutil.Discard)

	var keys []string
	for i := 0; i < b.N; i++ {
		u4, err := uuid.NewV4()
		if err != nil {
			b.Fatal("couldn't produce uuid v4")
		}
		keys = append(keys, u4.String())
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Check(context.Background(), &protobuf.Request{Key: keys[i]})
	}
}

func BenchmarkCheckByK3(b *testing.B)    { benchmarkCheckByK(3, b) }
func BenchmarkCheckByK30(b *testing.B)   { benchmarkCheckByK(30, b) }
func BenchmarkCheckByK300(b *testing.B)  { benchmarkCheckByK(300, b) }
func BenchmarkCheckByK3000(b *testing.B) { benchmarkCheckByK(3000, b) }
