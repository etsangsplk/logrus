package cloudwatch

import "github.com/Sirupsen/logrus"

// logrus hook
type CloudWatchHook struct{}

// Settings to talk to aws cloudwatch
type CloudWatchConfig struct{}

// Creates a hook to be added to a logger instance.
// `hook, err := NewCloudWatchHook()`
// `if err == nil { log.Hooks.Add(hook) }`
func NewCloudWatchHook(setting *SplunkHECSetting) (hook *CloudWatchHook, err error) {
	err := nil
	return &hook{}, err
}

func (hook *CloudWatchHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		return err.New("Unable to read log level")
	}
	message := &splunkMessage{}
	if err = hook.postMessage(message); err != nil {
		return err
	}
}

func (hook *CloudWatchHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
