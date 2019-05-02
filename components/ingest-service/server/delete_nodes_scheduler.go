package server

// RPC functions for the job: 'delete_node'

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/chef/automate/api/interservice/ingest"
	"github.com/chef/automate/components/ingest-service/config"
	"github.com/chef/automate/lib/workflow"

	log "github.com/sirupsen/logrus"
)

// DeleteMarkedNodes - run the delete node task now
func (server *JobSchedulerServer) DeleteMarkedNodes(ctx context.Context,
	empty *ingest.DeleteMarkedNodesRequest) (*ingest.DeleteMarkedNodesResponse, error) {

	var threshold string
	err := server.workflowManager.GetScheduledWorkflowParameters(ctx, "periodic_delete_nodes", "delete_nodes", &threshold)
	if err != nil {
		return &ingest.DeleteMarkedNodesResponse{}, status.Error(codes.Internal, err.Error())
	}

	deletedNodes, err := server.client.DeleteMarkedNodes(ctx, threshold)
	if err != nil {
		return &ingest.DeleteMarkedNodesResponse{}, status.Error(codes.Internal, err.Error())
	}

	log.WithFields(log.Fields{
		"nodes_deleted": deletedNodes,
	}).Info("Nodes deleted")
	return &ingest.DeleteMarkedNodesResponse{}, nil
}

// StartDeleteNodesScheduler - start the scheduled task of deleting nodes
func (server *JobSchedulerServer) StartDeleteNodesScheduler(ctx context.Context,
	empty *ingest.StartDeleteNodesSchedulerRequest) (*ingest.StartDeleteNodesSchedulerResponse, error) {
	log.WithFields(log.Fields{"func": nameOfFunc()}).Debug("rpc call")

	err := server.workflowManager.UpdateWorkflowScheduleByName(
		ctx, "periodic_delete_nodes", "delete_nodes", workflow.UpdateEnabled(true))
	if err != nil {
		return &ingest.StartDeleteNodesSchedulerResponse{}, status.Error(codes.Internal, err.Error())
	}
	return &ingest.StartDeleteNodesSchedulerResponse{}, nil
}

// StopDeleteNodesScheduler - stop the scheduled delete node task from running
func (server *JobSchedulerServer) StopDeleteNodesScheduler(ctx context.Context,
	empty *ingest.StopDeleteNodesSchedulerRequest) (*ingest.StopDeleteNodesSchedulerResponse, error) {
	log.WithFields(log.Fields{"func": nameOfFunc()}).Debug("rpc call")

	err := server.workflowManager.UpdateWorkflowScheduleByName(
		ctx, "periodic_delete_nodes", "delete_nodes", workflow.UpdateEnabled(false))
	if err != nil {
		return &ingest.StopDeleteNodesSchedulerResponse{}, status.Error(codes.Internal, err.Error())
	}
	return &ingest.StopDeleteNodesSchedulerResponse{}, nil
}

// ConfigureDeleteNodesScheduler changed the configurations for the delete node scheduler
func (server *JobSchedulerServer) ConfigureDeleteNodesScheduler(ctx context.Context,
	settings *ingest.JobSettings) (*ingest.ConfigureDeleteNodesSchedulerResponse, error) {
	log.WithFields(log.Fields{
		"Freq":      settings.Every,
		"Running":   settings.Running,
		"Threshold": settings.Threshold,
	}).Info("ConfigureDeleteNodesScheduler")

	oldRRule, err := server.workflowManager.GetScheduledWorkflowRecurrence(
		ctx, "periodic_delete_nodes", "delete_nodes")
	if err != nil {
		return &ingest.ConfigureDeleteNodesSchedulerResponse{}, status.Error(codes.Internal, err.Error())
	}

	updateOpts, err := config.JobSettingsToUpdateOpts(settings, oldRRule)
	if err != nil {
		return &ingest.ConfigureDeleteNodesSchedulerResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	err = server.workflowManager.UpdateWorkflowScheduleByName(
		ctx, "periodic_delete_nodes", "delete_nodes", updateOpts...)
	if err != nil {
		return &ingest.ConfigureDeleteNodesSchedulerResponse{}, status.Error(codes.Internal, err.Error())
	}
	return &ingest.ConfigureDeleteNodesSchedulerResponse{}, nil
}
