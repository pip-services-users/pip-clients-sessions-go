package test_version1

import (
	"os"
	"testing"

	"github.com/pip-services-users/pip-clients-sessions-go/version1"
	"github.com/pip-services3-go/pip-services3-commons-go/config"
)

type sessionsHttpCommandableClientV1Test struct {
	client  *version1.SessionsHttpCommandableClientV1
	fixture *SessionsClientFixtureV1
}

func newSessionsHttpCommandableClientV1Test() *sessionsHttpCommandableClientV1Test {
	return &sessionsHttpCommandableClientV1Test{}
}

func (c *sessionsHttpCommandableClientV1Test) setup(t *testing.T) *SessionsClientFixtureV1 {
	var HTTP_HOST = os.Getenv("HTTP_HOST")
	if HTTP_HOST == "" {
		HTTP_HOST = "localhost"
	}
	var HTTP_PORT = os.Getenv("HTTP_PORT")
	if HTTP_PORT == "" {
		HTTP_PORT = "8080"
	}

	var httpConfig = config.NewConfigParamsFromTuples(
		"connection.protocol", "http",
		"connection.host", HTTP_HOST,
		"connection.port", HTTP_PORT,
	)

	c.client = version1.NewSessionsHttpCommandableClientV1()
	c.client.Configure(httpConfig)
	c.client.Open("")

	c.fixture = NewSessionsClientFixtureV1(c.client)

	return c.fixture
}

func (c *sessionsHttpCommandableClientV1Test) teardown(t *testing.T) {
	c.client.Close("")
}

func TestHttpOpenSession(t *testing.T) {
	c := newSessionsHttpCommandableClientV1Test()
	fixture := c.setup(t)
	defer c.teardown(t)

	fixture.TestOpenSession(t)
}

func TestHttpCloseSession(t *testing.T) {
	c := newSessionsHttpCommandableClientV1Test()
	fixture := c.setup(t)
	defer c.teardown(t)

	fixture.TestCloseSession(t)
}
