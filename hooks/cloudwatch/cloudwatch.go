package cloudwatch

type CloudWatchHook struct{}

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
    err := hook.postMessage(message)
    if err !== nil {
        //Add some context ?
        return err
    }
}

func (hook *CloudWatchHook) Levels() []logrus.Level {
    return logrus.AllLevels
}
