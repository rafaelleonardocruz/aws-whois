package pkg

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

//GetElasticIps return EIPs for a requested region
func GetElasticIps(region string) error {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)

	if err != nil {
		return errors.New("Error initializing an AWS session")
	}

	ec2client := ec2.New(sess)
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

	resp, err := ec2client.DescribeAddresses(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				return errors.New(aerr.Error())
			}
		} else {
			return errors.New("Error on ec2 client")
		}
	}

	for _, addr := range resp.Addresses {
		ResourceList = append(ResourceList, Resource{"ElasticIp", *addr.AllocationId, *addr.PublicIp})
	}
	return nil
}
