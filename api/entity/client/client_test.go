package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClientWithoutName(t *testing.T) {
	client := NewClient()
	client.Name = ""
	client.UniqueClientID = 122

	err := client.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "Client's name is empty", err.Error())
}

func TestClientWithoutUniqueClientID(t *testing.T) {
	customer := NewClient()
	customer.Name = "Client Test"
	customer.UniqueClientID = 0

	err := customer.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "Client is invalid", err.Error())
}
