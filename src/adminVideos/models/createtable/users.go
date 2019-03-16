package createtable

import (
	"github.com/jinzhu/gorm"
)

type Users struct {
	gorm.Model
	UserName string `gorm:"column:username;type:varchar;not null;"json:"username"`
	PassWord string `gorm:"column:password;type:varchar;not null;"json:"password"`

}


type Persons struct {
	gorm.Model
	Users Users `gorm:"FOREIGNKEY:UsersId;"json:"users_id"`
	UsersId int `gorm:"column:users_id;type:integer;not null"json:"users_id"`
	Age int `gorm:"column:age;type:integer;"json:"age"` //年龄
	Sex string `gorm:"column:sex;type:varchar;"json:"sex"` //性别
	Birthday string `gorm:"column:birthday;type:varchar;"json:"birthday"` //生日
	NickName string  `gorm:"column:nick_name;type:varchar;"json:"nick_name"` //昵称
	Vip int `gorm:"column:vip;type:integer;"json:"vip"`  //会员
}