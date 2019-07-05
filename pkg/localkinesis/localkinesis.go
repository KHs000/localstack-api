package localkinesis

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesis"
	"github.com/google/uuid"
)

const (
	knsEndp = "http://localhost:4568"
)

// Client TODO
type Client struct {
	Kns *kinesis.Kinesis
}

func kinesisClient() *kinesis.Kinesis {
	sess := session.Must(session.NewSession(&aws.Config{
		Region:   aws.String("us-east-1"),
		Endpoint: aws.String(knsEndp),
	}))
	return kinesis.New(sess)
}

// NewClient TODO
func NewClient() Client {
	return Client{
		Kns: kinesisClient(),
	}
}

// Create TODO
func (c Client) Create(name string) error {
	_, err := c.Kns.CreateStream(&kinesis.CreateStreamInput{
		ShardCount: aws.Int64(1),
		StreamName: aws.String(name),
	})
	if err != nil {
		return err
	}
	return nil
}

// List TODO
func (c Client) List() ([]string, error) {
	out, err := c.Kns.ListStreams(&kinesis.ListStreamsInput{})
	if err != nil {
		return nil, err
	}

	var names []string
	for _, v := range out.StreamNames {
		names = append(names, aws.StringValue(v))
	}

	return names, nil
}

// PutRecord TODO
func (c Client) PutRecord(data []byte, name string) error {
	_, err := c.Kns.PutRecord(&kinesis.PutRecordInput{
		Data:         data,
		StreamName:   aws.String(name),
		PartitionKey: aws.String(uuid.New().String()),
	})
	if err != nil {
		return err
	}
	return nil
}
