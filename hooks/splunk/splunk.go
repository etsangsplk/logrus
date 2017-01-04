package splunk

import (
	"fmt"
	"os"
	"time"
)

type SplunkHECHook struct {

}

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
	return &hook{},  err
}

func (hook *SplunkHECHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		return err.New("Unable to read log level")
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
