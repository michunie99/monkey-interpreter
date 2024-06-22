package object

// import (
// 	"sync"
// 	"time"
// )
//
// type HashCache struct {
// 	store sync.Map
// }
//
// func NewHashCache() *HashCache {
// 	return &HashCache{}
// }
//
// func (hc *HashCache) Set(key string, value uint64, ttl time.Duration) {
// 	hc.store.Store(key, value)
// 	time.AfterFunc(ttl, func() {
// 		hc.store.Delete(key)
// 	})
// }
//
// func (hc *HashCache) Get(key string) (uint64, bool) {
// 	val, ok := hc.store.Load(key)
// 	if ok {
// 		val, ok := val.(uint64)
// 		return val, ok
// 	}
// 	return val, ok
// }
