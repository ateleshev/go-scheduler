package scheduler

import (
	"log"
	"sync"
	"time"
)

func NewEvent(name string, duration time.Duration, action Action, periodic bool) *Event { // {{{
	return &Event{
		name:     name,
		duration: duration,
		action:   action,
		periodic: periodic,
	}
} // }}}

func NewPeriodicEvent(name string, duration time.Duration, action Action) *Event { // {{{
	return NewEvent(name, duration, action, true)
} // }}}

func NewSingleEvent(name string, duration time.Duration, action Action) *Event { // {{{
	return NewEvent(name, duration, action, false)
} // }}}

type Event struct {
	EventInterface
	sync.RWMutex

	name     string
	duration time.Duration // time.Hour | 5 * time.Minute | 30 * time.Second | 200 * time.Millisecond | time.Microsecond | time.Nanosecond
	action   Action
	periodic bool
	executed bool
	stop     chan bool
}

func (this *Event) Name() string { // {{{
	return this.name
} // }}}

func (this *Event) Duration() time.Duration { // {{{
	return this.duration
} // }}}

func (this *Event) Execute() { // {{{
	if err := this.action(); err != nil {
		log.Println("[scheduler] Event error:", err)
	}
} // }}}

func (this *Event) runSingle() { // {{{
	select {
	case <-time.After(this.duration):
		this.action()
	case <-this.stop:
		return
	}
} // }}}

func (this *Event) runPeriodic() { // {{{
	for {
		select {
		case <-time.After(this.duration):
			this.Execute()
		case <-this.stop:
			return
		}
	}
} // }}}

func (this *Event) Run() { // {{{
	if !this.executed {
		this.Lock()
		defer this.Unlock()

		this.executed = true

		if this.periodic {
			this.runPeriodic()
		} else {
			this.runSingle()
		}
	}
} // }}}

func (this *Event) Start() { // {{{
	if !this.executed {
		go this.Run()
	}
} // }}}

func (this *Event) Stop() { // {{{
	this.Lock()
	defer this.Unlock()

	this.stop <- true
	this.executed = false
} // }}}
