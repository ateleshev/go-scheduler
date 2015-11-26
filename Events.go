package scheduler

import (
	"sync"
)

type Events struct {
	sync.RWMutex

	data map[string]*Event
}

func NewEvents() *Events { // {{{
	return &Events{
		data: make(map[string]*Event, 0),
	}
} // }}}

func (this *Events) Len() int { // {{{
	return len(this.data)
} // }}}

func (this *Events) Data() map[string]*Event { // {{{
	return this.data
} // }}}

func (this *Events) Set(name string, event *Event) { // {{{
	this.Lock()
	defer this.Unlock()

	this.data[name] = event
} // }}}

func (this *Events) Has(name string) bool { // {{{
	_, ok := this.data[name]
	return ok
} // }}}

func (this *Events) Get(name string) *Event { // {{{
	return this.data[name]
} // }}}
