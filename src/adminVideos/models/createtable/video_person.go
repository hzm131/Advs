package createtable

import "github.com/jinzhu/gorm"

//影视人员
type VideoPerson struct {
	gorm.Model
	Director string `gorm:"column:director;type:varchar;not null;"json:"director"` //导演
	Actor []Actor `gorm:"FOREIGNKEY:ActorId;"json:"actor_id"`
	ActorId int `gorm:"column:actor_id;type:integer;not null;"json:"actor_id"`
}

//演员
type Actor struct {
	gorm.Model
	Starring1 string `gorm:"column:starring1;type:varchar;"json:"starring1"` //主演
	Starring2 string `gorm:"column:starring2;type:varchar;"json:"starring2"`
	Starring3 string `gorm:"column:starring3;type:varchar;"json:"starring3"`
	Lead1 string `gorm:"column:lead1;type:varchar;"json:"lead1"` //领衔主演
	Lead2 string `gorm:"column:lead2;type:varchar;"json:"lead2"`
	Lead3 string `gorm:"column:lead3;type:varchar;"json:"lead3"`
}