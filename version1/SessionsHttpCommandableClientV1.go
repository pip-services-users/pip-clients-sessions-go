package version1

import (
	"reflect"

	"github.com/pip-services3-go/pip-services3-commons-go/data"
	cdata "github.com/pip-services3-go/pip-services3-commons-go/data"
	cclients "github.com/pip-services3-go/pip-services3-rpc-go/clients"
)

type SessionsHttpCommandableClientV1 struct {
	*cclients.CommandableHttpClient
	dataPageType  reflect.Type
	sessionV1Type reflect.Type
}

func NewSessionsHttpCommandableClientV1() *SessionsHttpCommandableClientV1 {
	c := &SessionsHttpCommandableClientV1{
		CommandableHttpClient: cclients.NewCommandableHttpClient("v1/sessions"),
		dataPageType:          reflect.TypeOf(&data.DataPage{}),
		sessionV1Type:         reflect.TypeOf(&SessionV1{}),
	}
	return c
}

func (c *SessionsHttpCommandableClientV1) GetSessions(correlationId string, filter *data.FilterParams,
	paging *data.PagingParams) (result *cdata.DataPage, err error) {

	params := cdata.NewAnyValueMapFromTuples(
		"filter", filter,
		"paging", paging,
	)

	res, err := c.CallCommand(c.dataPageType, "get_sessions", correlationId, params)
	if err != nil {
		return nil, err
	}

	result, _ = res.(*cdata.DataPage)
	return result, nil
}

func (c *SessionsHttpCommandableClientV1) GetSessionById(correlationId string, id string) (result *SessionV1, err error) {
	params := cdata.NewAnyValueMapFromTuples(
		"session_id", id,
	)

	res, err := c.CallCommand(c.sessionV1Type, "get_session_by_id", correlationId, params)
	if err != nil {
		return nil, err
	}

	result, _ = res.(*SessionV1)
	return result, nil
}

func (c *SessionsHttpCommandableClientV1) OpenSession(correlationId string, userId string, userName string,
	address string, client string, user interface{},
	data interface{}) (result *SessionV1, err error) {

	params := cdata.NewAnyValueMapFromTuples(
		"user_id", userId,
		"user_name", userName,
		"address", address,
		"client", client,
		"user", user,
		"data", nil,
	)

	res, err := c.CallCommand(c.sessionV1Type, "open_session", correlationId, params)
	if err != nil {
		return nil, err
	}

	result, _ = res.(*SessionV1)
	return result, nil

}

func (c *SessionsHttpCommandableClientV1) StoreSessionData(correlationId string, sessionId string,
	data interface{}) (result *SessionV1, err error) {
	params := cdata.NewAnyValueMapFromTuples(
		"session_id", sessionId,
		"data", data,
	)

	res, err := c.CallCommand(c.sessionV1Type, "store_session_data", correlationId, params)
	if err != nil {
		return nil, err
	}

	result, _ = res.(*SessionV1)
	return result, nil
}

func (c *SessionsHttpCommandableClientV1) UpdateSessionUser(correlationId string, sessionId string,
	user interface{}) (result *SessionV1, err error) {
	params := cdata.NewAnyValueMapFromTuples(
		"session_id", sessionId,
		"user", user,
	)

	res, err := c.CallCommand(c.sessionV1Type, "update_session_user", correlationId, params)
	if err != nil {
		return nil, err
	}

	result, _ = res.(*SessionV1)
	return result, nil
}

func (c *SessionsHttpCommandableClientV1) CloseSession(correlationId string, sessionId string) (result *SessionV1, err error) {
	params := cdata.NewAnyValueMapFromTuples(
		"session_id", sessionId,
	)

	res, err := c.CallCommand(c.sessionV1Type, "close_session", correlationId, params)
	if err != nil {
		return nil, err
	}

	result, _ = res.(*SessionV1)
	return result, nil
}

func (c *SessionsHttpCommandableClientV1) DeleteSessionById(correlationId string, sessionId string) (result *SessionV1, err error) {
	params := cdata.NewAnyValueMapFromTuples(
		"session_id", sessionId,
	)

	res, err := c.CallCommand(c.sessionV1Type, "delete_session_by_id", correlationId, params)
	if err != nil {
		return nil, err
	}

	result, _ = res.(*SessionV1)
	return result, nil
}
