package main

import (
	"github.com/jinzhu/gorm"
)

type Videos struct {
	gorm.Model
	VideoName string
	VideoPath string
	ClassIfication string
	VideoIntroduce string
	VideoImageId int
	VideoImage VideoImage
}

type VideoImage struct {
	gorm.Model
	ImagePath string
}


func ForVideoImage(){
	ImagePath := []string{
		"http://127.0.0.1:8000/static/images/akali.jpg",
		"http://127.0.0.1:8000/static/images/jie.jpg",
		"http://127.0.0.1:8000/static/images/yasuo.jpg",
		"http://127.0.0.1:8000/static/images/daomei.jpg",
	}

	for i:=0;i< len(ImagePath);i++{
		db.Create(&VideoImage{
			ImagePath: ImagePath[i],
		})
	}

}

func ForVideos(){
	VideoName := []string{
		"挖法1",
		"挖法2",
		"挖法3",
		"挖法4",
	}
	VideoPath := []string{
		"http://127.0.0.1:8000/static/videos/akali.jpg",
		"http://127.0.0.1:8000/static/videos/jie.jpg",
		"http://127.0.0.1:8000/static/videos/yasuo.jpg",
		"http://127.0.0.1:8000/static/videos/daomei.jpg",
	}
	ClassIfication := []string{
		"恐怖",
		"喜剧",
		"科幻",
		"限制级",
	}

	VideoIntroduce := []string{
		"世上最恐怖",
		"世上最喜剧",
		"世上最科幻",
		"世上限制级",
	}


	for i:=0;i<4;i++{
		db.Create(&Videos{
			VideoName: VideoName[i],
			VideoPath: VideoPath[i],
			ClassIfication: ClassIfication[i],
			VideoIntroduce: VideoIntroduce[i],
			VideoImageId: i+1,
		})
	}
}