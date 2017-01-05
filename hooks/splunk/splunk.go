package splunk

import (
	"time"

	"github.com/Sirupsen/logrus"
)

// This is a hook for logrus
type SplunkHECHook struct {
}

// Config settings to talk to splunk hec
type SplunkHECConfig struct {
	url         string
	auth        string
	nullMessage *splunkMessage

	// http compression
	gzipCompression      bool
	gzipCompressionLevel int

	// Advanced options
	postMessagesFrequency time.Duration
	postMessagesBatchSize int
	bufferMaximum         int
}

// Creates a hook to be added to a logger instance.
// `hook, err := NewSplunkHECHook()`
// `if err == nil { log.Hooks.Add(hook) }`
func NewSplunkHECHook(setting *SplunkHECSetting) (hook *SplunkHECHook, err error) {
	err := nil
	return &hook{}, err
}

func (hook *SplunkHECHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		return err.New("Unable to read log level")
	}
	message := &splunkMessage{}
	if err = hook.postMessage(message); err != nil {
		return err
	}
}

func (hook *SplunkHECHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
