package funciton

import "github.com/name5566/leaf/gate"

func SendMsg(a interface{}, m interface{}) {
	agent := a.(gate.Agent)
	agent.WriteMsg(m)

}