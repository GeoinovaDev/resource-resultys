package resource

import (
	"sync"

	"github.com/GeoinovaDev/lower-resultys/collection/queue"
)

// Resource struct
type Resource struct {
	Limit int

	used      int
	mutex     *sync.Mutex
	onRelease func()
	fila      *queue.Queue
}

// New cria o ralloc
func New(limit int) *Resource {
	resource := &Resource{Limit: limit}

	resource.mutex = &sync.Mutex{}
	resource.fila = queue.New()

	return resource
}

// Alloc aloca resource
func (resource *Resource) Alloc(callback func()) *Resource {
	resource.mutex.Lock()

	if resource.used == resource.Limit {
		resource.fila.Push(item{cb: callback})
		resource.mutex.Unlock()
		return resource
	}

	resource.used++
	resource.mutex.Unlock()

	callback()

	resource.Release()

	return resource
}

// Release libera o resource
func (resource *Resource) Release() *Resource {
	resource.mutex.Lock()
	resource.used--
	resource.mutex.Unlock()

	if resource.fila.IsEmpty() {
		return resource
	}

	item := resource.fila.Pop().(item)
	resource.Alloc(item.cb)

	return resource
}
