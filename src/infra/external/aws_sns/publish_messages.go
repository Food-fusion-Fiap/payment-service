package aws_sns

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"log"
	"os"
)

type PubSubInterface struct {
}

func (r PubSubInterface) NotifyPaymentApproved(orderId string) error {
	sdkConfig, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		fmt.Println("Couldn't load default configuration. Have you set up your AWS account?")
		fmt.Println(err)
		return err
	}
	log.Print("veio aqui 01")
	snsClient := sns.NewFromConfig(sdkConfig)
	log.Print("veio aqui 02")
	topicArn := os.Getenv("PAYMENT_SUCCESS_SNS_TOPIC_ARN")
	log.Print("veio aqui 03")
	message := fmt.Sprintf("{\"orderId\": \"%s\",  \"changeToStatus\": \"received\"}", orderId)
	log.Print("veio aqui 04")

	err = Publish(*snsClient, topicArn, message)
	log.Print("veio aqui 05")
	if err != nil {
		return err
	}
	log.Print("veio aqui 06")
	return nil
}

func (r PubSubInterface) NotifyPaymentError(orderId string) error {
	sdkConfig, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		fmt.Println("Couldn't load default configuration. Have you set up your AWS account?")
		fmt.Println(err)
		return err
	}
	snsClient := sns.NewFromConfig(sdkConfig)
	topicArn := os.Getenv("PAYMENT_FAILURE_SNS_TOPIC_ARN")
	message := fmt.Sprintf("{\"orderId\": \"%s\",  \"changeToStatus\": \"cancelled\"}", orderId)

	err = Publish(*snsClient, topicArn, message)
	if err != nil {
		return err
	}

	return nil
}

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
