package scheduling

import (
	"log/slog"
	"time"

	"github.com/bamzi/jobrunner"
	"github.com/robfig/cron/v3"
)

type (
	Schedule struct {
		ID          interface{} `json:"id"`
		Name        string      `json:"name"`
		JobSchedule string      `json:"jobSchedule"`
		Executions  int         `json:"executions"`
		Job         func()      `json:"-"`
	}

	Scheduler struct {
		Schedules []*Schedule
	}
)

// Run is the function that will be called by the scheduler
func (s *Schedule) Run() {
	s.Job()
}

// NewScheduler creates a new scheduler
func NewScheduler() *Scheduler {
	return &Scheduler{}
}

// Start starts the scheduler
func (s *Scheduler) Start() {

	jobrunner.Start()

	for _, schedule := range s.Schedules {
		err := jobrunner.Schedule(schedule.JobSchedule, schedule)

		e := jobrunner.Entries()
		j := e[len(e)-1]

		schedule.ID = j.ID

		slog.Info("Scheduling job", "schedule", schedule)
		if err != nil {
			slog.Error("Error while scheduling job", "error", err.Error())
		}
	}
}

// AddSchedule adds a new schedule to the scheduler
func (s *Scheduler) AddSchedule(message string, jobSchedule string, job func()) {
	s.Schedules = append(s.Schedules, &Schedule{Name: message, JobSchedule: jobSchedule, Job: job})
}

// GetSchedules returns all schedules
func (s *Scheduler) GetSchedules() []*Schedule {
	return s.Schedules
}

// ClearSchedules clears all schedules
func (s *Scheduler) ClearSchedules() {
	jobs := jobrunner.StatusJson()
	slog.Info("Clearing schedules", "jobs", jobs)

	// All running schedules must be stopped first
	s.Schedules = []*Schedule{}
}

func (s *Scheduler) StopJobById(id int) {

	for _, e := range jobrunner.Entries() {
		if e.ID == cron.EntryID(id) {
			jobrunner.Remove(e.ID)

			// Remove the Job from the Scheduler
			for i, schedule := range s.Schedules {
				if schedule.ID == e.ID {
					s.Schedules = append(s.Schedules[:i], s.Schedules[i+1:]...)
					break
				}
			}
		}
	}
}

func Scehdule() *Scheduler {
	slog.Info("Scheduling jobs")

	scheduler := NewScheduler()

	scheduler.AddSchedule("Poke", "@every 5s", func() {
		slog.Info("Scheduled 5s Poked...", "time", time.Now())
	})
	scheduler.AddSchedule("Poke2", "@every 7s", func() {
		slog.Info("Scheduled 7s Poked...", "time", time.Now())
	})

	slog.Info("Schedules", "schedules", scheduler.GetSchedules())

	scheduler.Start()
	return scheduler
}
