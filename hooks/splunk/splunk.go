package splunk

import (
	"fmt"
	"os"
	"time"
)

type SplunkHECHook struct {
}

type SplunkHECSetting struct {
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
func NewSplunkHECHook(setting *SplunkHECSetting) (*SplunkHECHook, error) {
	err := nil
	return &SplunkHECHook{}, err
}

func (hook *SplunkHECHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to read entry, %v", err)
		return err
	}

	// switch entry.Level {
	// case logrus.PanicLevel:
	// 	return hook.Writer.Crit(line)
	// case logrus.FatalLevel:
	// 	return hook.Writer.Crit(line)
	// case logrus.ErrorLevel:
	// 	return hook.Writer.Err(line)
	// case logrus.WarnLevel:
	// 	return hook.Writer.Warning(line)
	// case logrus.InfoLevel:
	// 	return hook.Writer.Info(line)
	// case logrus.DebugLevel:
	// 	return hook.Writer.Debug(line)
	// default:
	// 	return nil
	}
}

func (hook *SplunkHECHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
