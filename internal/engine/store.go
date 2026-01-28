package engine

import "sync"

type StringStore struct {
	shards []*Shard
}

var ShardCount = 256

type Shard struct {
	mu    sync.Mutex
	items map[string]string
}

func NewStringStore() *StringStore {
	s := &StringStore{
		shards: make([]*Shard, ShardCount),
	}

	for i := 0; i < ShardCount; i++ {
		s.shards[i] = &Shard{
			items: make(map[string]string),
		}
	}
	return s
}


func (s *StringStore) getShard(key string) *Shard {
	h := HashFNV32(key)

	index := h % uint32(ShardCount)

	return s.shards[index]
}

func (s *StringStore) Set(key, value string)  {
	
	shard := s.getShard(key)

	shard.mu.Lock()
	defer shard.mu.Unlock()

	shard.items[key] = value
}

func (s *StringStore) Get(key string) (string, bool)  {
	

	shard := s.getShard(key)

	shard.mu.Lock()
	defer shard.mu.Unlock()

	val, ok := shard.items[key]
	return val, ok
}