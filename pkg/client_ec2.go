package pkg

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type Instance struct {
	InstanceID string
	IpAddress  string
}

func GetInstances() ([]Instance, error) {
	var result []Instance

	ec2client := ec2.New(session.New())
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
