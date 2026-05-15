package client

import (
	"github.com/stretchr/testify/require"
	"log"
	"os"
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

func TestClasses(t *testing.T) {
	client := initTestClient(t)
	address := os.Getenv("TEST_ADDRESS")
	require.NotEmpty(t, address)
	classes, err := client.Classes(address)
	require.Nil(t, err)
	log.Println(FormatJSON(classes))
}
