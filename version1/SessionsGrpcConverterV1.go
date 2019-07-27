package version1

import (
	"encoding/json"

	"github.com/pip-services-users/pip-clients-sessions-go/protos"
	"github.com/pip-services3-go/pip-services3-commons-go/convert"
	"github.com/pip-services3-go/pip-services3-commons-go/data"
	"github.com/pip-services3-go/pip-services3-commons-go/errors"
)

func fromError(err error) *protos.ErrorDescription {
	if err == nil {
		return nil
	}

	desc := errors.ErrorDescriptionFactory.Create(err)
	obj := &protos.ErrorDescription{
		Type:          desc.Type,
		Category:      desc.Category,
		Code:          desc.Code,
		CorrelationId: desc.CorrelationId,
		Status:        convert.StringConverter.ToString(desc.Status),
		Message:       desc.Message,
		Cause:         desc.Cause,
		StackTrace:    desc.StackTrace,
		Details:       fromMap(desc.Details),
	}

	return obj
}

func toError(obj *protos.ErrorDescription) error {
	if obj == nil || (obj.Category == "" && obj.Message == "") {
		return nil
	}

	description := &errors.ErrorDescription{
		Type:          obj.Type,
		Category:      obj.Category,
		Code:          obj.Code,
		CorrelationId: obj.CorrelationId,
		Status:        convert.IntegerConverter.ToInteger(obj.Status),
		Message:       obj.Message,
		Cause:         obj.Cause,
		StackTrace:    obj.StackTrace,
		Details:       toMap(obj.Details),
	}

	return errors.ApplicationErrorFactory.Create(description)
}

func fromMap(val map[string]interface{}) map[string]string {
	r := map[string]string{}

	for k, v := range val {
		r[k] = convert.ToString(v)
	}

	return r
}

func toMap(val map[string]string) map[string]interface{} {
	var r map[string]interface{}

	for k, v := range val {
		r[k] = v
	}

	return r
}

func toJson(value interface{}) string {
	if value == nil {
		return ""
	}

	b, err := json.Marshal(value)
	if err != nil {
		return ""
	}
	return string(b[:])
}

func fromJson(value string) interface{} {
	if value == "" {
		return nil
	}

	var m interface{}
	json.Unmarshal([]byte(value), &m)
	return m
}

func fromSession(session *SessionV1) *protos.Session {
	if session == nil {
		return nil
	}

	obj := &protos.Session{
		Id:          session.Id,
		UserId:      session.UserId,
		UserName:    session.UserName,
		Active:      session.Active,
		OpenTime:    convert.StringConverter.ToString(session.OpenTime),
		CloseTime:   convert.StringConverter.ToString(session.CloseTime),
		RequestTime: convert.StringConverter.ToString(session.RequestTime),
		Address:     session.Address,
		Client:      session.Client,
		User:        toJson(session.User),
		Data:        toJson(session.Data),
	}

	return obj
}

func toSession(obj *protos.Session) *SessionV1 {
	if obj == nil {
		return nil
	}

	session := &SessionV1{
		Id:          obj.Id,
		UserId:      obj.UserId,
		UserName:    obj.UserName,
		Active:      obj.Active,
		OpenTime:    convert.DateTimeConverter.ToDateTime(obj.OpenTime),
		CloseTime:   convert.DateTimeConverter.ToDateTime(obj.CloseTime),
		RequestTime: convert.DateTimeConverter.ToDateTime(obj.RequestTime),
		Address:     obj.Address,
		Client:      obj.Client,
		User:        fromJson(obj.User),
		Data:        fromJson(obj.Data),
	}

	return session
}

func fromSessionPage(page *data.DataPage) *protos.SessionPage {
	if page == nil {
		return nil
	}

	obj := &protos.SessionPage{
		Total: *page.Total,
		Data:  make([]*protos.Session, len(page.Data)),
	}

	for i, v := range page.Data {
		session := v.(*SessionV1)
		obj.Data[i] = fromSession(session)
	}

	return obj
}

func toSessionPage(obj *protos.SessionPage) *data.DataPage {
	if obj == nil {
		return nil
	}

	sessions := make([]interface{}, len(obj.Data))

	for i, v := range obj.Data {
		sessions[i] = toSession(v)
	}

	page := data.NewDataPage(&obj.Total, sessions)

	return page
}
