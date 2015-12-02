package scheduler

import "time"

type EventInterface interface {
	Name() string
	Duration() time.Duration
	Execute()
	Run()
	Start()
	Stop()
}
