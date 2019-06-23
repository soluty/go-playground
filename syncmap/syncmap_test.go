package syncmap

import (
"sync/atomic"
"testing"
"github.com/gogf/gf/g/container/gmap"
"sync"
)

type MySyncMap struct {
	lock sync.RWMutex
	data map[interface{}]interface{}
}

func (m *MySyncMap) Set(k,v interface{}) {
	m.lock.Lock()
	m.data[k]=v
	m.lock.Unlock()
}

func (m *MySyncMap) Get(k interface{}) interface{} {
	m.lock.RLock()
	val, _ := m.data[k]
	m.lock.RUnlock()
	return val
}

var m1 = gmap.NewAnyAnyMap()
var m2 = sync.Map{}
var m3 = &MySyncMap{
	data: map[interface{}]interface{}{},
}

func BenchmarkGmapSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		go m1.Set(i, i)
	}
}

func BenchmarkMyMapSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		go m3.Set(i, i)
	}
}


func BenchmarkSyncmapSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		go m2.Store(i, i)
	}
}



func BenchmarkGmapGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		go m1.Get(i)
	}
}

func BenchmarkMyMapGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		go m3.Get(i)
	}
}

func BenchmarkSyncmapGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		go m2.Load(i)
	}
}

func BenchmarkGmapRemove(b *testing.B) {
	for i := 0; i < b.N; i++ {
		go m1.Remove(i)
	}
}

func BenchmarkSyncmapRmove(b *testing.B) {
	for i := 0; i < b.N; i++ {
		go m2.Delete(i)
	}
}


var testInt int64

func BenchmarkAtomicAdd(b *testing.B) {
	testInt = 0
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		go atomic.AddInt64(&testInt,1)
	}
}

func BenchmarkMutexAdd(b *testing.B) {
	testInt = 0
	var l sync.Mutex
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		go func() {
			l.Lock()
			testInt = testInt+1
			l.Unlock()
		}()
	}
}
