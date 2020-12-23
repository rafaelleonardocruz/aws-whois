package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTryMatch(t *testing.T) {

	mustBeFound1 := "1.1.1.1"
	mustBeFound2 := "3.3.3.3"
	mustBeNotFound1 := "5.5.5.5"
	mustBeNotFound2 := "6.6.6.6"

	ResourceList = append(ResourceList, Resource{"EC2", "i-12345", "1.1.1.1"})
	ResourceList = append(ResourceList, Resource{"EC2", "i-67890", "2.2.2.2"})
	ResourceList = append(ResourceList, Resource{"EIP", "i-12345", "3.3.3.3"})
	ResourceList = append(ResourceList, Resource{"EIP", "i-67890", "4.4.4.4"})

	var res Resource
	res, _ = TryMatch(mustBeFound1)
	assert.Equal(t, res.PublicIP, mustBeFound1, "The address must be equal")

	res, _ = TryMatch(mustBeFound2)
	assert.Equal(t, res.PublicIP, mustBeFound2, "The address must be equal")

	res, _ = TryMatch(mustBeNotFound1)
	assert.NotEqual(t, res.PublicIP, mustBeNotFound1, "The address must be not found")

	res, _ = TryMatch(mustBeNotFound2)
	assert.NotEqual(t, res.PublicIP, mustBeNotFound2, "The address must be not found")
}
