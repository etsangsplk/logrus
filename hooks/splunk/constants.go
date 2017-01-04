package splunk

import "time"

const (
	driverName                    = "splunk"
	splunkURLKey                  = "splunk-url"
	splunkTokenKey                = "splunk-token"
	splunkSourceKey               = "splunk-source"
	splunkSourceTypeKey           = "splunk-sourcetype"
	splunkIndexKey                = "splunk-index"
	splunkCAPathKey               = "splunk-capath"
	splunkCANameKey               = "splunk-caname"
	splunkInsecureSkipVerifyKey   = "splunk-insecureskipverify"
	splunkFormatKey               = "splunk-format"
	splunkVerifyConnectionKey     = "splunk-verify-connection"
	splunkGzipCompressionKey      = "splunk-gzip"
	splunkGzipCompressionLevelKey = "splunk-gzip-level"
	envKey                        = "env"
	labelsKey                     = "labels"
	tagKey                        = "tag"
)

const (
	// How often do we send messages (if we are not reaching batch size)
	defaultPostMessagesFrequency = 5 * time.Second
	// How big can be batch of messages
	defaultPostMessagesBatchSize = 1000
	// Maximum number of messages we can store in buffer
	defaultBufferMaximum = 10 * defaultPostMessagesBatchSize
	// Number of messages allowed to be queued in the channel
	defaultStreamChannelSize = 4 * defaultPostMessagesBatchSize
)

const (
	envVarPostMessagesFrequency = "SPLUNK_LOGGING_DRIVER_POST_MESSAGES_FREQUENCY"
	envVarPostMessagesBatchSize = "SPLUNK_LOGGING_DRIVER_POST_MESSAGES_BATCH_SIZE"
	envVarBufferMaximum         = "SPLUNK_LOGGING_DRIVER_BUFFER_MAX"
	envVarStreamChannelSize     = "SPLUNK_LOGGING_DRIVER_CHANNEL_SIZE"
)

const (
	splunkFormatRaw    = "raw"
	splunkFormatJSON   = "json"
	splunkFormatInline = "inline"
)
