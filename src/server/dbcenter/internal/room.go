package internal

import (
	"wjleafserver/src/server/base/model"
	"wjleafserver/src/server/msg"
	"time"
)

type Room struct {
	ID            int
	Name          string
	PeopleNumbers int //how man people
	CreatedAt     time.Time
	OverAt        time.Time
	GameNumbers   int
	CreatedBy     uint
	Seats         map[int]*Seat
	Status        int //0-wait 1-playing 2-apply dissolve 3-disssolved 4-over

	//Game          IGame
	ZhuangIndex   int
	AplDisUserID  uint
	PlaiedNumbers int
}

type Seat struct {
	IDX             int
	UID             uint
	UserName        string
	IsOnline        int
	IsReday         int //0-no 1-yes
	IsAgreeDissolve int //0-init 1-yes  2-no

	Score int
}

// 创建房间
func (r *Room)CreateRoomID() int {
	return 10010
}

// 加入房间
func (r *Room)Join(u *model.User) (*Room, bool) {
	return nil, false
}

// 准备
func (r *Room)SetReady(u *model.User) {

}

// 开始
func (r *Room)Begin()  {

}

// 离开
func (r *Room)Leave(u *model.User)  {

}

// 申请解散
func (r *Room) Dissolve(u *model.User) {

}

// 解散
func (r *Room) DoDissolve() {

}

func (r *Room) UnDissolve() {

}

// 清空房间用户
func (r *Room) Clear() {

}

// 同意解散
func (r *Room) AgreeDissolve(u *model.User) bool {
	return true
}

// 获取座位
func (r *Room) GetSeatByUID(uid uint) *Seat {
	return nil
}

// 获取房间信息
func (r *Room) GetRoomBaseInfo() *msg.RoomBaseInfo {
	return nil
}