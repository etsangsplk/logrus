package cloudwatch

import (
	"errors"

	"github.com/Sirupsen/logrus"
)

// logrus hook
type CloudWatchHook struct{}

// Settings to talk to aws cloudwatch
type CloudWatchConfig struct{}

// Creates a hook to be added to a logger instance.
// `hook, err := NewCloudWatchHook()`
// `if err == nil { log.Hooks.Add(hook) }`
func NewCloudWatchHook(setting *CloudWatchConfig) (hook *CloudWatchHook, err error) {
	err = nil
	return &CloudWatchHook{}, err
}

func (hook *CloudWatchHook) Fire(entry *logrus.Entry) error {
	message, err := entry.String()
	if err != nil {
		return errors.New("Unable to read log level")
	}

	if err = hook.post(&message); err != nil {
		return err
	}
	return nil
}

func (hook *CloudWatchHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (hook *CloudWatchHook) post(message *string) error {
	return nil
}
