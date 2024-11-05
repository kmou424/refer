package refer

import (
	"container/list"
	"sync"
)

type refContainer struct {
	orderedKeys *list.List
	objects     map[string]any
	keys        map[any]*list.Element
	lock        sync.RWMutex
}

func newContainer() *refContainer {
	return &refContainer{
		orderedKeys: list.New(),
		objects:     make(map[string]any),
		keys:        make(map[any]*list.Element),
		lock:        sync.RWMutex{},
	}
}

func (c *refContainer) storeRef(key string, value interface{}) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.objects[key] = value
	if _, ok := c.objects[key]; !ok {
		v := c.orderedKeys.PushBack(key)
		c.keys[value] = v
	}
}

func (c *refContainer) loadRef(key string) interface{} {
	c.lock.RLock()
	defer c.lock.RUnlock()
	if obj, ok := c.objects[key]; ok {
		return obj
	}
	return nil
}

func (c *refContainer) deleteRef(key string) {
	c.lock.Lock()
	defer c.lock.Unlock()
	if _, ok := c.objects[key]; ok {
		v := c.keys[c.objects[key]]
		c.orderedKeys.Remove(v)
		delete(c.objects, key)
		delete(c.keys, c.objects[key])
	}
}

func (c *refContainer) hasRef(key string) bool {
	c.lock.RLock()
	defer c.lock.RUnlock()
	_, ok := c.objects[key]
	return ok
}

func (c *refContainer) lookupRef(key string) (interface{}, bool) {
	c.lock.RLock()
	defer c.lock.RUnlock()
	if obj, ok := c.objects[key]; ok {
		return obj, true
	}
	return nil, false
}
