package internal

import (
	"github.com/name5566/leaf/module"
	"wjleafserver/src/server/base"
	"github.com/name5566/leaf/gate"
)

var (
	skeleton = base.NewSkeleton()
	ChanRPC  = skeleton.ChanRPCServer

	Clients = make(map[string]client)
)

type Module struct {
	*module.Skeleton
}

func (m *Module) OnInit() {
	m.Skeleton = skeleton
}

func (m *Module) OnDestroy() {

}

type client struct {
	Token		string
	UserID		uint
	Agent		gate.Agent
	IsOnline	bool
}

func init()  {
	Clients = make(map[string]client)
}