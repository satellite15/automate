package server

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

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

	return &ingest.JobSchedulerStatus{}, status.Error(codes.Unimplemented, "TODO(ssd): unimplemented")
}
