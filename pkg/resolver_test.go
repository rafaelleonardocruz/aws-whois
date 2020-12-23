package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var validAddressPublic string = "18.211.191.170"
var validAddressPrivate string = "10.1.0.0"
var validAddressGoogle string = "8.8.8.8"
var invalidAddress string = "8.8.8"
var validIpv6Address string = "fe80::38e7:b57:59e:1773"
var validHostAws string = "ec2-54-207-114-71.sa-east-1.compute.amazonaws.com"
var invalidHostAws string = "google.com"

func TestIsValidAddress(t *testing.T) {

	var resBol bool
	resBol, _ = isValidAddress(validAddressPublic)
	assert.Equal(t, resBol, true, "The address must be valid")
	resBol, _ = isValidAddress(validAddressPrivate)
	assert.Equal(t, resBol, true, "The address must be valid")
	resBol, _ = isValidAddress(validAddressGoogle)
	assert.Equal(t, resBol, true, "The address must be valid")
	resBol, _ = isValidAddress(invalidAddress)
	assert.Equal(t, resBol, false, "The address must be invalid")
	resBol, _ = isValidAddress(validIpv6Address)
	assert.Equal(t, resBol, false, "The address must be invalid")
}

func TestPrivateAddress(t *testing.T) {
	var resBol bool

	resBol, _ = isPrivateAddress(validAddressPublic)
	assert.Equal(t, resBol, false, "The address must be public")
	resBol, _ = isPrivateAddress(validAddressPrivate)
	assert.Equal(t, resBol, true, "The address must be public")
	resBol, _ = isPrivateAddress(validAddressGoogle)
	assert.Equal(t, resBol, false, "The address must be private")
}

func TestGetReverseDNSName(t *testing.T) {
	var err error

	_, err = getReverseDNSName(validAddressPublic)
	assert.Nil(t, err, "The address must be resolvable")
	_, err = getReverseDNSName(validAddressPrivate)
	assert.NotNil(t, err, "The address must be resolvable")
	_, err = getReverseDNSName(validAddressGoogle)
	assert.Nil(t, err, "The address must be resolvable")
	_, err = getReverseDNSName(invalidAddress)
	assert.NotNil(t, err, "The address must be resolvable")
	_, err = isValidAddress(validIpv6Address)
	assert.NotNil(t, err, "The address must be resolvable")
}

func TestIsAwsAddress(t *testing.T) {
	var err error

	_, err = isAwsAddress(validHostAws)
	assert.Nil(t, err, "The address must be owned by AWS")
	_, err = isAwsAddress(invalidHostAws)
	assert.NotNil(t, err, "The address must be owmed by AWS")
}
