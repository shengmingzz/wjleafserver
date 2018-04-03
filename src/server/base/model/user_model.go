package model

import (
	"time"
	"github.com/jinzhu/gorm"
	"github.com/name5566/leaf/gate"
)
/*
* 用户表
*/
type User struct {
	gorm.Model
	UserName  string
	IsOnline  int `gorm:"-"` // Ignore this field
	RoomID    int `gorm:"-"`
	RoomCards int
	Status    int
	Nickname  string
	Type      int
	Password  string
	CreatedAt time.Time
	Agent     gate.Agent `gorm:"-"`
}

func (u *User)TableName() string  {
	return "users"
}

func (u *User)GetOne(userId uint) (*User,error) {
	var user User
	result := MysqlConn.Where("id = ?",userId).First(&user)
	if result.Error != nil {
		return nil,result.Error
	}
	return &user,nil
}

func (u *User)FindByUserName(userName string) (*User,error) {
	var user User
	result := MysqlConn.Where("user_name = ?",userName).First(&user)
	if result.Error != nil {
		return nil,result.Error
	}
	return &user,nil
}
