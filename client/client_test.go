package client

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func initTestClient(t *testing.T) *LearnClient {
	Init("learnctl", VERSION, "testdata/config.yaml")
	client := NewLearnClient("127.0.0.1", 6379)
	require.NotNil(t, client)
	return client
}

func TestPing(t *testing.T) {
	client := initTestClient(t)
	err := client.Ping()
	require.Nil(t, err)
}
