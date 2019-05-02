package server

// RPC functions for the job: 'missing_node'

import (
	"context"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/chef/automate/api/interservice/ingest"
	"github.com/chef/automate/components/ingest-service/config"
	"github.com/chef/automate/lib/workflow"
)

const (
	markNodesMissingWorkflowName = "missing_nodes"
	markNodesMissingScheduleName = "periodic_missing_nodes"
)

// MarkNodesMissing - run the mark nodes missing task now
func (server *JobSchedulerServer) MarkNodesMissing(ctx context.Context,
	empty *ingest.MarkNodesMissingRequest) (*ingest.MarkNodesMissingResponse, error) {

	var threshold string
	err := server.workflowManager.GetScheduledWorkflowParameters(ctx,
		markNodesMissingScheduleName, markNodesMissingWorkflowName,
		&threshold)
	if err != nil {
		return &ingest.MarkNodesMissingResponse{}, status.Error(codes.Internal, err.Error())
	}

	nodesMissing, err := server.client.MarkNodesMissing(ctx, threshold)
	if err != nil {
		return &ingest.MarkNodesMissingResponse{}, status.Error(codes.Internal, err.Error())
	}

	log.WithFields(log.Fields{
		"nodes_updated": nodesMissing,
		"status":        "missing",
	}).Info("Marked nodes missing")
	return &ingest.MarkNodesMissingResponse{}, nil
}

// StartNodesMissingScheduler - start the scheduled task of deleting nodes
func (server *JobSchedulerServer) StartNodesMissingScheduler(ctx context.Context,
	empty *ingest.StartNodesMissingSchedulerRequest) (*ingest.StartNodesMissingSchedulerResponse, error) {
	log.Info("StartNodesMissingScheduler")

	err := server.workflowManager.UpdateWorkflowScheduleByName(
		ctx,
		markNodesMissingScheduleName, markNodesMissingWorkflowName,
		workflow.UpdateEnabled(true))
	if err != nil {
		return &ingest.StartNodesMissingSchedulerResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &ingest.StartNodesMissingSchedulerResponse{}, nil
}

// StopNodesMissingScheduler - stop the scheduled delete node task from running
func (server *JobSchedulerServer) StopNodesMissingScheduler(ctx context.Context,
	empty *ingest.StopNodesMissingSchedulerRequest) (*ingest.StopNodesMissingSchedulerResponse, error) {
	log.Info("StopNodesMissingScheduler")

	err := server.workflowManager.UpdateWorkflowScheduleByName(
		ctx,
		markNodesMissingScheduleName, markNodesMissingWorkflowName,
		workflow.UpdateEnabled(false))
	if err != nil {
		return &ingest.StopNodesMissingSchedulerResponse{}, status.Error(codes.Internal, err.Error())
	}
	return &ingest.StopNodesMissingSchedulerResponse{}, nil
}

// ConfigureNodesMissingScheduler rpc call to configure the NodesMissing Job
func (server *JobSchedulerServer) ConfigureNodesMissingScheduler(ctx context.Context,
	settings *ingest.JobSettings) (*ingest.ConfigureNodesMissingSchedulerResponse, error) {
	log.WithFields(log.Fields{
		"settings": settings.String(),
	}).Info("Incoming job")

	oldRRule, err := server.workflowManager.GetScheduledWorkflowRecurrence(
		ctx, markNodesMissingScheduleName, markNodesMissingWorkflowName)
	if err != nil {
		return &ingest.ConfigureNodesMissingSchedulerResponse{}, status.Error(codes.Internal, err.Error())
	}

	updateOpts, err := config.JobSettingsToUpdateOpts(settings, oldRRule)
	if err != nil {
		return &ingest.ConfigureNodesMissingSchedulerResponse{},
			status.Error(codes.InvalidArgument, err.Error())
	}

	// apply job settings to the job config, then update the job if needed
	err = server.workflowManager.UpdateWorkflowScheduleByName(
		ctx, markNodesMissingScheduleName, markNodesMissingWorkflowName, updateOpts...)
	if err != nil {
		return &ingest.ConfigureNodesMissingSchedulerResponse{},
			status.Error(codes.Internal, err.Error())

	}
	return &ingest.ConfigureNodesMissingSchedulerResponse{}, nil
}
