package cron

import (
	"context"

	"github.com/TIBCOSoftware/flogo-lib/core/trigger"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/robfig/cron"
)

var log = logger.GetLogger("trigger-flogo-cron")

// CronTriggerFactory My Trigger factory
type CronTriggerFactory struct {
	metadata *trigger.Metadata
}

// NewFactory create a new Trigger factory
func NewFactory(md *trigger.Metadata) trigger.Factory {
	return &CronTriggerFactory{metadata: md}
}

// New Creates a new trigger instance for a given id
func (t *CronTriggerFactory) New(config *trigger.Config) trigger.Trigger {
	return &CronTrigger{metadata: t.metadata, config: config}
}

// CronTrigger is a stub for your Trigger implementation
type CronTrigger struct {
	metadata *trigger.Metadata
	config   *trigger.Config
	cron     *cron.Cron
	handlers []*trigger.Handler
}

// Initialize implements trigger.Init.Initialize
func (t *CronTrigger) Initialize(ctx trigger.InitContext) error {
	t.cron = cron.New()
	t.handlers = ctx.GetHandlers()
	return nil
}

// Metadata implements trigger.Trigger.Metadata
func (t *CronTrigger) Metadata() *trigger.Metadata {
	return t.metadata
}

// Start implements trigger.Trigger.Start
func (t *CronTrigger) Start() error {
	// start the trigger
	expression := t.config.Settings["expression"].(string)

	for _, handler := range t.handlers {

		cmd := func() {
			_, err := handler.Handle(context.Background(), nil)
			if err != nil {
				log.Error("Error running handler: ", err.Error())
			}
		}

		err := t.cron.AddFunc(expression, cmd)

		if err != nil {
			log.Errorf("Error adding cron handler with expression %s : %s", expression, err.Error())
			return err
		}
	}

	t.cron.Start()

	return nil
}

// Stop implements trigger.Trigger.Start
func (t *CronTrigger) Stop() error {
	// stop the trigger
	t.cron.Stop()
	return nil
}
