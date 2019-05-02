package server

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"

	"github.com/chef/automate/components/ingest-service/backend"
	"github.com/chef/automate/lib/workflow"
)

// TODO(ssd) 2019-05-15: This is a helper to avoid having to write
// workflows for things that are just single tasks. Perhaps the
// workflow library could have a helper that..helps with this.
type SingleTaskWorkflow struct {
	taskName string
}

func NewSingleTaskWorkflow(taskName string) *SingleTaskWorkflow {
	return &SingleTaskWorkflow{taskName}
}

func (s *SingleTaskWorkflow) OnStart(w workflow.WorkflowInstanceHandler, ev workflow.StartEvent) workflow.Decision {
	var params string
	err := w.GetParameters(&params)
	if err != nil {
		logrus.WithError(err).Error("failed to get parameters")
		w.Complete()
	}

	w.EnqueueTask(s.taskName, params)
	return w.Continue(0)
}

func (s *SingleTaskWorkflow) OnTaskComplete(w workflow.WorkflowInstanceHandler,
	ev workflow.TaskCompleteEvent) workflow.Decision {
	return w.Complete()
}

func (s *SingleTaskWorkflow) OnCancel(w workflow.WorkflowInstanceHandler,
	ev workflow.CancelEvent) workflow.Decision {
	return w.Complete()
}

type DeleteExpiredMarkedNodesTask struct {
	Client backend.Client
}

func (t *DeleteExpiredMarkedNodesTask) Run(ctx context.Context, task workflow.TaskQuerier) (interface{}, error) {
	var threshold string
	err := task.GetParameters(&threshold)
	if err != nil {
		return nil, errors.Wrap(err, "could not get threshold parameter")
	}

	return nil, deleteExpiredMarkedNodes(ctx, threshold, t.Client)
}

type MarkNodesMissingTask struct {
	Client backend.Client
}

func (t *MarkNodesMissingTask) Run(ctx context.Context, task workflow.TaskQuerier) (interface{}, error) {
	var threshold string
	err := task.GetParameters(&threshold)
	if err != nil {
		return nil, errors.Wrap(err, "could not get threshold parameter")
	}

	return nil, markNodesMissing(ctx, threshold, t.Client)
}

type MarkMissingNodesForDeletionTask struct {
	Client backend.Client
}

func (t *MarkMissingNodesForDeletionTask) Run(ctx context.Context, task workflow.TaskQuerier) (interface{}, error) {
	var threshold string
	err := task.GetParameters(&threshold)
	if err != nil {
		return nil, errors.Wrap(err, "could not get threshold parameter")
	}

	return nil, markMissingNodesForDeletion(ctx, threshold, t.Client)
}

// markNodesMissing is a job that will call the backend to mark
// all nodes that haven't checked in passed the threshold
func markNodesMissing(ctx context.Context, threshold string, client backend.Client) error {
	logctx := log.WithFields(log.Fields{
		"threshold": threshold,
		"job":       "MarkMissingNodesForDeletion",
	})

	logctx.Debug("Starting job")
	updateCount, err := client.MarkNodesMissing(ctx, threshold)
	if err != nil {
		logctx.WithError(err).Error("Job failed")
		return err
	}

	f := log.Fields{"nodes_updated": updateCount}
	if updateCount > 0 {
		logctx.WithFields(f).Info("Job updated nodes")
	} else {
		logctx.WithFields(f).Debug("Job ran without updates")
	}
	return nil
}

// markMissingNodesForDeletion is a job that will call the backend to mark all missing nodes
// that haven't checked in passed the threshold ready for deletion
func markMissingNodesForDeletion(ctx context.Context, threshold string, client backend.Client) error {
	logctx := log.WithFields(log.Fields{
		"threshold": threshold,
		"job":       "MarkMissingNodesForDeletion",
	})

	logctx.Debug("Starting job")
	updateCount, err := client.MarkMissingNodesForDeletion(ctx, threshold)
	if err != nil {
		log.WithError(err).Error("Job failed")
		return err
	}

	f := log.Fields{"nodes_updated": updateCount}
	if updateCount > 0 {
		logctx.WithFields(f).Info("Job updated nodes")
	} else {
		logctx.WithFields(f).Debug("Job ran without updates")
	}
	return nil
}

// deleteExpiredMarkedNodes is a job that will call the backend to delete all expired
// nodes marked for deletion
func deleteExpiredMarkedNodes(ctx context.Context, threshold string, client backend.Client) error {
	logctx := log.WithFields(log.Fields{
		"job":       "DeleteExpiredMarkedNodes",
		"threshold": threshold,
	})

	logctx.Debug("Starting Job")
	updateCount, err := client.DeleteMarkedNodes(ctx, threshold)
	if err != nil {
		logctx.WithError(err).Error("Job Failed")
		return err
	}

	f := log.Fields{"nodes_deleted": updateCount}
	if updateCount > 0 {
		logctx.WithFields(f).Info("Job deleted nodes")
	} else {
		logctx.WithFields(f).Debug("Job ran without updates")
	}
	return nil
}
