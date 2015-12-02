package scheduler

func NewScheduler(name string) *Scheduler { // {{{
	return &Scheduler{
		name:   name,
		events: NewEvents(),
	}
} // }}}

type Scheduler struct {
	SchedulerInterface

	name   string
	events *Events
}

func (this *Scheduler) Name() string { // {{{
	return this.name
} // }}}

func (this *Scheduler) Set(event *Event) { // {{{
	this.events.Set(event.Name(), event)
} // }}}

func (this *Scheduler) Has(eventName string) bool { // {{{
	return this.events.Has(eventName)
} // }}}

func (this *Scheduler) Get(eventName string) *Event { // {{{
	return this.events.Get(eventName)
} // }}}

func (this *Scheduler) Start() { // {{{
	for _, event := range this.events.Data() {
		event.Start()
	}
} // }}}

func (this *Scheduler) Stop() { // {{{
	for _, event := range this.events.Data() {
		event.Stop()
	}
} // }}}

func (this *Scheduler) StartEvent(name string) { // {{{
	if this.Has(name) {
		this.Get(name).Start()
	}
} // }}}

func (this *Scheduler) StopEvent(name string) { // {{{
	if this.Has(name) {
		this.Get(name).Stop()
	}
} // }}}
