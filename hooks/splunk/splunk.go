package splunk

import (
	"errors"
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
func NewSplunkHECHook(setting *SplunkHECConfig) (hook *SplunkHECHook, err error) {
	err = nil
	return &SplunkHECHook{}, err
}

func (hook *SplunkHECHook) Fire(entry *logrus.Entry) error {
	_, err := entry.String()
	if err != nil {
		return errors.New("Unable to read log level")
	}
	message := &splunkMessage{}
	if err = hook.post(message); err != nil {
		return err
	}
	return nil
}

func (hook *SplunkHECHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (hook *SplunkHECHook) post(message *splunkMessage) error {
	return nil
}
