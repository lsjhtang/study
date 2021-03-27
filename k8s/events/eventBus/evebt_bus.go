package eventBus

import (
	"sync"
)

type EventBus struct {
	subscribes map[string]EventDataChannel
	handlers   map[string]*EventHandler
	lock       sync.RWMutex
}

func NewEventBus() *EventBus {
	return &EventBus{subscribes: make(map[string]EventDataChannel), handlers: make(map[string]*EventHandler)}
}

//订阅
func (e *EventBus) Sub(topic string, fn interface{}) EventDataChannel {
	e.lock.Lock()
	defer e.lock.Unlock()

	e.subscribes[topic] = make(EventDataChannel)
	e.handlers[topic] = NewEventHandler(fn)
	return e.subscribes[topic]
}

//发布
func (e *EventBus) Pub(topic string, params ...interface{}) {
	e.lock.RLock()
	defer e.lock.RUnlock()
	if s, ok := e.subscribes[topic]; ok {
		handler := e.handlers[topic]
		go func() {
			s <- NewEventData(handler.Call(params...))
		}()
	}
}
