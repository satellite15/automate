package server

// RPC functions for the job: 'missing_nodes_for_deletion'

import (
	"context"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/chef/automate/api/interservice/ingest"
	"github.com/chef/automate/components/ingest-service/config"
	"github.com/chef/automate/lib/workflow"
)

// MarkMissingNodesForDeletion - run the mark missing nodes for deletion task now
func (server *JobSchedulerServer) MarkMissingNodesForDeletion(ctx context.Context,
	empty *ingest.MarkMissingNodesForDeletionRequest) (*ingest.MarkMissingNodesForDeletionResponse, error) {

	var threshold string
	err := server.workflowManager.GetScheduledWorkflowParameters(ctx, "periodic_missing_nodes_for_deletion", "missing_nodes_for_deletion", &threshold)
	if err != nil {
		return &ingest.MarkMissingNodesForDeletionResponse{}, status.Error(codes.Internal, err.Error())
	}

	nodes4Deletion, err := server.client.MarkMissingNodesForDeletion(ctx, threshold)
	if err != nil {
		return &ingest.MarkMissingNodesForDeletionResponse{}, status.Error(codes.Internal, err.Error())
	}

	log.WithFields(log.Fields{
		"nodes_updated": nodes4Deletion,
		"exists":        false,
	}).Info("Node(s) marked for deletion")
	return &ingest.MarkMissingNodesForDeletionResponse{}, nil
}

// StartMissingNodesForDeletionScheduler - start the scheduled task of marking nodes for deletion
func (server *JobSchedulerServer) StartMissingNodesForDeletionScheduler(ctx context.Context,
	empty *ingest.StartMissingNodesForDeletionSchedulerRequest) (*ingest.StartMissingNodesForDeletionSchedulerResponse, error) {
	log.Info("StartMissingNodesForDeletionScheduler")

	err := server.workflowManager.UpdateWorkflowScheduleByName(
		ctx, "periodic_missing_nodes_for_deletion", "missing_nodes_for_deletion", workflow.UpdateEnabled(true))

	if err != nil {
		return &ingest.StartMissingNodesForDeletionSchedulerResponse{}, status.Error(codes.Internal, err.Error())
	}
	return &ingest.StartMissingNodesForDeletionSchedulerResponse{}, nil
}

// StopMissingNodesForDeletionScheduler - stop the scheduled marking nodes for deletion task from running
func (server *JobSchedulerServer) StopMissingNodesForDeletionScheduler(ctx context.Context,
	empty *ingest.StopMissingNodesForDeletionSchedulerRequest) (*ingest.StopMissingNodesForDeletionSchedulerResponse, error) {
	log.Info("StopMissingNodesForDeletionScheduler")

	err := server.workflowManager.UpdateWorkflowScheduleByName(
		ctx, "periodic_missing_nodes_for_deletion", "missing_nodes_for_deletion", workflow.UpdateEnabled(false))
	if err != nil {
		return &ingest.StopMissingNodesForDeletionSchedulerResponse{}, status.Error(codes.Internal, err.Error())
	}
	return &ingest.StopMissingNodesForDeletionSchedulerResponse{}, nil
}

// ConfigureMissingNodesForDeletionScheduler rpc call to configure the MissingNodesForDeletion Job
func (server *JobSchedulerServer) ConfigureMissingNodesForDeletionScheduler(ctx context.Context,
	settings *ingest.JobSettings) (*ingest.ConfigureMissingNodesForDeletionSchedulerResponse, error) {
	log.WithFields(log.Fields{
		"settings": settings.String(),
	}).Info("Incoming job")

	oldRRule, err := server.workflowManager.GetScheduledWorkflowRecurrence(
		ctx, "periodic_missing_nodes_for_deletion", "missing_nodes_for_deletion")
	if err != nil {
		return &ingest.ConfigureMissingNodesForDeletionSchedulerResponse{}, status.Error(codes.Internal, err.Error())
	}

	updateOpts, err := config.JobSettingsToUpdateOpts(settings, oldRRule)
	if err != nil {
		return &ingest.ConfigureMissingNodesForDeletionSchedulerResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	err = server.workflowManager.UpdateWorkflowScheduleByName(
		ctx, "periodic_missing_nodes_for_deletion", "missing_nodes_for_deletion", updateOpts...)
	if err != nil {
		return &ingest.ConfigureMissingNodesForDeletionSchedulerResponse{}, status.Error(codes.Internal, err.Error())
	}
	return &ingest.ConfigureMissingNodesForDeletionSchedulerResponse{}, nil
}
