package localsqs

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

const (
	sqsEndp = "http://localhost:4576"
)

// Pong TODO
func Pong() string {
	return "pong"
}

func sqsClient() *sqs.SQS {
	sess := session.Must(session.NewSession(&aws.Config{
		Region:   aws.String("us-east-1"),
		Endpoint: aws.String(sqsEndp),
	}))
	return sqs.New(sess)
}

// Create TODO
func Create(name string) (string, error) {
	client := sqsClient()

	out, err := client.CreateQueue(&sqs.CreateQueueInput{
		QueueName: aws.String(name),
	})
	if err != nil {
		return "", err
	}

	return aws.StringValue(out.QueueUrl), nil
}
