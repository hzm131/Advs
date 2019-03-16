package operating

import (
	"adminVideos/models"
	"adminVideos/models/createtable"
	"fmt"
)

//创建视频封面返回id
func CreatedImage(str string) (int,error){
	var img createtable.VideoImage
	img.ImagePath = str
	find := models.Db.Create(&img)
	if err:=find.Error; err!=nil{
		fmt.Println("创建失败",err)
		return 0,err
	}
	//创建成功后返回id
	query:=models.Db.First(&img)
	if err:=query.Error; err!=nil{
		fmt.Println("没有查询到",err)
		return 0,err
	}
	id := img.ID  //拿到id
	return id,nil //返回id
}

func GetVideos(offset int,limit int) ([]createtable.Videos){
	var videos []createtable.Videos
	models.Db.Offset(offset).Limit(limit).Order("created_at").Find(&videos)
	return videos
}




func CreatVideo(video createtable.Videos){
	models.Db.Create(video)
}