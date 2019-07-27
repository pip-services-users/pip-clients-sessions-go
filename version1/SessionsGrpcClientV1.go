package version1

import (
	"github.com/pip-services-users/pip-clients-sessions-go/protos"
	"github.com/pip-services3-go/pip-services3-commons-go/data"
	"github.com/pip-services3-go/pip-services3-grpc-go/clients"
)

type SessionGrpcClientV1 struct {
	clients.GrpcClient
}

func NewSessionGrpcClientV1() *SessionGrpcClientV1 {
	return &SessionGrpcClientV1{
		GrpcClient: *clients.NewGrpcClient("sessions_v1.Sessions"),
	}
}

func (c *SessionGrpcClientV1) GetSessions(correlationId string, filter *data.FilterParams,
	paging *data.PagingParams) (result *data.DataPage, err error) {
	req := &protos.SessionPageRequest{
		CorrelationId: correlationId,
	}
	if filter != nil {
		req.Filter = filter.Value()
	}
	if paging != nil {
		req.Paging = &protos.PagingParams{
			Skip:  paging.GetSkip(0),
			Take:  (int32)(paging.GetTake(100)),
			Total: paging.Total,
		}
	}

	reply := new(protos.SessionPageReply)
	err = c.Call("get_sessions", correlationId, req, reply)
	if err != nil {
		return nil, err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return nil, err
	}

	result = toSessionPage(reply.Page)

	return result, nil
}

func (c *SessionGrpcClientV1) GetSessionById(correlationId string, id string) (result *SessionV1, err error) {
	req := &protos.SessionIdRequest{
		CorrelationId: correlationId,
		SessionId:     id,
	}

	reply := new(protos.SessionObjectReply)
	err = c.Call("get_session_by_id", correlationId, req, reply)
	if err != nil {
		return nil, err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return nil, err
	}

	result = toSession(reply.Session)

	return result, nil
}

func (c *SessionGrpcClientV1) OpenSession(correlationId string, userId string,
	userName string, address string, client string, user interface{},
	data interface{}) (result *SessionV1, err error) {
	req := &protos.SessionOpenRequest{
		CorrelationId: correlationId,
		UserId:        userId,
		UserName:      userName,
		Address:       address,
		Client:        client,
		User:          toJson(user),
		Data:          toJson(data),
	}

	reply := new(protos.SessionObjectReply)
	err = c.Call("open_session", correlationId, req, reply)
	if err != nil {
		return nil, err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return nil, err
	}

	result = toSession(reply.Session)

	return result, nil
}

func (c *SessionGrpcClientV1) StoreSessionData(correlationId string,
	sessionId string, data interface{}) (result *SessionV1, err error) {
	req := &protos.SessionStoreDataRequest{
		CorrelationId: correlationId,
		SessionId:     sessionId,
		Data:          toJson(data),
	}

	reply := new(protos.SessionObjectReply)
	err = c.Call("store_session_data", correlationId, req, reply)
	if err != nil {
		return nil, err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return nil, err
	}

	result = toSession(reply.Session)

	return result, nil
}

func (c *SessionGrpcClientV1) UpdateSessionUser(correlationId string,
	sessionId string, user interface{}) (result *SessionV1, err error) {
	req := &protos.SessionUpdateUserRequest{
		CorrelationId: correlationId,
		SessionId:     sessionId,
		User:          toJson(user),
	}

	reply := new(protos.SessionObjectReply)
	err = c.Call("update_session_user", correlationId, req, reply)
	if err != nil {
		return nil, err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return nil, err
	}

	result = toSession(reply.Session)

	return result, nil
}

func (c *SessionGrpcClientV1) CloseSession(correlationId string,
	sessionId string) (result *SessionV1, err error) {
	req := &protos.SessionIdRequest{
		CorrelationId: correlationId,
		SessionId:     sessionId,
	}

	reply := new(protos.SessionObjectReply)
	err = c.Call("close_session", correlationId, req, reply)
	if err != nil {
		return nil, err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return nil, err
	}

	result = toSession(reply.Session)

	return result, nil
}

func (c *SessionGrpcClientV1) DeleteSessionById(correlationId string,
	sessionId string) (result *SessionV1, err error) {
	req := &protos.SessionIdRequest{
		CorrelationId: correlationId,
		SessionId:     sessionId,
	}

	reply := new(protos.SessionObjectReply)
	err = c.Call("delete_session_by_id", "", req, reply)
	if err != nil {
		return nil, err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return nil, err
	}

	result = toSession(reply.Session)

	return result, nil
}
