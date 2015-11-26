package scheduler

import (
	"time"
)

type IEvent interface {
	Name() string
	Duration() time.Duration
	Execute()
	Run()
	Start()
	Stop()
}
