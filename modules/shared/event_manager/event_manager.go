package eventmanager

import "sync"

type Event interface {
	Name() string
	Data() interface{}
}

type Listener func(Event)

type EventManager struct {
	listeners map[string][]Listener
	mu        sync.RWMutex
}

func NewEventManager() *EventManager {
	return &EventManager{
		listeners: make(map[string][]Listener),
	}
}

// GlobalInstance is the shared, global event manager
var Dispatcher = &EventManager{
	listeners: make(map[string][]Listener),
}

func (em *EventManager) On(eventType string, listener Listener) {
	em.mu.Lock()
	defer em.mu.Unlock()
	em.listeners[eventType] = append(em.listeners[eventType], listener)
}

func (em *EventManager) Emit(event Event) {
	em.mu.RLock()
	defer em.mu.RUnlock()

	// event.Type() calls the method on your specific struct
	for _, listener := range em.listeners[event.Name()] {
		go listener(event)
	}
}
