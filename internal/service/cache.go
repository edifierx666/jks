package service

import "github.com/gogf/gf/v2/os/gcache"

type Cache struct {
  onlineUserCache *gcache.Cache
}

func NewCache() *Cache {
  return &Cache{onlineUserCache: gcache.New()}
}

func (s *Cache) OnlineUserCache() *gcache.Cache {
  return s.onlineUserCache
}
