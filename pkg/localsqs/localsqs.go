package localsqs

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

const (
	sqsEndp = "http://localhost:4576"
)

// Client TODO
type Client struct {
	Sqs *sqs.SQS
}

func sqsClient() *sqs.SQS {
	sess := session.Must(session.NewSession(&aws.Config{
		Region:   aws.String("us-east-1"),
		Endpoint: aws.String(sqsEndp),
	}))
	return sqs.New(sess)
}

// NewClient TODO
func NewClient() Client {
	return Client{
		Sqs: sqsClient(),
	}
}

// Create TODO
func (c Client) Create(name string) (string, error) {
	out, err := c.Sqs.CreateQueue(&sqs.CreateQueueInput{
		QueueName: aws.String(name),
	})
	if err != nil {
		return "", err
	}

	return aws.StringValue(out.QueueUrl), nil
}

// GetAttributes TODO
func (c Client) GetAttributes(url string, attr ...string) (map[string]string, error) {
	var atn []*string
	for _, t := range attr {
		atn = append(atn, aws.String(t))
	}

	out, err := c.Sqs.GetQueueAttributes(&sqs.GetQueueAttributesInput{
		AttributeNames: atn,
		QueueUrl:       aws.String(url),
	})
	if err != nil {
		return nil, err
	}

	tm := make(map[string]string)
	for k, v := range out.Attributes {
		tm[k] = aws.StringValue(v)
	}

	return tm, nil
}

// List TODO
func (c Client) List() ([]string, error) {
	out, err := c.Sqs.ListQueues(&sqs.ListQueuesInput{})
	if err != nil {
		return nil, err
	}

	return aws.StringValueSlice(out.QueueUrls), nil
}

// Purge TODO
func (c Client) Purge(url string) error {
	_, err := c.Sqs.PurgeQueue(&sqs.PurgeQueueInput{
		QueueUrl: aws.String(url),
	})
	if err != nil {
		return err
	}
	return nil
}
