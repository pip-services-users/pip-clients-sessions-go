package test_version1

import (
	"os"
	"testing"

	"github.com/pip-services-users/pip-clients-sessions-go/version1"
	"github.com/pip-services3-go/pip-services3-commons-go/config"
)

type sessionGrpcClientV1Test struct {
	client  *version1.SessionGrpcClientV1
	fixture *SessionsClientFixtureV1
}

func newSessionGrpcClientV1Test() *sessionGrpcClientV1Test {
	return &sessionGrpcClientV1Test{}
}

func (c *sessionGrpcClientV1Test) setup(t *testing.T) *SessionsClientFixtureV1 {
	var GRPC_HOST = os.Getenv("GRPC_HOST")
	if GRPC_HOST == "" {
		GRPC_HOST = "localhost"
	}
	var GRPC_PORT = os.Getenv("GRPC_PORT")
	if GRPC_PORT == "" {
		GRPC_PORT = "8090"
	}

	var httpConfig = config.NewConfigParamsFromTuples(
		"connection.protocol", "http",
		"connection.host", GRPC_HOST,
		"connection.port", GRPC_PORT,
	)

	c.client = version1.NewSessionGrpcClientV1()
	c.client.Configure(httpConfig)
	c.client.Open("")

	c.fixture = NewSessionsClientFixtureV1(c.client)

	return c.fixture
}

func (c *sessionGrpcClientV1Test) teardown(t *testing.T) {
	c.client.Close("")
}

func TestGrpcOpenSession(t *testing.T) {
	c := newSessionGrpcClientV1Test()
	fixture := c.setup(t)
	defer c.teardown(t)

	fixture.TestOpenSession(t)
}

func TestGrpcCloseSession(t *testing.T) {
	c := newSessionGrpcClientV1Test()
	fixture := c.setup(t)
	defer c.teardown(t)

	fixture.TestCloseSession(t)
}
