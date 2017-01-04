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

	// switch entry.Level {
	// case logrus.PanicLevel:
	//  return hook.Writer.Crit(line)
	// case logrus.FatalLevel:
	//  return hook.Writer.Crit(line)
	// case logrus.ErrorLevel:
	//  return hook.Writer.Err(line)
	// case logrus.WarnLevel:
	//  return hook.Writer.Warning(line)
	// case logrus.InfoLevel:
	//  return hook.Writer.Info(line)
	// case logrus.DebugLevel:
	//  return hook.Writer.Debug(line)
	// default:
	//  return nil
	//}
}

func (hook *CloudWatchHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
