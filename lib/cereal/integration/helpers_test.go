package integration_test

import (
	"context"
	"time"

	"github.com/stretchr/testify/suite"

	"github.com/chef/automate/lib/cereal"
)

type taskExecutor struct {
	Name     string
	Executor cereal.TaskExecutor
}
type workflowExecutor struct {
	Name     string
	Executor cereal.WorkflowExecutor
}

type managerOpt struct {
	TaskExecutors     []taskExecutor
	WorkflowExecutors []workflowExecutor
	NoStart           bool
}
type managerOptFunc func(*managerOpt)

func WithWorkflowExecutor(name string, executor cereal.WorkflowExecutor) managerOptFunc {
	return func(o *managerOpt) {
		o.WorkflowExecutors = append(o.WorkflowExecutors, workflowExecutor{
			Name:     name,
			Executor: executor,
		})
	}
}

func WithTaskExecutor(name string, executor cereal.TaskExecutor) managerOptFunc {
	return func(o *managerOpt) {
		o.TaskExecutors = append(o.TaskExecutors, taskExecutor{
			Name:     name,
			Executor: executor,
		})
	}
}

func WithNoStart() managerOptFunc {
	return func(o *managerOpt) {
		o.NoStart = true
	}
}

type taskExecutorWrapper struct {
	f func(context.Context, cereal.Task) (interface{}, error)
}

func (w *taskExecutorWrapper) Run(ctx context.Context, task cereal.Task) (interface{}, error) {
	return w.f(ctx, task)
}

func WithTaskExecutorF(name string, f func(context.Context, cereal.Task) (interface{}, error)) managerOptFunc {
	return WithTaskExecutor(name, &taskExecutorWrapper{f})
}

type workflowExecutorWrapper struct {
	onStart        func(w cereal.WorkflowInstance, ev cereal.StartEvent) cereal.Decision
	onTaskComplete func(w cereal.WorkflowInstance, ev cereal.TaskCompleteEvent) cereal.Decision
	onCancel       func(w cereal.WorkflowInstance, ev cereal.CancelEvent) cereal.Decision
}

func (wrapper *workflowExecutorWrapper) OnStart(w cereal.WorkflowInstance, ev cereal.StartEvent) cereal.Decision {
	return wrapper.onStart(w, ev)
}

func (wrapper *workflowExecutorWrapper) OnTaskComplete(w cereal.WorkflowInstance,
	ev cereal.TaskCompleteEvent) cereal.Decision {
	return wrapper.onTaskComplete(w, ev)
}

func (wrapper *workflowExecutorWrapper) OnCancel(w cereal.WorkflowInstance,
	ev cereal.CancelEvent) cereal.Decision {
	return wrapper.onCancel(w, ev)
}

func randName(name string) string {
	now := time.Now()
	return name + now.String()
}

type CerealTestSuite struct {
	suite.Suite
	newManager func(...managerOptFunc) *cereal.Manager
}
