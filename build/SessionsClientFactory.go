package build

import (
	clients1 "github.com/pip-services-users/pip-clients-sessions-go/version1"
	cref "github.com/pip-services3-go/pip-services3-commons-go/refer"
	cbuild "github.com/pip-services3-go/pip-services3-components-go/build"
)

type SessionsClientFactory struct {
	cbuild.Factory
}

func NewSessionsClientFactory() *SessionsClientFactory {
	c := &SessionsClientFactory{
		Factory: *cbuild.NewFactory(),
	}

	// nullClientDescriptor := cref.NewDescriptor("pip-services-sessions", "client", "null", "*", "1.0")
	// directClientDescriptor := cref.NewDescriptor("pip-services-sessions", "client", "direct", "*", "1.0")
	cmdHttpClientDescriptor := cref.NewDescriptor("pip-services-sessions", "client", "commandable-http", "*", "1.0")
	grpcClientDescriptor := cref.NewDescriptor("pip-services-sessions", "client", "grpc", "*", "1.0")

	// c.RegisterType(nullClientDescriptor, clients1.NewSessionsNullClientV1)
	// c.RegisterType(directClientDescriptor, clients1.NewSessionsDirectClientV1)
	c.RegisterType(cmdHttpClientDescriptor, clients1.NewSessionsHttpCommandableClientV1)
	c.RegisterType(grpcClientDescriptor, clients1.NewSessionGrpcClientV1)

	return c
}
