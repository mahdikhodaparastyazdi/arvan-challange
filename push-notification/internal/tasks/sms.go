package tasks

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"notification/internal/constants"
	"notification/internal/dto"
	"time"

	log "notification/pkg/logger"

	"github.com/hibiken/asynq"
)

func NewTask(msg dto.Message, typeName string, maxRetry int, timeoutInSec time.Duration) (*asynq.Task, error) {
	p, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}

	return asynq.NewTask(typeName, p, asynq.MaxRetry(maxRetry), asynq.Timeout(timeoutInSec)), nil
}

type Queue struct {
	client       *asynq.Client
	maxRetry     int
	timeoutInSec time.Duration
}

func NewQueue(c *asynq.Client, maxRetry int, timoutInSec time.Duration) *Queue {
	return &Queue{
		client:       c,
		maxRetry:     maxRetry,
		timeoutInSec: timoutInSec,
	}
}

func (q *Queue) Enqueue(msg dto.Message, typeName string) error {
	t, err := NewTask(msg, typeName, q.maxRetry, q.timeoutInSec)
	if err != nil {
		return fmt.Errorf("failed to create task: %w", err)
	}
	_, err = q.client.Enqueue(t, asynq.Queue(typeName))
	if err != nil {
		return fmt.Errorf("asynq enqueue failed: %w", err)
	}

	return nil
}

type Consumer interface {
	Consume(ctx context.Context, msg dto.Message, retry, maxRetry int) error
}

type Worker struct {
	server   *asynq.Server
	consumer Consumer
	logger   log.Logger
}

func NewWorker(s *asynq.Server, c Consumer, logger log.Logger) *Worker {
	return &Worker{
		server:   s,
		consumer: c,
		logger:   logger,
	}
}

func (w *Worker) StartWorker(name string) error {
	mux := asynq.NewServeMux()
	mux.HandleFunc(name, w.HandleTask)

	return w.server.Run(mux)
}

func (w *Worker) HandleTask(ctx context.Context, t *asynq.Task) error {
	retry, _ := asynq.GetRetryCount(ctx)
	maxRetry, _ := asynq.GetMaxRetry(ctx)

	var p dto.Message
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("unmarshalling json: %w", err)
	}
	if err := w.consumer.Consume(ctx, p, retry, maxRetry); err != nil {
		if errors.Is(err, constants.ErrExpiryReached) ||
			errors.Is(err, constants.ErrWrongStatus) ||
			errors.Is(err, constants.ErrBackOffRetry) {
			return nil
		}
		w.logger.Error("failed send notification", log.J{
			"providerName": p.ProviderName,
			"error":        err.Error(),
			"payload":      p,
		})
		return fmt.Errorf("notification task failed: %w", err)
	}

	return nil
}
