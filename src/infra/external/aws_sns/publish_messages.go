package aws_sns

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"log"
)

type PubSubInterface struct {
}

func (r PubSubInterface) NotifyPaymentApproved(orderId uint) error {
	return nil
}

func (r PubSubInterface) NotifyPaymentError(orderId uint) error {
	return nil
}

// Publish publishes a message to an Amazon SNS topic. The message is then sent to all
// subscribers. When the topic is a FIFO topic, the message must also contain a group ID
// and, when ID-based deduplication is used, a deduplication ID. An optional key-value
// filter attribute can be specified so that the message can be filtered according to
// a filter policy.
func Publish(snsClient sns.Client, topicArn string, message string) error {
	publishInput := sns.PublishInput{TopicArn: aws.String(topicArn), Message: aws.String(message)}
	_, err := snsClient.Publish(context.TODO(), &publishInput)
	if err != nil {
		log.Printf("Couldn't publish message to topic %v. Here's why: %v", topicArn, err)
	} else {
		log.Printf("Published!")
	}
	return err
}

// SubscribeQueue subscribes an Amazon Simple Queue Service (Amazon SQS) queue to an
// Amazon SNS topic. When filterMap is not nil, it is used to specify a filter policy
// so that messages are only sent to the queue when the message has the specified attributes.
func SubscribeQueue(snsClient sns.Client, topicArn string) (string, error) {
	var subscriptionArn string
	var attributes map[string]string

	output, err := snsClient.Subscribe(context.TODO(), &sns.SubscribeInput{
		Protocol:              aws.String("sqs"),
		TopicArn:              aws.String(topicArn),
		Attributes:            attributes,
		ReturnSubscriptionArn: true,
	})
	if err != nil {
		log.Printf("Couldn't susbscribe to topic %v. Here's why: %v\n", topicArn, err)
	} else {
		subscriptionArn = *output.SubscriptionArn
	}

	return subscriptionArn, err
}
