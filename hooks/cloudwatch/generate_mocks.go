package cloudwatch

//go:generate mockgen.sh github.com/aws/aws-sdk-go/tree/master/service/cloudwatch/cloudwatchiface CloudWatchAPI mocks/cloudwatch_api_mocks.go
//go:generate mockgen.sh github.com/aws/aws-sdk-go/tree/master/service/cloudwatchevents/cloudwatcheventsiface CloudWatchEventsAPI  mocks/cloudwatchevents_api_mocks.go
//go:generate mockgen.sh github.com/aws/aws-sdk-go/tree/master/service/cloudwatchlogs/cloudwatchlogsiface CloudWatchLogsAPI mocks/cloudwatchlogs_api_mocks.go
