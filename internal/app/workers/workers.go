package workers

import (
	cronv3 "github.com/robfig/cron/v3"
)

type Workers struct {
	Cron *cronv3.Cron
}

func NewWorkers() *Workers {
	cron := cronv3.New()
	cron.Start()
	return &Workers{Cron: cron}
}
