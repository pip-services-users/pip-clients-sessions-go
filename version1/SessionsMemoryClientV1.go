package version1

import cdata "github.com/pip-services3-go/pip-services3-commons-go/data"

type SessionsMemoryClientV1 struct {
	sessions []SessionV1
}

func NewSessionsMemoryClientV1() *SessionsMemoryClientV1 {

	c := SessionsMemoryClientV1{
		sessions: make([]SessionV1, 0),
	}
	return &c
}

func (c *SessionsMemoryClientV1) GetSessions(correlationId string, filter *cdata.FilterParams, paging *cdata.PagingParams) (page *cdata.DataPage, err error) {

	var total int64 = (int64)(len(c.sessions))
	items := make([]interface{}, 0)
	for _, v := range c.sessions {
		item := v
		items = append(items, &item)
	}
	return cdata.NewDataPage(&total, items), nil
}

func (c *SessionsMemoryClientV1) GetSessionById(correlationId string, sessionId string) (session *SessionV1, err error) {
	for _, d := range c.sessions {
		if d.Id == sessionId {
			s := d
			session = &s
			break
		}
	}
	return session, nil
}

func (c *SessionsMemoryClientV1) OpenSession(correlationId string, userId string, userName string,
	address string, client string, user interface{}, data interface{}) (session *SessionV1, err error) {

	session = NewSessionV1("", userId, userName)
	session.Address = address
	session.Client = client
	session.User = user
	session.Data = data

	c.sessions = append(c.sessions, *session)
	return session, nil
}

func (c *SessionsMemoryClientV1) StoreSessionData(correlationId string, sessionId string, data interface{}) (session *SessionV1, err error) {

	for i := range c.sessions {
		if c.sessions[i].Id == sessionId {
			c.sessions[i].Data = data
			item := c.sessions[i]
			session = &item
			break
		}
	}

	return session, nil
}

func (c *SessionsMemoryClientV1) UpdateSessionUser(correlationId string, sessionId string, user interface{}) (session *SessionV1, err error) {

	for i := range c.sessions {
		if c.sessions[i].Id == sessionId {
			c.sessions[i].User = user
			item := c.sessions[i]
			session = &item
			break
		}
	}

	return session, nil
}

func (c *SessionsMemoryClientV1) CloseSession(correlationId string, sessionId string) (session *SessionV1, err error) {

	for i := range c.sessions {
		if c.sessions[i].Id == sessionId {
			c.sessions[i].Active = false
			item := c.sessions[i]
			session = &item
			break
		}
	}

	return session, nil
}

func (c *SessionsMemoryClientV1) DeleteSessionById(correlationId string, sessionId string) (session *SessionV1, err error) {

	var index = -1
	for i, v := range c.sessions {
		if v.Id == sessionId {
			index = i
			break
		}
	}

	if index < 0 {
		return nil, nil
	}

	var item = c.sessions[index]
	if index == len(c.sessions) {
		c.sessions = c.sessions[:index-1]
	} else {
		c.sessions = append(c.sessions[:index], c.sessions[index+1:]...)
	}
	return &item, nil
}
