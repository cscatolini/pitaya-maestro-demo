package servers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"time"

	"github.com/cscatolini/pitaya-maestro-demo/protos"
	"github.com/gogo/protobuf/proto"
	"github.com/topfreegames/pitaya"
	"github.com/topfreegames/pitaya/component"
	"github.com/topfreegames/pitaya/timer"
)

type (
	// Room represents a component that contains a bundle of room related handler
	// like Join/Message
	Room struct {
		component.Base
		group *pitaya.Group
		timer *timer.Timer
	}

	// UserMessage represents a message that user sent
	UserMessage struct {
		Name    string `json:"name"`
		Content string `json:"content"`
	}

	// Status represents a status string
	Status struct {
		ServerID string `json:"serverId"`
		Status   string `json:"status"`
	}

	// RPCResponse represents a rpc message
	RPCResponse struct {
		Msg string `json:"msg"`
	}

	// SendRPCMsg represents a rpc message
	SendRPCMsg struct {
		ServerID string `json:"serverId"`
		Route    string `json:"route"`
		Msg      string `json:"msg"`
	}

	// NewUser message will be received when new user join room
	NewUser struct {
		Content string `json:"content"`
	}

	// AllMembers contains all members uid
	AllMembers struct {
		Members []string `json:"members"`
	}

	// JoinResponse represents the result of joining room
	JoinResponse struct {
		Code   int    `json:"code"`
		Result string `json:"result"`
	}
)

// NewRoom returns a new room
func NewRoom() *Room {
	return &Room{
		group: pitaya.NewGroup("room"),
	}
}

// Init runs on service initialization
func (r *Room) Init() {}

// AfterInit component lifetime callback
func (r *Room) AfterInit() {
	// TODO: maestro logic here
	r.timer = pitaya.NewTimer(time.Minute, func() {
		println("UserCount: Time=>", time.Now().String(), "Count=>", r.group.Count())
	})
}

// Entry is the entrypoint
func (r *Room) Entry(ctx context.Context, uid []byte) (*JoinResponse, error) {
	logger := pitaya.GetDefaultLoggerFromCtx(ctx)
	s := pitaya.GetSessionFromCtx(ctx)

	err := s.Bind(ctx, string(uid))
	if err != nil {
		logger.Error("Failed to bind session")
		logger.Error(err)
		return nil, pitaya.Error(err, "RH-000", map[string]string{"failed": "bind"})
	}
	logger.Infof("Bound session to user id %s", string(uid))
	return &JoinResponse{Result: fmt.Sprintf("ok for uid %s", string(uid))}, nil
}

// GetSessionData gets the session data
func (r *Room) GetSessionData(ctx context.Context) (*SessionData, error) {
	time.Sleep(2 * time.Second)
	s := pitaya.GetSessionFromCtx(ctx)
	sv := pitaya.GetServer()
	msg := fmt.Sprintf("ServerID: %s, Type: %s", sv.ID, sv.Type)
	return &SessionData{
		Data: s.GetData(),
		Msg:  msg,
	}, nil
}

// SetSessionData sets the session data
func (r *Room) SetSessionData(ctx context.Context, data *SessionData) ([]byte, error) {
	logger := pitaya.GetDefaultLoggerFromCtx(ctx)
	s := pitaya.GetSessionFromCtx(ctx)
	err := s.SetData(data.Data)
	if err != nil {
		logger.Error("Failed to set session data")
		logger.Error(err)
		return nil, err
	}
	err = s.PushToFront(ctx)
	if err != nil {
		return nil, err
	}
	return []byte("success"), nil
}

// Join room
func (r *Room) Join(ctx context.Context) (*JoinResponse, error) {
	logger := pitaya.GetDefaultLoggerFromCtx(ctx)
	s := pitaya.GetSessionFromCtx(ctx)
	err := r.group.Add(s)
	if err != nil {
		logger.Error("Failed to join room")
		logger.Error(err)
		return nil, err
	}
	s.Push("onMembers", &AllMembers{Members: r.group.Members()})
	r.group.Broadcast("onNewUser", &NewUser{Content: fmt.Sprintf("New user: %d", s.ID())})
	if err != nil {
		logger.Error("Failed to broadcast onNewUser")
		logger.Error(err)
		return nil, err
	}
	return &JoinResponse{Result: "success"}, nil
}

// Message sync last message to all members
func (r *Room) Message(ctx context.Context, msg *UserMessage) {
	logger := pitaya.GetDefaultLoggerFromCtx(ctx)
	err := r.group.Broadcast("onMessage", msg)
	if err != nil {
		logger.Error("Error broadcasting message")
		logger.Error(err)
	}
}

// MessageRemote just echoes the given message
func (r *Room) MessageRemote(ctx context.Context, msg *UserMessage, b bool, s string) (*UserMessage, error) {
	return msg, nil
}

// ChangeRoomStatus changes a room status
func (r *Room) ChangeRoomStatus(ctx context.Context, msg *protos.RPCMsg) (*protos.RPCRes, error) {
	serverID := pitaya.GetServerID()
	res := fmt.Sprintf("Changed server %s status to  %s", serverID, msg.Msg)
	return &protos.RPCRes{Msg: res}, nil
}

// Docs returns documentation
func (r *Room) Docs(ctx context.Context, flag *protos.DocMsg) (*protos.Doc, error) {
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

// Proto returns server protos
func (r *Room) Proto(ctx context.Context, name *protos.ProtoName) (*protos.ProtoDescriptor, error) {
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
