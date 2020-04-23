import (
	"sync"
)

type data struct {
	mutex  sync.Mutex
	values map[string]int
}

func (d *data) Add(key string, value int) {
	d.mutex.Lock()
	d.values[key] = value
	d.mutex.Unlock()
}
