package specs

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Scheduler int

const (
	SchedulerDFS Scheduler = iota
	SchedulerRoundRobin
)

var AllSchedulers = Schedulers{SchedulerDFS, SchedulerRoundRobin}
var AllSchedulerNames = [...]string{
	SchedulerDFS:        "dfs",
	SchedulerRoundRobin: "round-robin",
}

type Schedulers []Scheduler

func (s Schedulers) String() string {
	var buffer bytes.Buffer
	for i, scheduler := range s {
		if i > 0 {
			buffer.WriteString(", ")
		}
		buffer.WriteString(scheduler.String())
	}
	return buffer.String()
}

func (s Scheduler) String() string {
	return AllSchedulerNames[s]
}
func (s Scheduler) MarshalJSON() ([]byte, error) {
	return []byte(`"` + s.String() + `"`), nil
}

func (s *Scheduler) UnmarshalJSON(data []byte) (err error) {
	var scheduler string
	if err := json.Unmarshal(data, &scheduler); err != nil {
		return err
	}
	if *s, err = SchedulerFromString(scheduler); err != nil {
		return err
	}
	return nil
}

func SchedulerFromString(s string) (Scheduler, error) {
	for i, scheduler := range AllSchedulerNames {
		if s == scheduler {
			return Scheduler(i), nil
		}
	}
	return SchedulerDFS, fmt.Errorf("unknown scheduler %s", s)
}
