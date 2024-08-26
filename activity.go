package personal

import (
	"context"
	"time"

	"go.temporal.io/sdk/activity"
)

// @@@SNIPSTART samples-go-cancellation-activity-definition
type Activities struct{}

func (a *Activities) ActivityToBeCancelled(ctx context.Context) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("activity started, to cancel the Workflow Execution, use 'go run cancellation/cancel/main.go " +
		"-w <WorkflowID>' or use the CLI: 'tctl wf cancel -w <WorkflowID>'")
	for {
		// TODO: Make your own project to test the inserted heartbeat here
		// insert heartbeat
		activity.RecordHeartbeat(ctx, "HAHAHAHAHAHAHAHAHAHAHAHAHAHAHAHAHAHA")
		select {
		case <-time.After(1 * time.Second):
			logger.Info("heartbeating...")
			activity.RecordHeartbeat(ctx, "")
		case <-ctx.Done():
			logger.Info("context is cancelled")
			return "I am cancelled by Done", nil
		}
	}
}

func (a *Activities) CleanupActivity(ctx context.Context) error {
	logger := activity.GetLogger(ctx)
	logger.Info("Cleanup Activity started")
	return nil
}

func (a *Activities) ActivityToBeSkipped(ctx context.Context) error {
	logger := activity.GetLogger(ctx)
	logger.Info("this Activity will be skipped due to cancellation")
	return nil
}

// @@@SNIPEND
