package servers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"

	"github.com/cscatolini/pitaya-maestro-demo/protos"
	"github.com/gogo/protobuf/proto"
	"github.com/topfreegames/pitaya"
	"github.com/topfreegames/pitaya/component"
)

// ConnectorRemote is a remote that will receive rpc's
type ConnectorRemote struct {
	component.Base
}

// Connector struct
type Connector struct {
	component.Base
}

// SessionData struct
type SessionData struct {
	Data map[string]interface{}
	Msg  string
}

// Response struct
type Response struct {
	Code int32
	Msg  string
}

func reply(code int32, msg string) (*Response, error) {
	res := &Response{
		Code: code,
		Msg:  msg,
	}
	return res, nil
}

// GetSessionData gets the session data
func (c *Connector) GetSessionData(ctx context.Context) (*SessionData, error) {
	s := pitaya.GetSessionFromCtx(ctx)
	sv := pitaya.GetServer()
	msg := fmt.Sprintf("ServerID: %s, Type: %s", sv.ID, sv.Type)
	res := &SessionData{
		Data: s.GetData(),
		Msg:  msg,
	}
	return res, nil
}

// SetSessionData sets the session data
func (c *Connector) SetSessionData(ctx context.Context, data *SessionData) (*Response, error) {
	s := pitaya.GetSessionFromCtx(ctx)
	err := s.SetData(data.Data)
	if err != nil {
		return nil, pitaya.Error(err, "CN-000", map[string]string{"failed": "set data"})
	}
	sv := pitaya.GetServer()
	msg := fmt.Sprintf("ServerID: %s, Type: %s", sv.ID, sv.Type)
	return reply(200, msg)
}

// Docs returns documentation
func (c *ConnectorRemote) Docs(ctx context.Context, flag *protos.DocMsg) (*protos.Doc, error) {
	d, err := pitaya.Documentation(flag.GetGetProtos())

	if err != nil {
		return nil, err
	}
	doc, err := json.Marshal(d)

	if err != nil {
		return nil, err
	}

	return &protos.Doc{Doc: string(doc)}, nil
}

// ChangeRoomStatus changes a room status
func (c *Connector) ChangeRoomStatus(ctx context.Context, serverID, status string) (*Response, error) {
	rep := &protos.RPCRes{}
	arg := &protos.RPCMsg{Msg: status}
	err := pitaya.RPCTo(ctx, serverID, "room.changeroomstatus", rep, arg)
	if err != nil {
		return nil, pitaya.Error(err, "CN-000", map[string]string{"failed": "set data"})
	}
	return reply(200, rep.Msg)
}

// Proto returns server protos
func (c *ConnectorRemote) Proto(ctx context.Context, name *protos.ProtoName) (*protos.ProtoDescriptor, error) {
	protoName := name.Name
	protoReflectTypePointer := proto.MessageType(protoName)
	protoReflectType := protoReflectTypePointer.Elem()
	protoValue := reflect.New(protoReflectType)
	descriptorMethod, ok := protoReflectTypePointer.MethodByName("Descriptor")

	if !ok {
		return nil, errors.New("failed to get proto descriptor")
	}

	descriptorValue := descriptorMethod.Func.Call([]reflect.Value{protoValue})
	protoDescriptor := descriptorValue[0].Bytes()
	return &protos.ProtoDescriptor{
		Desc: protoDescriptor,
	}, nil
}
