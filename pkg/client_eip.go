package pkg

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type ElasticIp struct {
	AllocationID string
	PuclicIP     string
}

func GetElasticIps() ([]ElasticIp, error) {
	var result []ElasticIp
	svc := ec2.New(session.New())
	input := &ec2.DescribeAddressesInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String("domain"),
				Values: []*string{
					aws.String("vpc"),
				},
			},
		},
	}

	resp, err := svc.DescribeAddresses(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				return nil, errors.New(aerr.Error())
			}
		} else {
			return nil, errors.New("Error on ec2 client")
		}
		return nil, errors.New("Error on ec2 client")
	}

	for _, addr := range resp.Addresses {
		result = append(result, ElasticIp{*addr.AllocationId, *addr.PublicIp})

	}
	return (result), nil
}
