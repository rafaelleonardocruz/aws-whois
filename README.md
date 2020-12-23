# aws-whois 

  ![Testing](https://github.com/rafaelleonardocruz/aws-whois/workflows/Testing/badge.svg)
  ![Generate release](https://github.com/rafaelleonardocruz/aws-whois/workflows/Generate%20release/badge.svg)
  ![CodeQL](https://github.com/rafaelleonardocruz/aws-whois/workflows/CodeQL/badge.svg)

## Motivation

This project intends to find inside AWS accounts which resource is using a specific IP address

## Supported resoucers
- [x] Ec2 Instances
- [x] Ec2 Elastic IP
- [ ] Ec2 ALB/NLB/CLB
- [ ] API Gateway
- [ ] RDS Instance

## Usage
```sh
aws-whois - found which resource has a certain IP address

Usage:
  aws-whois [command]

Available Commands:
  find        find which resouce is using an IP address
  help        Help about any command

Flags:
  -h, --help      help for aws-whois
  -v, --version   version for aws-whois

Use "aws-whois [command] --help" for more information about a command.
```
