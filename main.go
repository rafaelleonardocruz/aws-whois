package main

import (
	"github.com/rafaelleonardocruz/aws-whois/cli"
)

// VERSION holds the version string. Final value injected during build.
var VERSION = "dev"

func main() {
	cli.Execute(VERSION)
}
