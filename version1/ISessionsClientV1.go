package version1

import (
	"github.com/pip-services3-go/pip-services3-commons-go/data"
)

type ISessionsClientV1 interface {
	GetSessions(correlationId string, filter *data.FilterParams,
		paging *data.PagingParams) (result *data.DataPage, err error)

	GetSessionById(correlationId string, id string) (result *SessionV1, err error)

	OpenSession(correlationId string, userId string, userName string,
		address string, client string, user interface{},
		data interface{}) (result *SessionV1, err error)

	StoreSessionData(correlationId string, sessionId string,
		data interface{}) (result *SessionV1, err error)

	UpdateSessionUser(correlationId string, sessionId string,
		user interface{}) (result *SessionV1, err error)

	CloseSession(correlationId string, sessionId string) (result *SessionV1, err error)

	DeleteSessionById(correlationId string, sessionId string) (result *SessionV1, err error)
}
