package pkg

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type Instance struct {
	InstanceID string
	IPAddress  string
}

//GetInstances function get ec2-instances to a specifc AWS Region
func GetInstances(region string) ([]Instance, error) {

	var result []Instance

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)

	if err != nil {
		return nil, errors.New("Error initializing an AWS session")
	}

	ec2client := ec2.New(sess)

	resp, err := ec2client.DescribeInstances(nil)

	if err != nil {
		return nil, errors.New("Error on ec2 client")
	}

	for i, _ := range resp.Reservations {
		for _, instance := range resp.Reservations[i].Instances {
			if *instance.PublicDnsName != "" {
				result = append(result, Instance{*instance.InstanceId, *instance.PublicIpAddress})
			}
		}
	}

	return (result), nil
}
