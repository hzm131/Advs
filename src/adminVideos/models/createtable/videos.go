package createtable

import (
	"github.com/jinzhu/gorm"
)

/*
	上传视频视频
*/
type Videos struct {
	gorm.Model
	VideoName string `gorm:"column:video_name;type:varchar;not null;"json:"src"validate:"string,min=2,max=10"` //视频名称
	VideoType string `gorm:"column:video_type;type:varchar;not null"json:"video_type"` //视频类型
	VideoDate string  `gorm:"column:video_date;type:varchar;not null"json:"video_date"` //上映时间

	VideoPerson VideoPerson `gorm:"FOREIGNKEY:VideoPersonId;"json:"video_person_id"` //演员表
	VideoPersonId int `gorm:"column:video_person_id;type:integer;not null;"json:"video_person_id"`

	VideoIntroduce string  `gorm:"column:video_introduce;type:varchar;not null"json:"video_introduce"` //视频简介

	CommentId int `gorm:"column:comment_id;type:integer;not null;"json:"comment_id"` // 评论 一对多
	Comment []Comment `gorm:"FOREIGNKEY:VideosId;"json:"actor_id"`

	VideoImageId int `gorm:"column:video_image_id;type:varchar;not null;"json:"video_image_id"` //引入视频封面表
	VideoImage VideoImage `gorm:"FOREIGNKEY:VideoImageId;"json:"video_image"` //封面 belongs to 视频

	VideoSrcId int `gorm:"column:video_src_id;type:integer;not null;"json:"video_src_id"` //引入视频路径表
	VideoSrc VideoSrc `gorm:"FOREIGNKEY:VideoSrcId;"json:"video_src_id"` //视频路径
}



//视频路径
type VideoSrc struct {
	gorm.Model
	VideoPatch string `gorm:"column:video_path;type:varchar;not null;"json:"video_path"`
}

//视频封面
type VideoImage struct {
	gorm.Model
	ImagePath string `gorm:"column:image_path;type:varchar;not null;"json:"image_path"`
}








