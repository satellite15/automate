package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	rrule "github.com/teambition/rrule-go"

	"github.com/chef/automate/api/interservice/ingest"
	"github.com/chef/automate/components/ingest-service/backend"
	"github.com/chef/automate/lib/workflow"

	log "github.com/sirupsen/logrus"
)

type JobSchedulerServer struct {
	client          backend.Client
	workflowManager *workflow.WorkflowManager
}

// NewJobSchedulerServer - create a new JobSchedulerServer
func NewJobSchedulerServer(client backend.Client, workflowManager *workflow.WorkflowManager) *JobSchedulerServer {
	return &JobSchedulerServer{
		client:          client,
		workflowManager: workflowManager,
	}
}

// StartJobScheduler - Start the Job Scheduler
func (server *JobSchedulerServer) StartJobScheduler(ctx context.Context,
	empty *ingest.StartJobSchedulerRequest) (*ingest.StartJobSchedulerResponse, error) {
	log.WithFields(log.Fields{"func": nameOfFunc()}).Debug("rpc call")

	return &ingest.StartJobSchedulerResponse{}, status.Error(codes.Unimplemented, "TODO(ssd): unimplemented")
}

// StopJobScheduler - Stop the Job Scheduler
func (server *JobSchedulerServer) StopJobScheduler(ctx context.Context,
	empty *ingest.StopJobSchedulerRequest) (*ingest.StopJobSchedulerResponse, error) {
	log.WithFields(log.Fields{"func": nameOfFunc()}).Debug("rpc call")

	return &ingest.StopJobSchedulerResponse{}, status.Error(codes.Unimplemented, "TODO(ssd): unimplemented")
}

// GetStatusJobScheduler - collect and return the status of all the jobs in the Job Scheduler
func (server *JobSchedulerServer) GetStatusJobScheduler(ctx context.Context,
	empty *ingest.JobSchedulerStatusRequest) (*ingest.JobSchedulerStatus, error) {
	log.WithFields(log.Fields{"func": nameOfFunc()}).Debug("rpc call")

	schedules, err := server.workflowManager.ListWorkflowSchedules(ctx)
	if err != nil {
		return &ingest.JobSchedulerStatus{}, status.Error(codes.Internal, err.Error())
	}

	jobs := make([]*ingest.Job, 0, len(schedules))
	for _, sched := range schedules {
		every, err := rruleToEvery(sched.Recurrence)
		if err != nil {
			return &ingest.JobSchedulerStatus{}, status.Error(codes.Internal, err.Error())
		}

		var threshold string
		err = json.Unmarshal(sched.Parameters, &threshold)
		if err != nil {
			return &ingest.JobSchedulerStatus{}, status.Error(codes.Internal, err.Error())
		}

		jobs = append(jobs, &ingest.Job{
			Running:   sched.Enabled,
			Name:      sched.WorkflowName,
			Every:     every,
			Threshold: threshold,
			NextRun:   getTimeString(sched.NextDueAt),
			// TODO(ssd) 2019-05-16: We should decide how we want to support these
			// LastRun:
			// LastElapsed:
			// StartedOn:

		})
	}

	return &ingest.JobSchedulerStatus{
		Jobs:    jobs,
		Running: true,
	}, nil
}

func getTimeString(dateTime time.Time) string {
	if dateTime.IsZero() {
		return ""
	}
	return dateTime.Format(time.RFC3339Nano)
}

func rruleToEvery(rruleStr string) (string, error) {
	r, err := rrule.StrToRRule(rruleStr)
	if err != nil {
		return "", err
	}

	switch r.OrigOptions.Freq {
	case rrule.HOURLY:
		return fmt.Sprintf("%dh", r.OrigOptions.Interval), nil
	case rrule.MINUTELY:
		return fmt.Sprintf("%dm", r.OrigOptions.Interval), nil
	case rrule.SECONDLY:
		return fmt.Sprintf("%ds", r.OrigOptions.Interval), nil
	default:
		return "", errors.New("Unsupported rrule to duration conversion")
	}
}
